// Package handlers
package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"restapi/internal/models"
	"restapi/internal/repositories/sqlconnect"
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
	db, err := sqlconnect.ConnectDB()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}

	defer db.Close()
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	fmt.Println("ID String: ", idStr)

	if idStr == "" {
		// query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1"
		var query strings.Builder
		query.WriteString(
			"SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1",
		)
		// var args []interface{}
		var args []any

		args = addFilters(
			r,
			&query,
			args,
		) // Do not copy a non-zero Builder. https://go.dev/doc/effective_go?utm_source=chatgpt.com#pointers_vs_values

		// rows, err := db.Query(query, args...)
		rows, err := db.Query(query.String(), args...)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Database query error", http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		teacherList := make([]models.Teacher, 0)

		for rows.Next() {
			var teacher models.Teacher
			err := rows.Scan(
				&teacher.ID,
				&teacher.FirstName,
				&teacher.LastName,
				&teacher.Email,
				&teacher.Class,
				&teacher.Subject,
			)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error Scanning database results", http.StatusInternalServerError)
				return
			}
			// Added recommend by LLM to remove Diagnostics: 1. sql.Rows "rows" is used in Next loop at line 89 without final check of rows.Err() [default] in `rows, err := db.Query(query, args...)`

			if err := rows.Err(); err != nil {
				fmt.Println("Error iterating rows:", err)
				http.Error(w, "Error reading database results", http.StatusInternalServerError)
				return
			}

			teacherList = append(teacherList, teacher)

		}

		response := struct {
			Status string           `json:"status"`
			Count  int              `json:"count"`
			Data   []models.Teacher `json:"data"`
		}{
			Status: "Success",
			Count:  len(teacherList),
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

	var teacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE ID = ? ", id).
		Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)

	if err == sql.ErrNoRows {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("error in select query: ", err)
		http.Error(w, "Database query error", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
}

func addFilters(r *http.Request, query *strings.Builder, args []any) []any {
	params := map[string]string{
		"first_name": "first_name",
		"last_name":  "last_name",
		"email":      "email",
		"class":      "class",
		"subject":    "subject",
	}

	for param, dbField := range params {
		value := r.URL.Query().Get(param)
		if value != "" {
			// query += " AND " + dbField + " =?"
			query.WriteString(" AND ")
			query.WriteString(dbField)
			query.WriteString(" =?")
			args = append(args, value)
		}
	}
	return args
}

func addTeachersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlconnect.ConnectDB()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	var newTeachers []models.Teacher
	err = json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare(
		"INSERT INTO teachers (first_name, last_name, email, class, subject) VALUES(?,?,?,?,?);")
	if err != nil {
		http.Error(w, "Error in preparing query", http.StatusInternalServerError)
		fmt.Println("error preparing: ", err)
		return
	}
	defer stmt.Close()

	addedTeachers := make([]models.Teacher, len(newTeachers))

	for i, newTeacher := range newTeachers {
		resp, err := stmt.Exec(
			newTeacher.FirstName,
			newTeacher.LastName,
			newTeacher.Email,
			newTeacher.Class,
			newTeacher.Subject,
		)
		if err != nil {
			http.Error(w, "Error inserting data to database", http.StatusInternalServerError)
			return
		}

		lastID, err := resp.LastInsertId()
		if err != nil {
			http.Error(w, "Error getting last ID", http.StatusInternalServerError)
			return
		}

		newTeacher.ID = int(lastID)
		addedTeachers[i] = newTeacher

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
