// Package router
package router

import (
	"net/http"

	"restapi/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.RootHandler)

	mux.HandleFunc("GET /teachers", handlers.GetTeachersHandler)
	mux.HandleFunc("POST /teachers", handlers.AddTeachersHandler)
	mux.HandleFunc("PATCH /teachers", handlers.PatchTeachersHandler)
	mux.HandleFunc("DELETE /teachers", handlers.DeleteTeachersHandler)

	mux.HandleFunc("PUT /teacher/{id}", handlers.PutTeacherHandler)
	mux.HandleFunc("GET /teacher/{id}", handlers.GetOneTeacherHandler)
	mux.HandleFunc("PATCH /teacher/{id}", handlers.PatchOneTeacherHandler)
	mux.HandleFunc("DELETE /teacher/{id}", handlers.DeleteOneTeacherHandler)

	mux.HandleFunc("GET /students", handlers.GetStudentsHandler)
	mux.HandleFunc("POST /students", handlers.AddStudentsHandler)
	mux.HandleFunc("PATCH /students", handlers.PatchStudentsHandler)
	mux.HandleFunc("DELETE /students", handlers.DeleteStudentsHandler)

	mux.HandleFunc("PUT /student/{id}", handlers.PutStudentHandler)
	mux.HandleFunc("GET /student/{id}", handlers.GetOneStudentHandler)
	mux.HandleFunc("PATCH /student/{id}", handlers.PatchOneStudentHandler)
	mux.HandleFunc("DELETE /student/{id}", handlers.DeleteOneStudentHandler)

	mux.HandleFunc("GET /execs/", handlers.ExecsHandler)

	return mux
}
