// Package middlewares
package middlewares

import (
	"fmt"
	"net/http"
	"slices"
)

var allowedOrigins = []string{
	"https://my-origin-url.com",
	"https://localhost:3000",
}

func Cors(next http.Handler) http.Handler {
	fmt.Println("Cors Middleware .....")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Cors Middleware return.....")
		origin := r.Header.Get("Origin")

		if origin != "" {
			if slices.Contains(allowedOrigins, origin) {
				w.Header().
					Set("Access-Control-Allow-Origin", origin)
					// Setting w.Header().Set("Access-Control-Allow-Credentials", "true") is useless if you don't return Access-Control-Allow-Origin: <origin>. The browser will throw a CORS error and block the response.
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Expose-Headers", "true")
				w.Header().Set("Access-Control-Max-Age", "3600")
			} else {
				http.Error(w, "Not allwed by Cors", http.StatusForbidden)
				return
			}
		}

		// Handle preflight OPTIONS requests directly
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent) // 204 No Content
			return
		}

		next.ServeHTTP(w, r)
		fmt.Println("Cors Middleware sent.....")
	})
}
