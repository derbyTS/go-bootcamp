// Package handlers
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"restapi/internal/models"
)

var (
	teachers = make(map[int]models.Teacher)
	mutex    = &sync.Mutex{}
	nextID   = 1
)

func init() {
	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Calculus",
	}
	nextID++
	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Smith",
		Class:     "10A",
		Subject:   "Algebra",
	}
	nextID++
	teachers[nextID] = models.Teacher{
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

		teacherList := make([]models.Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) &&
				(lastName == "" || teacher.LastName == lastName) {
				teacherList = append(teacherList, teacher)
			}
		}

		response := struct {
			Status string           `json:"status"`
			Count  int              `json:"count"`
			Data   []models.Teacher `json:"data"`
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

	var newTeachers []models.Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
	}

	addedTeachers := make([]models.Teacher, len(newTeachers))

	for i, newTeacher := range newTeachers {
		newTeacher.ID = nextID
		teachers[nextID] = newTeacher
		addedTeachers[i] = newTeacher
		nextID++
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}

	json.NewEncoder(w).Encode(response)
}

func TeachersHandler(w http.ResponseWriter, r *http.Request) {
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
