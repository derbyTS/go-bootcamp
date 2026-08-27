package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type user struct {
	// Name []string `json:"name"` //You can make it an array if the request is expected to be array. Cannot unmarshal array into Go struct field user.name of type string
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

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
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
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
	})

	fmt.Println("Server is running on: ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
