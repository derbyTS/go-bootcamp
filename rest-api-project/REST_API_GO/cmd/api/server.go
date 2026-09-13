package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	// "time"

	mw "restapi/internal/api/middlewares"
)

type Teacher struct {
	ID        int    `json:"id",omitempty`
	FirstName string `json:"first_name",omitempty`
	LastName  string `json:"last_name",omitempty`
	Class     string `json:"class",omitempty`
	Subject   string `json:"subject",omitempty`
}

var (
	teachers = make(map[int]Teacher)
	mutex    = &sync.Mutex{}
	nextID   = 1
)

func init() {
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Calculus",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Smith",
		Class:     "10A",
		Subject:   "Algebra",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Doe",
		Class:     "11A",
		Subject:   "Physics",
	}
	nextID++
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	fmt.Println("ID String: ", idStr)

	if idStr == "" && idStr != " " {
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")

		teacherList := make([]Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) &&
				(lastName == "" || teacher.LastName == lastName) {
				teacherList = append(teacherList, teacher)
			}
		}

		response := struct {
			Status string    `json:"status"`
			Count  int       `json:"count"`
			Data   []Teacher `json:"data"`
		}{
			Status: "Success",
			Count:  len(teachers),
			Data:   teacherList,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
	// Handler Path Parameter
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	teacher, exist := teachers[id]
	if !exist {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(teacher)
}

func addTeachersHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var newTeachers []Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
	}

	addedTeachers := make([]Teacher, len(newTeachers))

	for i, newTeacher := range newTeachers {
		newTeacher.ID = nextID
		teachers[nextID] = newTeacher
		addedTeachers[i] = newTeacher
		nextID++
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string    `json:"status"`
		Count  int       `json:"count"`
		Data   []Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}

	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTeachersHandler(w, r)
	case http.MethodPost:
		addTeachersHandler(w, r)
	case http.MethodPut:
		w.Write([]byte("Hello Put Teachers Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Teachers Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Teachers Route"))
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
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

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello Get Execs Route"))
	case http.MethodPost:
		w.Write([]byte("Hello Post Execs Route"))
	case http.MethodPut:
		w.Write([]byte("Hello Put Execs Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Execs Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Execs Route"))
	}
}

func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/students/", studentsHandler)

	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// rl := mw.NewRateLimiter(5, time.Minute)
	//
	// hppOPtions := mw.HppOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www.form-urlencoded",
	// 	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// secureMux := mw.Cors(
	// 	rl.Throttle(
	// 		mw.ResponseTimeMiddleware(
	// 			mw.SecurityHeaders(
	// 				mw.Compression(
	// 					mw.Hpp(hppOPtions)(mux)))),
	// 	))
	// secureMux := applyMiddlewares(
	// 	mux,
	// 	mw.Hpp(hppOPtions),
	// 	mw.Compression,
	// 	mw.SecurityHeaders,
	// 	mw.ResponseTimeMiddleware,
	// 	rl.Throttle,
	// 	mw.Cors,
	// )

	secureMux := mw.SecurityHeaders(mux)

	server := &http.Server{
		Addr:      port,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on: ", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln(err)
	}
}

// Middleware is a function that wraps an http.Handler with additional functionality

type Middleware func(http.Handler) http.Handler

func applyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}
