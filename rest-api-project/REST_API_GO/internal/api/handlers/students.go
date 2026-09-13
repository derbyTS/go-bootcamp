// Package handlers
package handlers

import "net/http"

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello Get Students Route"))
	case http.MethodPost:
		w.Write([]byte("Hello Post Students Route"))
	case http.MethodPut:
		w.Write([]byte("Hello Put Students Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Students Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Students Route"))
	}
}
