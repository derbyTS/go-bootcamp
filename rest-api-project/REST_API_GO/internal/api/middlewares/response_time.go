package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request in ResponseTime")
		start := time.Now()

		// Create a custom ResponseWriter to capture the status code
		wrappedWriter := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		// Calculate the duration
		period := time.Since(start)

		// To be corrected
		wrappedWriter.Header().Set("X-Response-Time", period.String())

		next.ServeHTTP(wrappedWriter, r)

		period = time.Since(start)

		// Log the request details
		fmt.Printf(
			"Method: %s, URL: %s, Status: %d, Duration: %v\n",
			r.Method,
			r.URL,
			wrappedWriter.statusCode,
			period,
		)
		fmt.Println("Sent response from ResponseTimeMiddleware")
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
