package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/integrii/flaggy"
)

var errorRegistry = map[int]string{
	0001: "",
}

func main() {
	var caCertPath string
	var clientCertPath string
	var clientCertKeyPath string

	flaggy.String(&caCertPath, "", "cacert", "CA certificate to verify peer against")
	flaggy.String(&clientCertPath, "", "cert", "Client certificate file and password")
	flaggy.String(&clientCertKeyPath, "", "cert", "Client certificate file and password")
	flaggy.Parse()

	var tlsClientConfig *tls.Config

	clientCert, err := tls.LoadX509KeyPair(clientCertPath, clientCertKeyPath)

	if err != nil {
		panic(err)
	}

	if caCertPath != "" {
		caCert, err := os.ReadFile(caCertPath)

		if err != nil {
			panic(err)
		}

		caCertPool := x509.NewCertPool()

		if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
			panic("")
		}

		tlsClientConfig = &tls.Config{
			Certificates: []tls.Certificate{clientCert},
		}
		tlsClientConfig.RootCAs = caCertPool
	}

	transport := &http.Transport{
		TLSClientConfig: tlsClientConfig,
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}
	response, err := client.Get("")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
