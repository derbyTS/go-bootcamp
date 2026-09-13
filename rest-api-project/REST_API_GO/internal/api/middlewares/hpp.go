// Package middlewares
package middlewares

import (
	"fmt"
	"net/http"
	"slices"
	"strings"
)

type HppOptions struct {
	CheckQuery                  bool
	CheckBody                   bool
	CheckBodyOnlyForContentType string
	Whitelist                   []string
}

// Hpp - HTTP Parameter Pollution
func Hpp(options HppOptions) func(http.Handler) http.Handler {
	fmt.Println("HPP Middleware .....")
	return func(next http.Handler) http.Handler {
		fmt.Println("HPP Middleware return.....")
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if options.CheckBody && r.Method == http.MethodPost &&
				isCorrectContentType(r, options.CheckBodyOnlyForContentType) {
				filterBodyParams(r, options.Whitelist)
			}
			if options.CheckQuery && r.URL.Query() != nil {
				filterQueryParams(r, options.Whitelist)
			}
			next.ServeHTTP(w, r)
			fmt.Println("HPP Middleware sent.....")
		})
	}
}

func isCorrectContentType(r *http.Request, contentType string) bool {
	return strings.Contains(r.Header.Get("Content-Type"), contentType)
}

func filterBodyParams(r *http.Request, whiteList []string) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range r.Form {
		if len(v) > 1 {
			r.Form.Set(k, v[0])
			// r.Form.Set(k, v[len(v)-1])
		}
		if !isWhiteListed(k, whiteList) {
			delete(r.Form, k)
		}
	}
}

func filterQueryParams(r *http.Request, whiteList []string) {
	query := r.URL.Query()

	for k, v := range query {
		if len(v) > 1 {
			// query.Set(k, v[0])
			query.Set(k, v[len(v)-1])
		}
		if !isWhiteListed(k, whiteList) {
			delete(query, k)
		}
	}

	r.URL.RawQuery = query.Encode()
}

func isWhiteListed(param string, whiteList []string) bool {
	return slices.Contains(whiteList, param)
}
