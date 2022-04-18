package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lucas-clemente/quic-go/http3"
)

var httpVersion = flag.String("version", "HTTP/1.1", "HTTP version")
var port = flag.String("port", "8080", "Port")
var name = flag.String("name", "http1", "Server Name")

var certPath = os.Getenv("CERT_PATH")
var keyPath = os.Getenv("KEY_PATH")

type HTTPProtocol string

const (
	HTTP1TLS HTTPProtocol = "HTTP/1.1TLS"
	HTTP2                 = "HTTP/2"
	HTTP3                 = "HTTP/3"
	HTTP1                 = "HTTP/1.1"
)

func main() {
	flag.Parse()
	os.Setenv("ELASTIC_APM_SERVICE_NAME", *name)
	fmt.Println(*httpVersion)
	fmt.Println("Port", *port)
	fmt.Println(certPath, keyPath)

	address := fmt.Sprintf(":%s", *port)

	version := HTTPProtocol(*httpVersion)

	SetupHttpHandler()

	switch version {
	case HTTP1:
		log.Fatal(http.ListenAndServe(address, nil))
		break
	case HTTP1TLS:
		log.Fatal(http.ListenAndServeTLS(address, certPath, keyPath, nil))
		break
	case HTTP2:
		log.Fatal(http.ListenAndServeTLS(address, certPath, keyPath, nil))
		break
	case HTTP3:
		log.Fatal(http3.ListenAndServeQUIC(address, certPath, keyPath, nil))
		break
	default:
		log.Fatalf("Not supported protocol: %s", *httpVersion)
	}
}
