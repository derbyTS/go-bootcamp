// Package sqlconnect
package sqlconnect

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"restapi/internal/models"
	"restapi/pkg/utils"

	"golang.org/x/crypto/argon2"
)

func GetExecsDBHandler(execs []models.Exec, r *http.Request) ([]models.Exec, error) {
	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error connecting to database")
	}

	defer db.Close()
	var query strings.Builder
	query.WriteString(
		"SELECT id, first_name, last_name, email, username, user_created_at, inactive_status, role FROM execs WHERE 1=1",
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
		var exec models.Exec
		err := rows.Scan(
			&exec.ID,
			&exec.FirstName,
			&exec.LastName,
			&exec.Email,
			&exec.Username,
			&exec.UserCreatedAt,
			&exec.InactiveStatus,
			&exec.Role,
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

		execs = append(execs, exec)

	}
	return execs, nil
}

func GetExecsDBHandlerByID(id int) (models.Exec, error) {
	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return models.Exec{}, utils.ErrorHandler(err, "Error connecting to database")
	}

	defer db.Close()

	var exec models.Exec
	err = db.QueryRow("SELECT id, first_name, last_name, email, username, user_created_at, inactive_status, role FROM execs WHERE ID = ? ", id).
		Scan(&exec.ID, &exec.FirstName, &exec.LastName, &exec.Email, &exec.Username, &exec.UserCreatedAt, &exec.InactiveStatus, &exec.Role)

	if err == sql.ErrNoRows {
		// http.Error(w, "Exec not found", http.StatusNotFound)
		return models.Exec{}, utils.ErrorHandler(err, "Exec not found")
	} else if err != nil {
		fmt.Println("error in select query: ", err)
		// http.Error(w, "Database query error", http.StatusNotFound)
		return models.Exec{}, utils.ErrorHandler(err, "Database query error")
	}
	return exec, nil
}

func AddExecsDBHandler(newExecs []models.Exec) ([]models.Exec, error) {
	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	stmt, err := db.Prepare(utils.GenerateInsertQuery("execs", models.Exec{}))
	if err != nil {
		log.Println("error preparing: ", err)
		// http.Error(w, "Error in preparing query", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error in preparing query")
	}
	defer stmt.Close()

	addedExecs := make([]models.Exec, len(newExecs))

	for i, newExec := range newExecs {
		if newExec.Password == "" {
			return nil, utils.ErrorHandler(errors.New("Password is blank"), "Please enter password")
		}
		salt := make([]byte, 16)
		_, err = rand.Read(salt)
		if err != nil {
			return nil, utils.ErrorHandler(
				errors.New("Failed to generate salt"),
				"Error adding data",
			)
		}
		hash := argon2.IDKey([]byte(newExec.Password), salt, 1, 64*1024, 4, 32)
		saltBase64 := base64.StdEncoding.EncodeToString(salt)
		hashBase64 := base64.StdEncoding.EncodeToString(hash)

		encodeHash := fmt.Sprintf("%s.%s", saltBase64, hashBase64)
		newExec.Password = encodeHash

		values := utils.GetStructValues(newExec)
		resp, err := stmt.Exec(values...)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error inserting data to database")
		}

		lastID, err := resp.LastInsertId()
		if err != nil {
			// http.Error(w, "Error getting last ID", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error getting last ID")
		}

		newExec.ID = int(lastID)
		addedExecs[i] = newExec

	}
	return addedExecs, nil
}

func PatchExecsDBHandler(updates []map[string]any) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Exec not found")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Exec not found")
	}

	for _, update := range updates {
		idStr, ok := update["id"].(string)
		if !ok {
			tx.Rollback()
			// http.Error(w, "Invalid execs ID in update", http.StatusBadRequest)
			return utils.ErrorHandler(err, "Exec not found")
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println(err)
			return utils.ErrorHandler(err, "Error with strconv.Atoi")
		}

		var execFromDB models.Exec
		err = db.QueryRow("SELECT id, first_name, last_name, email, username, inactive_status, role FROM execs WHERE ID = ? ", id).
			Scan(&execFromDB.ID, &execFromDB.FirstName, &execFromDB.LastName, &execFromDB.Email, &execFromDB.Username, &execFromDB.InactiveStatus, &execFromDB.Role)
		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {
				// http.Error(w, "Exec not found", http.StatusNotFound)
				return utils.ErrorHandler(err, "Exec not found")
			}
			// http.Error(w, "Unable to retrieve to database", http.StatusInternalServerError)
			return utils.ErrorHandler(err, "Exec not found")
		}

		// Apply updates using reflect
		execVal := reflect.ValueOf(&execFromDB).Elem()
		execType := execVal.Type()

		for k, v := range update {
			if string(k) == "id" {
				continue
			}
			for i := 0; i < execVal.NumField(); i++ {
				field := execType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := execVal.Field(i)
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
			"UPDATE execs SET first_name = ?, last_name = ?, email = ?, username = ?, inactive_status = ?,role = ? WHERE id = ?",
			execFromDB.FirstName,
			execFromDB.LastName,
			execFromDB.Email,
			execFromDB.Username,
			execFromDB.InactiveStatus,
			execFromDB.Role,
			execFromDB.ID,
		)
		if err != nil {
			tx.Rollback()
			return utils.ErrorHandler(err, "Error updating exec")
		}
	}
	err = tx.Commit()
	if err != nil {
		// http.Error(w, "Error committing transaction", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Exec not found")
	}
	return nil
}

func PatchOneExecDBHandler(
	id int,
	updates map[string]any,
) (models.Exec, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return models.Exec{}, utils.ErrorHandler(err, "Problem connecting to db")
	}
	defer db.Close()

	var existing models.Exec
	err = db.QueryRow("SELECT id, first_name, last_name, email, username, inactive_status, role FROM execs WHERE ID = ? ", id).
		Scan(&existing.ID, &existing.FirstName, &existing.LastName, &existing.Email, &existing.Username, &existing.InactiveStatus, &existing.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			// http.Error(w, "Exec not found", http.StatusNotFound)
			return models.Exec{}, utils.ErrorHandler(err, "Exec not found")
		}
		// http.Error(w, "Unable to retrieve to database", http.StatusInternalServerError)
		return models.Exec{}, utils.ErrorHandler(err, "Unable to retrieve to database")
	}

	// Apply updates using reflect
	execVal := reflect.ValueOf(&existing).Elem()
	execType := execVal.Type()

	for k, v := range updates {
		for i := 0; i < execVal.NumField(); i++ {
			field := execType.Field(i)
			if field.Tag.Get("json") == k+",omitempty" {
				execVal.Field(i).Set(reflect.ValueOf(v).Convert(execVal.Field(i).Type()))
			}
		}
	}

	_, err = db.Exec(
		"UPDATE execs SET first_name = ?, last_name = ?, email = ?, username = ?, inactive_status = ?,role = ? WHERE id = ?",
		existing.FirstName,
		existing.LastName,
		existing.Email,
		existing.Username,
		existing.InactiveStatus,
		existing.Role,
		existing.ID,
	)
	if err != nil {
		// http.Error(w, "Error updaating exec", http.StatusInternalServerError)
		return models.Exec{}, utils.ErrorHandler(err, "Error updaating exec")
	}
	return existing, nil
}

func DeleteOneExecDBHandler(id int) error {
	db, err := ConnectDB()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Problem connecting to db", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Problem connecting to db")
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM execs where id = ?", id)
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error deleting exec", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Exec not found")
	}

	fmt.Println(result.RowsAffected())

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		// http.Error(w, "Error retrieving delete exec", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Exec not found")
	}

	if rowsAffected == 0 {
		// http.Error(w, "Exec not found", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Exec not found")
	}
	return nil
}

func GetUserByUsername(username string) (models.Exec, error) {
	db, err := ConnectDB()
	if err != nil {
		return models.Exec{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	user := models.Exec{}
	err = db.QueryRow(`SELECT id, first_name, last_name, email, username, password, inactive_status, role FROM execs where username = ?`, username).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Password, &user.InactiveStatus, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Exec{}, utils.ErrorHandler(err, "User not found")
		}
		return models.Exec{}, utils.ErrorHandler(err, "Database error")
	}
	return user, nil
}
