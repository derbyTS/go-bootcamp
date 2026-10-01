package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Method based routing
	mux.HandleFunc("POST /items/create/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Items Created")
	})

	mux.HandleFunc("DELETE /items/create/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Items Deleted")
	})

	// Wildcard in pattern - path parameter
	mux.HandleFunc("GET /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Item %s\n", r.PathValue("id"))
	})

	// Wildcard with ...
	mux.HandleFunc("/files/{path...}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Path: %s\n", r.PathValue("path"))
	})

	mux.HandleFunc("/path/{param1}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Param1: %s\n", r.PathValue("param1"))
	})

	// This will cause an error cause the above pattern confuse with this
	// mux.HandleFunc("/{param2}/path2", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Param2: %s\n", r.PathValue("param2"))
	// })

	mux.HandleFunc("/path1/path2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Nothing here")
	})

	http.ListenAndServe(":8080", mux)
}
