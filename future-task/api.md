# Sample API standard by LLM

```go
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var u user

	// 1. Limit request body size to 1MB to prevent Denial of Service (DoS)
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	// 2. Initialize decoder and disallow unknown JSON fields
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// 3. Decode JSON stream into struct
	if err := decoder.Decode(&u); err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			http.Error(w, "Request body exceeds 1MB limit", http.StatusRequestEntityTooLarge)
			return
		}
		http.Error(w, "Invalid or malformed JSON payload", http.StatusBadRequest)
		return
	}

	// 4. Ensure there is no trailing garbage data after the JSON object
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		http.Error(w, "Request body must only contain a single JSON object", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"status":"success"}`)
}
```
