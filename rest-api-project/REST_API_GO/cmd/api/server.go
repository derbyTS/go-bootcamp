package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	mw "restapi/internal/api/middlewares"
)

type user struct {
	// Name []string `json:"name"` //You can make it an array if the request is expected to be array. Cannot unmarshal array into Go struct field user.name of type string
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello Get Teachers Route"))
		fmt.Println("Hello Get Teachers Route")
		return
	case http.MethodPost:
		w.Write([]byte("Hello Post Teachers Route"))
		fmt.Println("Hello Post Teachers Route")
		return
	case http.MethodPut:
		w.Write([]byte("Hello Put Teachers Route"))
		fmt.Println("Hello Put Teachers Route")
		return
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Teachers Route"))
		fmt.Println("Hello Patch Teachers Route")
		return
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Teachers Route"))
		fmt.Println("Hello Delete Teachers Route")
	}

	w.Write([]byte("Hello Default Teachers Route"))
	fmt.Println("Hello Default Teachers Route")
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello Get Students Route"))
		fmt.Println("Hello Get Students Route")
		return
	case http.MethodPost:
		w.Write([]byte("Hello Post Students Route"))
		fmt.Println("Hello Post Students Route")
		return
	case http.MethodPut:
		w.Write([]byte("Hello Put Students Route"))
		fmt.Println("Hello Put Students Route")
		return
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Students Route"))
		fmt.Println("Hello Patch Students Route")
		return
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Students Route"))
		fmt.Println("Hello Delete Students Route")
		return
	}
	w.Write([]byte("Hello Students Route"))
	fmt.Println("Hello Students Route")
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello Get Execs Route"))
		fmt.Println("Hello Get Execs Route")
		return
	case http.MethodPost:
		w.Write([]byte("Hello Post Execs Route"))
		fmt.Println("Hello Post Execs Route")
		return
	case http.MethodPut:
		w.Write([]byte("Hello Put Execs Route"))
		fmt.Println("Hello Put Execs Route")
		return
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Execs Route"))
		fmt.Println("Hello Patch Execs Route")
		return
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Execs Route"))
		fmt.Println("Hello Delete Execs Route")
		return
	}
	w.Write([]byte("Hello Execs Route"))
	fmt.Println("Hello Execs Route")
}

func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/students/", studentsHandler)

	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      port,
		Handler:   mw.SecurityHeaders(mw.Cors(mux)),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on: ", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln(err)
	}
}
