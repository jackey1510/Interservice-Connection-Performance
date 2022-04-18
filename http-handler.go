package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"go.elastic.co/apm/module/apmhttp/v2"
)

type responseBody struct {
	Body           any `json:"body"`
	ResponseTimeMs int `json:"responseTimeMs"`
}

type getBody struct {
	Url          string
	Protocol     HTTPProtocol
	RequestCount int
}

func index(w http.ResponseWriter, r *http.Request) {
	defer log.Println(r.Method, r.UserAgent())
	fmt.Fprintf(w, "Hello World from %s", *name)
}

func httpGet(client *http.Client, url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func responseJson(w http.ResponseWriter, object any) {
	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(object); err != nil {
		handleBadRequest(w, err, "Error encodeing json")
	}
}

func sendRequest(gb getBody) ([]byte, error) {
	httpClient := GetHttpClient(gb.Protocol)
	var res []byte
	var err error
	// go routine
	var wg sync.WaitGroup
	wg.Add(gb.RequestCount)
	for i := 0; i < gb.RequestCount; i++ {
		go func() {
			res, err = httpGet(httpClient, gb.Url)
			wg.Done()
		}()
	}
	wg.Wait()
	return res, err
}

func handleBadRequest(w http.ResponseWriter, e error, message string) {
	errMessage := fmt.Sprintf("%s: %s", message, e.Error())
	log.Println(errMessage)
	http.Error(w, errMessage, http.StatusBadRequest)
}

func getFromUrl(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	url, protocol, count := queries.Get("url"), queries.Get("protocol"), queries.Get("requestCount")

	countInt, err := strconv.Atoi(count)

	if err != nil {
		handleBadRequest(w, err, fmt.Sprintf("request count is not integer: %s", count))
		return
	}

	gb := getBody{Url: url, Protocol: HTTPProtocol(protocol), RequestCount: countInt}

	startTime := time.Now().UnixMilli()
	res, err := sendRequest(gb)

	if err != nil {
		handleBadRequest(w, err, fmt.Sprintf("Error sending request to url %s", gb.Url))
		return
	}

	responseTimeMs := time.Now().UnixMilli() - startTime

	responseJson(w, responseBody{Body: string(res), ResponseTimeMs: int(responseTimeMs)})
}

func SetupHttpHandler() {
	http.Handle("/", apmhttp.Wrap(http.HandlerFunc(index)))
	http.Handle("/get", apmhttp.Wrap(http.HandlerFunc(getFromUrl)))
}
