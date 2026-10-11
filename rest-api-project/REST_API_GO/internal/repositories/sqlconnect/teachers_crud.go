// Package sqlconnect
package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"restapi/internal/models"
)

func GetTeachersDBHandler(teachers []models.Teacher, r *http.Request) ([]models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return nil, err
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
	// fmt.Println(query.String())
	rows, err := db.Query(query.String(), args...)
	if err != nil {
		fmt.Println(err)
		// http.Error(w, "Database query error", http.StatusInternalServerError)
		return nil, err
	}

	defer rows.Close()

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
			// http.Error(w, "Error Scanning database results", http.StatusInternalServerError)
			return nil, err
		}
		// Added recommend by LLM to remove Diagnostics: 1. sql.Rows "rows" is used in Next loop at line 89 without final check of rows.Err() [default] in `rows, err := db.Query(query, args...)`

		if err := rows.Err(); err != nil {
			fmt.Println("Error iterating rows:", err)
			// http.Error(w, "Error reading database results", http.StatusInternalServerError)
			return nil, err
		}

		teachers = append(teachers, teacher)

	}
	return teachers, err
}

func GetTeachersDBHandlerByID(id int) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return models.Teacher{}, err
	}

	defer db.Close()

	var teacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE ID = ? ", id).
		Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)

	if err == sql.ErrNoRows {
		// http.Error(w, "Teacher not found", http.StatusNotFound)
		return models.Teacher{}, err
	} else if err != nil {
		fmt.Println("error in select query: ", err)
		// http.Error(w, "Database query error", http.StatusNotFound)
		return models.Teacher{}, err
	}
	return teacher, nil
}

func AddTeachersDBHandler(newTeachers []models.Teacher) ([]models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare(
		"INSERT INTO teachers (first_name, last_name, email, class, subject) VALUES(?,?,?,?,?);")
	if err != nil {
		// http.Error(w, "Error in preparing query", http.StatusInternalServerError)
		log.Println("error preparing: ", err)
		return nil, err
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
			// http.Error(w, "Error inserting data to database", http.StatusInternalServerError)
			return nil, err
		}

		lastID, err := resp.LastInsertId()
		if err != nil {
			// http.Error(w, "Error getting last ID", http.StatusInternalServerError)
			return nil, err
		}

		newTeacher.ID = int(lastID)
		addedTeachers[i] = newTeacher

	}
	return addedTeachers, nil
}

func PutTeacherDBHandler(id int, updateTeacher models.Teacher) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	defer db.Close()

	var existing models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE ID = ? ", id).
		Scan(&existing.ID, &existing.FirstName, &existing.LastName, &existing.Email, &existing.Class, &existing.Subject)
	if err != nil {
		if err == sql.ErrNoRows {
			// http.Error(w, "Teacher not found", http.StatusNotFound)
			return models.Teacher{}, err
		}
		// http.Error(w, "Unable to retrieve to database", http.StatusInternalServerError)
		return models.Teacher{}, err
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
		// http.Error(w, "Error updaating teacher", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	return updateTeacher, nil
}

func PatchTeachersDBHandler(updates []map[string]any) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return err
	}

	for _, update := range updates {
		idStr, ok := update["id"].(string)
		if !ok {
			tx.Rollback()
			// http.Error(w, "Invalid teachers ID in update", http.StatusBadRequest)
			return err
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println(err)
			return err
		}

		var teacherFromDB models.Teacher
		err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE ID = ? ", id).
			Scan(&teacherFromDB.ID, &teacherFromDB.FirstName, &teacherFromDB.LastName, &teacherFromDB.Email, &teacherFromDB.Class, &teacherFromDB.Subject)
		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {
				// http.Error(w, "Teacher not found", http.StatusNotFound)
				return err
			}
			// http.Error(w, "Unable to retrieve to database", http.StatusInternalServerError)
			return err
		}

		// Apply updates using reflect
		teacherVal := reflect.ValueOf(&teacherFromDB).Elem()
		teacherType := teacherVal.Type()

		for k, v := range update {
			if string(k) == "id" {
				continue
			}
			for i := 0; i < teacherVal.NumField(); i++ {
				field := teacherType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := teacherVal.Field(i)
					if fieldVal.CanSet() {
						val := reflect.ValueOf(v)
						if val.Type().ConvertibleTo(field.Type) {
							fieldVal.Set(val.Convert(field.Type))
						} else {
							tx.Rollback()
							log.Printf("Cannot convert %v to %v\n", val.Type(), field.Type)
						}
					}
					break
				}
			}
		}

		_, err = tx.Exec(
			"UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?",
			teacherFromDB.FirstName,
			teacherFromDB.LastName,
			teacherFromDB.Email,
			teacherFromDB.Class,
			teacherFromDB.Subject,
			teacherFromDB.ID,
		)
		if err != nil {
			tx.Rollback()
			// http.Error(w, "Error updating teacher", http.StatusInternalServerError)
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		// http.Error(w, "Error committing transaction", http.StatusInternalServerError)
		return err
	}
	return err
}

func PatchOneTeacherDBHandler(
	id int,
	updates map[string]any,
) (models.Teacher, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	defer db.Close()

	var existing models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE ID = ? ", id).
		Scan(&existing.ID, &existing.FirstName, &existing.LastName, &existing.Email, &existing.Class, &existing.Subject)
	if err != nil {
		if err == sql.ErrNoRows {
			// http.Error(w, "Teacher not found", http.StatusNotFound)
			return models.Teacher{}, err
		}
		// http.Error(w, "Unable to retrieve to database", http.StatusInternalServerError)
		return models.

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
			Teacher{}, err
	}

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
		// http.Error(w, "Error updaating teacher", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	return existing, nil
}

func DeleteOneTeacherDBHandler(id int) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM teachers where id = ?", id)
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
		return err
	}

	fmt.Println(result.RowsAffected())

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error retrieving delete teacher", http.StatusInternalServerError)
		return err
	}

	if rowsAffected == 0 {
		// http.Error(w, "Teacher not found", http.StatusInternalServerError)
		return err
	}
	return nil
}

func DeleteTeachersDBHandler(ids []int) ([]int, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return nil, err
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return nil, err
	}

	stmt, err := tx.Prepare("DELETE from teachers WHERE id = ?")
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error preparing delete statement", http.StatusInternalServerError)
		return nil, err
	}
	defer stmt.Close()

	deletedIDs := []int{}

	for _, id := range ids {
		result, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			// http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
			return nil, err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			// http.Error(w, "Error retrieving delete teacher", http.StatusInternalServerError)
			log.Println(err)
			return nil, err
		}

		if rowsAffected > 0 {
			deletedIDs = append(deletedIDs, id)
		} else {
			// http.Error(w, fmt.Sprintf("ID: %d does not exist", id), http.StatusInternalServerError)
			tx.Rollback()
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println("error commiting: ", err)
		// http.Error(w, "Error commiting transaction", http.StatusInternalServerError)
		return nil, err
	}

	if len(deletedIDs) < 1 {
		// http.Error(w, "ID's do not exist", http.StatusBadRequest)
		return nil, err
	}
	return deletedIDs, err
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
