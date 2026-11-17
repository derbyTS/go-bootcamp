// Package sqlconnect
package sqlconnect

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"restapi/internal/models"
	"restapi/pkg/utils"

	"github.com/go-mail/mail/v2"
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
		newExec.Password, err = utils.HashPassword(newExec.Password)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error adding execs")
		}

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

func UpdatePasswordInDb(
	userID int,
	currentPassword string,
	newPassword string,
) error {
	// ) (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Error connecting to database")
		// return "", utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	var username string
	var userPassword string
	var userRole string

	err = db.QueryRow("SELECT username, password, role FROM execs WHERE id = ?", userID).
		Scan(&username, &userPassword, &userRole)
	if err != nil {
		return utils.ErrorHandler(err, "User not found")
		// return "", utils.ErrorHandler(err, "User not found")
	}

	err = utils.VerifyPassword(userPassword, currentPassword)
	if err != nil {
		return utils.ErrorHandler(err, "Password not match")
		// return "", utils.ErrorHandler(err, "Password not match")
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		utils.ErrorHandler(err, "Error during hashing")
		return utils.ErrorHandler(err, "Internal error")
		// return "", utils.ErrorHandler(err, "Internal error")
	}

	currentTime := time.Now().Format(time.RFC3339)

	_, err = db.Exec(
		"UPDATE execs SET password = ?, password_changed_at = ? where id = ?",
		hashedPassword,
		currentTime,
		userID,
	)
	if err != nil {
		return utils.ErrorHandler(err, "Failed to update password")
		// return "", utils.ErrorHandler(err, "Failed to update password")
	}
	_, err = utils.SignToken(userID, username, userRole)
	// token, err := utils.SignToken(userID, username, userRole)
	if err != nil {
		return utils.ErrorHandler(err, "Password could not create token")
		// return "", utils.ErrorHandler(err, "Password could not create token")
	}
	return nil
	// return token, nil
}

func ForgotPasswordDbHandler(email string) error {
	db, err := ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Error connecting db")
	}
	defer db.Close()

	var exec models.Exec
	err = db.QueryRow("SELECT id FROM execs where email = ?", email).Scan(&exec.ID)
	if err != nil {
		return utils.ErrorHandler(err, "User not found")
	}

	duration, err := strconv.Atoi(os.Getenv("RESET_TOKEN_EXP_DURATION"))
	if err != nil {
		return utils.ErrorHandler(err, "Failed to send password email")
	}

	mins := time.Duration(duration)

	expiry := time.Now().Add(mins * time.Minute).Format(time.RFC3339)

	tokenBytes := make([]byte, 32)
	_, err = rand.Read(tokenBytes)
	if err != nil {
		return utils.ErrorHandler(err, "Failed to send password email")
	}

	token := hex.EncodeToString(tokenBytes)

	hashedToken := sha256.Sum256(tokenBytes)

	hashedTokenString := hex.EncodeToString(hashedToken[:])

	_, err = db.Exec(
		"UPDATE execs SET password_reset_token = ?, password_token_expire = ? WHERE id= ?",
		hashedTokenString,
		expiry,
		exec.ID,
	)
	if err != nil {
		return utils.ErrorHandler(err, "Failed to send password email")
	}

	// Send the reset email
	resetURL := fmt.Sprintf("https://localhost:3000/execs/resetpassword/reset/%s", token)
	message := fmt.Sprintf(
		"Forgot your password? Reset your password using the following link: \n%s\nIf you didn't request a password reset, please ignore this email. This link is only available for %d minutes",
		resetURL,
		int(mins),
	)

	m := mail.NewMessage()
	m.SetHeader("From", "schooladmin@school.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Your password reset link")
	m.SetBody("text/plain", message)

	d := mail.NewDialer("localhost", 1025, "", "")

	err = d.DialAndSend(m)
	if err != nil {
		return utils.ErrorHandler(err, "Failed to send password email")
	}
	return nil
}

func ResetPasswordHandlerDb(token, newPassword string) error {
	decodedInputTokenByte, err := hex.DecodeString(token)
	if err != nil {
		return utils.ErrorHandler(err, "Internal Error")
	}

	hashedToken := sha256.Sum256(decodedInputTokenByte)
	hashedTokenString := hex.EncodeToString(hashedToken[:])

	db, err := ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Error connecting db")
	}
	defer db.Close()

	var user models.Exec
	query := "SELECT id, email FROM execs where password_reset_token = ? AND password_token_expire > ?"
	err = db.QueryRow(query, hashedTokenString, time.Now().Format(time.RFC3339)).
		Scan(&user.ID, &user.Email)
	if err != nil {
		return utils.ErrorHandler(err, "Invalid or token expired")
	}

	// Hash the new password
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return utils.ErrorHandler(err, "Internal error")
	}

	updateQuery := "UPDATE execs SET password = ?, password_reset_token = NULL, password_token_expire = NULL, password_changed_at = ? WHERE id = ?"
	_, err = db.Exec(updateQuery, hashedPassword, time.Now().Format(time.RFC3339), user.ID)
	if err != nil {
		return utils.ErrorHandler(err, "Internal error")
	}
	return nil
}
