package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// 1. Load the client certificate and private key
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatalln("Failed to load client certificate and key:", err)
	}

	// 2. Load the CA certificate to trust the server (since it's self-signed)
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Fatalln("Failed to read server CA file:", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// 3. Configure TLS with client certificates and server root CA
	// tlsConfig := &tls.Config{
	// 	Certificates: []tls.Certificate{cert}, // Sends client cert to server
	// 	RootCAs:      caCertPool,              // Verifies server cert against this CA
	// }

	// Option A: Skip hostname verification (easiest for local learning)
	// We add this because our current openssl has no `x509_extensions = req_ext`
	// If openssl has `x509_extensions = req_ext` you can use the above tlsConfig
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert}, // Still sends client cert for mTLS!
		RootCAs:            caCertPool,
		InsecureSkipVerify: true, // Tells Go to ignore the missing SAN field
	}

	// 4. Attach TLS config to an HTTP Transport and Client
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// 5. Send the mTLS request
	resp, err := client.Get("https://localhost:3000/orders")
	if err != nil {
		log.Fatalln("mTLS Request Failed:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Failed to read response:", err)
	}

	fmt.Println("Server Response:", string(body))
}
