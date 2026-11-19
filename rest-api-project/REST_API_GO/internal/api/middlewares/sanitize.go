package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"restapi/pkg/utils"

	"github.com/microcosm-cc/bluemonday"
)

func XSSMiddleware(next http.Handler) http.Handler {
	fmt.Println("------------- Initializing XSSMiddleware -------------")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("------------- XSSMiddleware Run -------------")

		// Clean the URL Path
		sanitizedPath, err := clean(r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("Original Path: ", r.URL.Path)
		fmt.Println("Sanitize Path: ", sanitizedPath)

		// Sanitize query parameters
		params := r.URL.Query()
		sanitizedQuery := make(map[string][]string)
		for k, values := range params {
			sanitizeKey, err := clean(k)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			var sanitizedValues []string
			for _, value := range values {
				cleanValue, err := clean(value)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				sanitizedValues = append(sanitizedValues, cleanValue.(string))
			}
			sanitizedQuery[sanitizeKey.(string)] = sanitizedValues
			fmt.Printf("Original Query: %s: %s\n", k, strings.Join(values, ", "))
			fmt.Printf(
				"Sanitized Query: %s: %s\n",
				sanitizeKey,
				strings.Join(sanitizedValues, ", "),
			)
		}

		r.URL.Path = sanitizedPath.(string)
		r.URL.RawQuery = url.Values(sanitizedQuery).Encode()
		fmt.Println("Updated URL: ", r.URL.String())

		// Sanitize request body
		if r.Header.Get("Content-Type") == "application/json" {
			if r.Body != nil {
				bodyByte, err := io.ReadAll(r.Body)
				if err != nil {
					http.Error(
						w,
						utils.ErrorHandler(err, "Error reading request body").Error(),
						http.StatusBadRequest,
					)
					return
				}

				bodyString := strings.TrimSpace(string(bodyByte))
				fmt.Println("Original body: ", bodyString)

				// Reset the body
				r.Body = io.NopCloser(bytes.NewReader([]byte(bodyString)))

				if len(bodyString) > 0 {
					var inputData interface{}
					err := json.NewDecoder(bytes.NewReader([]byte(bodyString))).Decode(&inputData)
					if err != nil {
						http.Error(
							w,
							utils.ErrorHandler(err, "Invalid Json body").Error(),
							http.StatusBadRequest,
						)
						return
					}
					fmt.Println("Original json data: ", inputData)

					// Sanitize the JSON body
					sanitizedData, err := clean(inputData)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}

					fmt.Println("Sanitize JSON data: ", sanitizedData)

					sanitizedBody, err := json.Marshal(sanitizedData)
					if err != nil {
						http.Error(
							w,
							utils.ErrorHandler(err, "Error sanitizing body").Error(),
							http.StatusBadRequest,
						)
						return
					}

					r.Body = io.NopCloser(bytes.NewReader(sanitizedBody))
					fmt.Println("Sanitized body: ", string(sanitizedBody))
				} else {
					log.Println("Request body is empty")
				}
			} else {
				log.Println("No body in request")
			}
		} else if r.Header.Get("Content-Type") != "" {
			log.Printf(
				"Recieved request with unsupported Content-Type: %s. Expected application/json\n",
				r.Header.Get("Content-Type"),
			)
			http.Error(
				w,
				"Unsupported Content-Type. Please use application/json",
				http.StatusUnsupportedMediaType,
			)
		}

		next.ServeHTTP(w, r)
		fmt.Println("------------- Sending response from XSSMiddleware -------------")
	})
}

// Clean sanitizes input data to prevent XSS attach
func clean(data interface{}) (interface{}, error) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			v[key] = sanitizeValue(value)
		}
		return v, nil
	case []interface{}:
		for i, value := range v {
			v[i] = sanitizeValue(value)
		}
		return v, nil
	case string:
		return sanitizeString(v), nil
	default:
		return nil, utils.ErrorHandler(
			fmt.Errorf("Unsupported type: %T", data),
			fmt.Sprintf("Unsupported type: %T", data),
		)
	}
}

func sanitizeValue(data interface{}) interface{} {
	switch v := data.(type) {
	case string:
		return sanitizeString(v)
	case map[string]interface{}:
		for k, value := range v {
			v[k] = sanitizeValue(value)
		}
		return v
	case []interface{}:
		for i, value := range v {
			v[i] = sanitizeValue(value)
		}
		return v
	default:
		return v
	}
}

func sanitizeString(value string) string {
	return bluemonday.UGCPolicy().Sanitize(value)
}
