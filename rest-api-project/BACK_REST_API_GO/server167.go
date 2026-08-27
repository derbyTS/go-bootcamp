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
			// Parse form data (necessary for x-www-form-urlencoded) **************************************

			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing from", http.StatusBadRequest)
				return
			}
			// fmt.Println("Form: ", r.Form)
			// fmt.Println(r.Form.Get("name"))

			response := make(map[string]any)

			for k, v := range r.Form {
				response[k] = v
			}

			fmt.Println("Form data: ", response)
			// names := r.Form["name"]
			// fmt.Println("Name(s): ", names[0])

			// Parse Raw Body **************************************

			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Fatalln(err)
				return
			}

			defer r.Body.Close()

			fmt.Println(string(body))

			// If you expect json data, then unmarshal it(according to instructor)

			var userInstance user

			err = json.Unmarshal(body, &userInstance)
			if err != nil {
				log.Fatalln(err)
				return
			}

			// Get the value without a struct
			response1 := make(map[string]any)

			err = json.Unmarshal(body, &response1)
			if err != nil {
				log.Fatalln(err)
				return
			}

			fmt.Println("userInstance: ", userInstance)
			fmt.Println("response1: ", response1)

			// I try using NewDecoder right here - Suggest by llm that since json.NewDecoder is designed to stream data directly from an io.Reader(like r.Body) without loading the entire request payload into memory first this is optimal
			// if err := json.NewDecoder(r.Body).Decode(&userInstance); err != nil {
			// 	http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			// 	return
			// }

			fmt.Println(userInstance.Name)

			// Access the request details **************************************
			fmt.Println("Body: ", r.Body)
			fmt.Println("Form: ", r.Form)
			fmt.Println("Header: ", r.Header)
			fmt.Println("Context: ", r.Context())
			fmt.Println("ContextLength: ", r.ContentLength)
			fmt.Println("Host: ", r.Host)
			fmt.Println("Method: ", r.Method)
			fmt.Println("Protocol: ", r.Proto)
			fmt.Println("Remote Addr: ", r.RemoteAddr)
			fmt.Println("Request URI: ", r.RequestURI)
			fmt.Println("TLS: ", r.TLS)
			fmt.Println("Trailer: ", r.Trailer)
			fmt.Println("Transfer Encoding: ", r.TransferEncoding)
			fmt.Println("URL: ", r.URL)
			fmt.Println("User Agent: ", r.UserAgent())
			fmt.Println("URL Port: ", r.URL.Port())

			// *****************************
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
