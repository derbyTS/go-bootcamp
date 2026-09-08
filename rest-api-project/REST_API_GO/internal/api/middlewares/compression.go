package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	fmt.Println("Compression Middleware .....")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Compression Middleware return.....")
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		// Set the response header
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()

		// Wrap the ResponseWriter
		w = &gzipResponseWritter{
			ResponseWriter: w,
			Writer:         gz,
		}

		next.ServeHTTP(w, r)
		fmt.Println("Sent response from Compression Middlware")
	})
}

type gzipResponseWritter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (g *gzipResponseWritter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}
