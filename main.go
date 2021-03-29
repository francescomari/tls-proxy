package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func run() error {
	var (
		source   string
		target   string
		certFile string
		keyFile  string
	)

	flag.StringVar(&source, "source", ":8080", "The address to listen to")
	flag.StringVar(&target, "target", "http://localhost:8081", "The URL to proxy")
	flag.StringVar(&certFile, "cert", "cert.pem", "Path to the certificate")
	flag.StringVar(&keyFile, "key", "key.pem", "Path to the private key")
	flag.Parse()

	targetURL, err := url.Parse(target)
	if err != nil {
		return err
	}

	return http.ListenAndServeTLS(source, certFile, keyFile, httputil.NewSingleHostReverseProxy(targetURL))
}
