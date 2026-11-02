package router

import (
	"net/http"

	"restapi/internal/api/handlers"
)

func studentsRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /students", handlers.GetStudentsHandler)
	mux.HandleFunc("POST /students", handlers.AddStudentsHandler)
	mux.HandleFunc("PATCH /students", handlers.PatchStudentsHandler)
	mux.HandleFunc("DELETE /students", handlers.DeleteStudentsHandler)

	mux.HandleFunc("PUT /student/{id}", handlers.PutStudentHandler)
	mux.HandleFunc("GET /student/{id}", handlers.GetOneStudentHandler)
	mux.HandleFunc("PATCH /student/{id}", handlers.PatchOneStudentHandler)
	mux.HandleFunc("DELETE /student/{id}", handlers.DeleteOneStudentHandler)

	return mux
}
