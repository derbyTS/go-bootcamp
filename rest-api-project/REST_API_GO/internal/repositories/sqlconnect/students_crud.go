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
	"restapi/pkg/utils"
)

func GetStudentsDBHandler(students []models.Student, r *http.Request) ([]models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error connecting to database")
	}

	defer db.Close()
	// query := "SELECT id, first_name, last_name, email, class, subject FROM students WHERE 1=1"
	var query strings.Builder
	query.WriteString(
		"SELECT id, first_name, last_name, email, class FROM students WHERE 1=1",
	)

	// var args []interface{}
	var args []any

	args = utils.AddFilters(
		r,
		&query,
		args,
	) // Do not copy a non-zero Builder. https://go.dev/doc/effective_go?utm_source=chatgpt.com#pointers_vs_values

	utils.AddSorting(r, &query)

	// rows, err := db.Query(query, args...)
	// fmt.Println(query.String())
	rows, err := db.Query(query.String(), args...)
	if err != nil {
		fmt.Println(err)
		// http.Error(w, "Database query error", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Database query error")
	}

	defer rows.Close()

	for rows.Next() {
		var student models.Student
		err := rows.Scan(
			&student.ID,
			&student.FirstName,
			&student.LastName,
			&student.Email,
			&student.Class,
		)
		if err != nil {
			fmt.Println(err)
			// http.Error(w, "Error Scanning database results", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error Scanning database results")
		}
		// Added recommend by LLM to remove Diagnostics: 1. sql.Rows "rows" is used in Next loop at line 89 without final check of rows.Err() [default] in `rows, err := db.Query(query, args...)`

		if err := rows.Err(); err != nil {
			fmt.Println("Error iterating rows:", err)
			// http.Error(w, "Error reading database results", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error reading database results")
		}

		students = append(students, student)

	}
	return students, nil
}

func GetStudentsDBHandlerByID(id int) (models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return models.Student{}, utils.ErrorHandler(err, "Error connecting to database")
	}

	defer db.Close()

	var student models.Student
	err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE ID = ? ", id).
		Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)

	if err == sql.ErrNoRows {
		// http.Error(w, "Student not found", http.StatusNotFound)
		return models.Student{}, utils.ErrorHandler(err, "Student not found")
	} else if err != nil {
		fmt.Println("error in select query: ", err)
		// http.Error(w, "Database query error", http.StatusNotFound)
		return models.Student{}, utils.ErrorHandler(err, "Database query error")
	}
	return student, nil
}

func AddStudentsDBHandler(newStudents []models.Student) ([]models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	// stmt, err := db.Prepare(
	// 	"INSERT INTO students (first_name, last_name, email, class, subject) VALUES(?,?,?,?,?);")
	stmt, err := db.Prepare(utils.GenerateInsertQuery("students", models.Student{}))
	if err != nil {
		log.Println("error preparing: ", err)
		// http.Error(w, "Error in preparing query", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error in preparing query")
	}
	defer stmt.Close()

	addedStudents := make([]models.Student, len(newStudents))

	for i, newStudent := range newStudents {
		values := utils.GetStructValues(newStudent)
		resp, err := stmt.Exec(values...)
		if err != nil {
			if strings.Contains(err.Error(), "Error 1452") {
				return nil, utils.ErrorHandler(err, "class/class teacher does not exist")
			}
			return nil, utils.ErrorHandler(err, "Error inserting data to database")
		}

		lastID, err := resp.LastInsertId()
		if err != nil {
			// http.Error(w, "Error getting last ID", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error getting last ID")
		}

		newStudent.ID = int(lastID)
		addedStudents[i] = newStudent

	}
	return addedStudents, nil
}

func PutStudentDBHandler(id int, updateStudent models.Student) (models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return models.Student{}, utils.ErrorHandler(err, "Problem connecting to db")
	}
	defer db.Close()

	var existing models.Student
	err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE ID = ? ", id).
		Scan(&existing.ID, &existing.FirstName, &existing.LastName, &existing.Email, &existing.Class)
	if err != nil {
		if err == sql.ErrNoRows {
			// http.Error(w, "Student not found", http.StatusNotFound)
			return models.Student{}, utils.ErrorHandler(err, "Student not found")
		}
		// http.Error(w, "Unable to retrieve to database", http.StatusInternalServerError)
		return models.Student{}, utils.ErrorHandler(err, "Unable to retrieve to database")
	}

	updateStudent.ID = existing.ID
	_, err = db.Exec(
		"UPDATE students SET first_name = ?, last_name = ?, email = ?, class = ? WHERE id = ?",
		updateStudent.FirstName,
		updateStudent.LastName,
		updateStudent.Email,
		updateStudent.Class,
		updateStudent.ID,
	)
	if err != nil {
		// http.Error(w, "Error updaating student", http.StatusInternalServerError)
		return models.Student{}, utils.ErrorHandler(err, "Student not found")
	}
	return updateStudent, nil
}

func PatchStudentsDBHandler(updates []map[string]any) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Student not found")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Student not found")
	}

	for _, update := range updates {
		idStr, ok := update["id"].(string)
		if !ok {
			tx.Rollback()
			// http.Error(w, "Invalid students ID in update", http.StatusBadRequest)
			return utils.ErrorHandler(err, "Student not found")
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println(err)
			return utils.ErrorHandler(err, "Error with strconv.Atoi")
		}

		var studentFromDB models.Student
		err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE ID = ? ", id).
			Scan(&studentFromDB.ID, &studentFromDB.FirstName, &studentFromDB.LastName, &studentFromDB.Email, &studentFromDB.Class)
		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {
				// http.Error(w, "Student not found", http.StatusNotFound)
				return utils.ErrorHandler(err, "Student not found")
			}
			// http.Error(w, "Unable to retrieve to database", http.StatusInternalServerError)
			return utils.ErrorHandler(err, "Student not found")
		}

		// Apply updates using reflect
		studentVal := reflect.ValueOf(&studentFromDB).Elem()
		studentType := studentVal.Type()

		for k, v := range update {
			if string(k) == "id" {
				continue
			}
			for i := 0; i < studentVal.NumField(); i++ {
				field := studentType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := studentVal.Field(i)
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
			"UPDATE students SET first_name = ?, last_name = ?, email = ?, class = ? WHERE id = ?",
			studentFromDB.FirstName,
			studentFromDB.LastName,
			studentFromDB.Email,
			studentFromDB.Class,
			studentFromDB.ID,
		)
		if err != nil {
			tx.Rollback()
			// http.Error(w, "Error updating student", http.StatusInternalServerError)
			return utils.ErrorHandler(err, "Error updating student")
		}
	}
	err = tx.Commit()
	if err != nil {
		// http.Error(w, "Error committing transaction", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Student not found")
	}
	return nil
}

func PatchOneStudentDBHandler(
	id int,
	updates map[string]any,
) (models.Student, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return models.Student{}, utils.ErrorHandler(err, "Problem connecting to db")
	}
	defer db.Close()

	var existing models.Student
	err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE ID = ? ", id).
		Scan(&existing.ID, &existing.FirstName, &existing.LastName, &existing.Email, &existing.Class)
	if err != nil {
		if err == sql.ErrNoRows {
			// http.Error(w, "Student not found", http.StatusNotFound)
			return models.Student{}, utils.ErrorHandler(err, "Student not found")
		}
		// http.Error(w, "Unable to retrieve to database", http.StatusInternalServerError)
		return models.Student{}, utils.ErrorHandler(err, "Unable to retrieve to database")

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
	}

	// Apply updates using reflect
	studentVal := reflect.ValueOf(&existing).Elem()
	studentType := studentVal.Type()

	for k, v := range updates {
		for i := 0; i < studentVal.NumField(); i++ {
			field := studentType.Field(i)
			if field.Tag.Get("json") == k+",omitempty" {
				studentVal.Field(i).Set(reflect.ValueOf(v).Convert(studentVal.Field(i).Type()))
			}
		}
	}

	_, err = db.Exec(
		"UPDATE students SET first_name = ?, last_name = ?, email = ?, class = ? WHERE id = ?",
		existing.FirstName,
		existing.LastName,
		existing.Email,
		existing.Class,
		existing.ID,
	)
	if err != nil {
		// http.Error(w, "Error updaating student", http.StatusInternalServerError)
		return models.Student{}, utils.ErrorHandler(err, "Error updaating student")
	}
	return existing, nil
}

func DeleteOneStudentDBHandler(id int) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Problem connecting to db")
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM students where id = ?", id)
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error deleting student", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Student not found")
	}

	fmt.Println(result.RowsAffected())

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error retrieving delete student", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Student not found")
	}

	if rowsAffected == 0 {
		// http.Error(w, "Student not found", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Student not found")
	}
	return nil
}

func DeleteStudentsDBHandler(ids []int) ([]int, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Problem connecting to db")
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error starting transaction")
	}

	stmt, err := tx.Prepare("DELETE from students WHERE id = ?")
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error preparing delete statement", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error preparing delete statement")
	}
	defer stmt.Close()

	deletedIDs := []int{}

	for _, id := range ids {
		result, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			// http.Error(w, "Error deleting student", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error deleting student")
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			log.Println(err)
			// http.Error(w, "Error retrieving delete student", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error retrieving delete student")
		}

		if rowsAffected > 0 {
			deletedIDs = append(deletedIDs, id)
		} else {
			// http.Error(w, fmt.Sprintf("ID: %d does not exist", id), http.StatusInternalServerError)
			tx.Rollback()
			return nil, utils.ErrorHandler(err, fmt.Sprintf("ID: %d does not exist", id))
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println("error commiting: ", err)
		// http.Error(w, "Error commiting transaction", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error commiting transaction")
	}

	if len(deletedIDs) < 1 {
		// http.Error(w, "ID's do not exist", http.StatusBadRequest)
		return nil, utils.ErrorHandler(err, "ID's do not exist")
	}
	return deletedIDs, nil
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
	}

	return validFields[field]
}
