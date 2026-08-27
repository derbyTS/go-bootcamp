package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Root Route"))
		fmt.Println("Hello Root Route")
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
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
			return
		}

		w.Write([]byte("Hello Default Teachers Route"))
		fmt.Println("Hello Default Teachers Route")
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Students Route"))
		fmt.Println("Hello Students Route")
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Execs Route"))
		fmt.Println("Hello Execs Route")
	})

	fmt.Println("Server is running on: ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
