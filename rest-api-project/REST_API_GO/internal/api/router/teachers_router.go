package router

import (
	"net/http"

	"restapi/internal/api/handlers"
)

func teachersRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /teachers", handlers.GetTeachersHandler)
	mux.HandleFunc("POST /teachers", handlers.AddTeachersHandler)
	mux.HandleFunc("PATCH /teachers", handlers.PatchTeachersHandler)
	mux.HandleFunc("DELETE /teachers", handlers.DeleteTeachersHandler)

	mux.HandleFunc("PUT /teacher/{id}", handlers.PutTeacherHandler)
	mux.HandleFunc("GET /teacher/{id}", handlers.GetOneTeacherHandler)
	mux.HandleFunc("PATCH /teacher/{id}", handlers.PatchOneTeacherHandler)
	mux.HandleFunc("DELETE /teacher/{id}", handlers.DeleteOneTeacherHandler)

	mux.HandleFunc("GET /teacher/{id}/students", handlers.GetStudentsByTeacherID)
	mux.HandleFunc("GET /teacher/{id}/studentscount", handlers.GetStudentsCountByTeacherID)

	return mux
}
