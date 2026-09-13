// Package handlers
package handlers

import "net/http"

func ExecsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello Get Execs Route"))
	case http.MethodPost:
		w.Write([]byte("Hello Post Execs Route"))
	case http.MethodPut:
		w.Write([]byte("Hello Put Execs Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Execs Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Execs Route"))
	}
}
