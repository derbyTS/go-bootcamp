// Package handlers
package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"restapi/internal/models"
	"restapi/internal/repositories/sqlconnect"
)

// AddTeachersHandler POST /teachers/
func AddTeachersHandler(w http.ResponseWriter, r *http.Request) {
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

// GetTeachersHandler GET /teachers/
func GetTeachersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlconnect.ConnectDB()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}

	defer db.Close()

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

	addSorting(r, &query)

	// rows, err := db.Query(query, args...)
	fmt.Println(query.String())
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

// GetTeacherHandler /teachers/{id}
func GetTeacherHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlconnect.ConnectDB()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	idStr := r.PathValue("id")

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

// PutTeacherHandler PUT /teachers/{id}
func PutTeacherHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/teachers/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updateTeacher models.Teacher
	err = json.NewDecoder(r.Body).Decode(&updateTeacher)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	db, err := sqlconnect.ConnectDB()
	if err != nil {
		log.Println(err)
		http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var existing models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE ID = ? ", id).
		Scan(&existing.ID, &existing.FirstName, &existing.LastName, &existing.Email, &existing.Class, &existing.Subject)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Unable to retrieve to database", http.StatusInternalServerError)
		return
	}

	updateTeacher.ID = existing.ID
	_, err = db.Exec(
		"UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?",
		updateTeacher.FirstName,
		updateTeacher.LastName,
		updateTeacher.Email,
		updateTeacher.Class,
		updateTeacher.Subject,
		updateTeacher.ID,
	)
	if err != nil {
		http.Error(w, "Error updaating teacher", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateTeacher)
}

// PatchTeacherHandler PATCH /teachers/{id}
func PatchTeacherHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/teachers/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updates map[string]any
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	db, err := sqlconnect.ConnectDB()
	if err != nil {
		log.Println(err)
		http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var existing models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE ID = ? ", id).
		Scan(&existing.ID, &existing.FirstName, &existing.LastName, &existing.Email, &existing.Class, &existing.Subject)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Unable to retrieve to database", http.StatusInternalServerError)
		return
	}

	// Apply updates
	// for k, v := range updates {
	// 	switch k {
	// 	case "first_name":
	// 		existing.FirstName = v.(string)
	// 	case "last_name":
	// 		existing.LastName = v.(string)
	// 	case "email":
	// 		existing.Email = v.(string)
	// 	case "class":
	// 		existing.Class = v.(string)
	// 	case "subject":
	// 		existing.Subject = v.(string)
	// 	}
	// }

	// Apply updates using reflect
	teacherVal := reflect.ValueOf(&existing).Elem()
	teacherType := teacherVal.Type()

	for k, v := range updates {
		for i := 0; i < teacherVal.NumField(); i++ {
			field := teacherType.Field(i)
			if field.Tag.Get("json") == k+",omitempty" {
				teacherVal.Field(i).Set(reflect.ValueOf(v).Convert(teacherVal.Field(i).Type()))
			}
		}
	}

	_, err = db.Exec(
		"UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?",
		existing.FirstName,
		existing.LastName,
		existing.Email,
		existing.Class,
		existing.Subject,
		existing.ID,
	)
	if err != nil {
		http.Error(w, "Error updaating teacher", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existing)
}

// DeleteTeacherHandler DELETE /teachers/{id}
func DeleteTeacherHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/teachers/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	db, err := sqlconnect.ConnectDB()
	if err != nil {
		log.Println(err)
		http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM teachers where id = ?", id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
		return
	}

	fmt.Println(result.RowsAffected())

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		http.Error(w, "Error retrieving delete teacher", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Teacher not found", http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "Teacher successfully deleted",
		ID:     id,
	}
	json.NewEncoder(w).Encode(response)
}

func addSorting(r *http.Request, query *strings.Builder) {
	sortParams := r.URL.Query()["sortBy"]
	if len(sortParams) > 0 {
		// query += " ORDER BY ?"
		query.WriteString(" ORDER BY")

		for i, param := range sortParams {
			parts := strings.Split(param, ":")
			if len(parts) != 2 {
				continue
			}

			field, order := parts[0], parts[1]
			fmt.Println("field: ", field)
			fmt.Println("order: ", order)

			if !isValidSortField(field) || !isValidSortOrder(order) {
				continue
			}

			if i > 0 {
				query.WriteString(",")
			}

			query.WriteString(" ")
			query.WriteString(field)
			query.WriteString(" ")
			query.WriteString(order)

		}
	}
}

func isValidSortOrder(order string) bool {
	return order == "asc" || order == "desc"
}

func isValidSortField(field string) bool {
	validFields := map[string]bool{
		"first_name": true,
		"last_name":  true,
		"email":      true,
		"class":      true,
		"subject":    true,
	}

	return validFields[field]
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
