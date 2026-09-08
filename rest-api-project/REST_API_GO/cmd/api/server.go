package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	// "time"

	mw "restapi/internal/api/middlewares"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello Get Teachers Route"))
		return
	case http.MethodPost:
		w.Write([]byte("Hello Post Teachers Route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello Put Teachers Route"))
		return
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Teachers Route"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Teachers Route"))
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello Get Students Route"))
		return
	case http.MethodPost:
		w.Write([]byte("Hello Post Students Route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello Put Students Route"))
		return
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Students Route"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Students Route"))
		return
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello Get Execs Route"))
		return
	case http.MethodPost:
		w.Write([]byte("Hello Post Execs Route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello Put Execs Route"))
		return
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Execs Route"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Execs Route"))
		return
	}
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

	// rl := mw.NewRateLimiter(5, time.Minute)
	//
	// hppOPtions := mw.HppOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www.form-urlencoded",
	// 	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// secureMux := mw.Cors(
	// 	rl.Throttle(
	// 		mw.ResponseTimeMiddleware(
	// 			mw.SecurityHeaders(
	// 				mw.Compression(
	// 					mw.Hpp(hppOPtions)(mux)))),
	// 	))
	// secureMux := applyMiddlewares(
	// 	mux,
	// 	mw.Hpp(hppOPtions),
	// 	mw.Compression,
	// 	mw.SecurityHeaders,
	// 	mw.ResponseTimeMiddleware,
	// 	rl.Throttle,
	// 	mw.Cors,
	// )

	secureMux := mw.SecurityHeaders(mux)

	server := &http.Server{
		Addr:      port,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on: ", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln(err)
	}
}

// Middleware is a function that wraps an http.Handler with additional functionality

type Middleware func(http.Handler) http.Handler

func applyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}
