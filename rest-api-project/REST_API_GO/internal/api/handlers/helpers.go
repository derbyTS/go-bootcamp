package handlers

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"restapi/pkg/utils"
)

func CheckBlankFields(value any) error {
	val := reflect.ValueOf(value)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String && field.String() == "" {
			// http.Error(w, "All fields are required", http.StatusBadRequest)
			return utils.ErrorHandler(
				errors.New("All fields are required"),
				"All fields are required",
			)
		}
	}
	return nil
}

func GetFieldNames(model any) []string {
	val := reflect.TypeOf(model)
	fields := []string{}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldToAdd := strings.TrimSuffix(field.Tag.Get("json"), ",omitempty")
		fields = append(fields, fieldToAdd)
	}
	return fields
}

func getPaginationParams(r *http.Request) (int, int) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	return page, limit
}
