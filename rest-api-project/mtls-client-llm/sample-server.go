package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func loadClientCAs() *x509.CertPool {
	clientCAs := x509.NewCertPool()
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Fatalln("Could not load client CA: ", err)
	}

	clientCAs.AppendCertsFromPEM(caCert)
	return clientCAs
}

func main() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling incoming users")
	})

	port := 3000

	// Load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	// Basic TLS configuration
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		// ClientAuth: tls.RequireAndVerifyClientCert, // Enforce mTLS
		// ClientCAs:  loadClientCAs(),                // Enforce mTLS
	}

	// Create a custom server
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	// Enable http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running in port: ", port)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln(err)
	}

	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}

func logRequestDetails(r *http.Request) {
	httpVersion := r.Proto
	fmt.Println("Received request with http version: ", httpVersion)

	if r.TLS != nil {
		tlsVersion := getTLSVersionName(r.TLS.Version)
		fmt.Println("Received request with TLS version ", tlsVersion)
	} else {
		fmt.Println("Request received without TLS")
	}
}

func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "1.0"
	case tls.VersionTLS11:
		return "1.1"
	case tls.VersionTLS12:
		return "1.2"
	case tls.VersionTLS13:
		return "1.3"
	default:
		return "Unknow TLS Version"
	}
}
