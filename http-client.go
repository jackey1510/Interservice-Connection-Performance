package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lucas-clemente/quic-go/http3"
	"go.elastic.co/apm/module/apmhttp/v2"
	"golang.org/x/net/http2"
)

func getTlsConfig(protocol HTTPProtocol) *tls.Config {
	if protocol == HTTP1 {
		return &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	// Create a pool with the server certificate since it is not signed
	// by a known CA

	caCert, err := ioutil.ReadFile(certPath)

	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
		// Allow self signed certs
		InsecureSkipVerify: true,
	}
	return tlsConfig
}

func GetHttpClient(protocol HTTPProtocol) *http.Client {
	client := &http.Client{}

	tlsConfig := getTlsConfig(protocol)
	// Use the proper transport in the client
	switch protocol {
	case HTTP1:
	case HTTP1TLS:
		client.Transport = &http.Transport{
			TLSClientConfig:   tlsConfig,
			DisableKeepAlives: false,
		}
		break
	case HTTP2:
		client.Transport = &http2.Transport{
			TLSClientConfig: tlsConfig,
		}
		break
	case HTTP3:
		client.Transport = &http3.RoundTripper{
			TLSClientConfig: tlsConfig,
		}
		break
	}

	return apmhttp.WrapClient(client)
}
