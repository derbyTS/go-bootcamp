// Package handlers
package handlers

import (
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"restapi/internal/models"
	"restapi/internal/repositories/sqlconnect"
	"restapi/pkg/utils"

	"golang.org/x/crypto/argon2"
)

// AddExecsHandler POST /execs/
func AddExecsHandler(w http.ResponseWriter, r *http.Request) {
	var newExecs []models.Exec
	var rawExecs []map[string]any

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, &rawExecs)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	fields := GetFieldNames(models.Exec{})

	allowedField := make(map[string]struct{})
	for _, field := range fields {
		allowedField[field] = struct{}{}
	}

	for _, exec := range rawExecs {
		for key := range exec {
			_, ok := allowedField[key]
			if !ok {
				http.Error(w, "Unnable to accept field found in request.", http.StatusBadRequest)
				return
			}
		}
	}

	err = json.Unmarshal(body, &newExecs)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	for _, exec := range newExecs {
		// if exec.FirstName == "" || exec.LastName == "" || exec.Email == "" ||
		// 	exec.Class == "" ||
		// 	exec.Subject == "" {
		// 	http.Error(w, "All fields are required", http.StatusBadRequest)
		// 	return
		// }
		err := CheckBlankFields(exec)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	addedExecs, err := sqlconnect.AddExecsDBHandler(newExecs)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string        `json:"status"`
		Count  int           `json:"count"`
		Data   []models.Exec `json:"data"`
	}{
		Status: "success",
		Count:  len(addedExecs),
		Data:   addedExecs,
	}

	json.NewEncoder(w).Encode(response)
}

// GetExecsHandler GET /execs/
func GetExecsHandler(w http.ResponseWriter, r *http.Request) {
	var execs []models.Exec
	execs, err := sqlconnect.GetExecsDBHandler(execs, r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := struct {
		Status string        `json:"status"`
		Count  int           `json:"count"`
		Data   []models.Exec `json:"data"`
	}{
		Status: "Success",
		Count:  len(execs),
		Data:   execs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetOneExecHandler /execs/{id}
func GetOneExecHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	// Handler Path Parameter
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	exec, err := sqlconnect.GetExecsDBHandlerByID(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exec)
}

// PatchExecsHandler PATCH /execs/
func PatchExecsHandler(w http.ResponseWriter, r *http.Request) {
	var updates []map[string]any
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = sqlconnect.PatchExecsDBHandler(updates)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// PatchOneExecHandler PATCH /execs/{id}
func PatchOneExecHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
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

	updatedExec, err := sqlconnect.PatchOneExecDBHandler(id, updates)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedExec)
}

// DeleteOneExecHandler DELETE /execs/{id}
func DeleteOneExecHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = sqlconnect.DeleteOneExecDBHandler(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "Exec successfully deleted",
		ID:     id,
	}
	json.NewEncoder(w).Encode(response)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Data Validation
	var req models.Exec
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password required", http.StatusBadRequest)
		return
	}

	// Check if the user exist
	db, err := sqlconnect.ConnectDB()
	if err != nil {
		utils.ErrorHandler(err, "Error connecting to database")
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	user := models.Exec{}
	err = db.QueryRow(`SELECT id, first_name, last_name, email, username, password, inactive_status, role FROM execs where username = ?`, req.Username).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Password, &user.InactiveStatus, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorHandler(err, "User not found")
			http.Error(w, "User not found", http.StatusInternalServerError)
			return
		}
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}

	// Check if the user is active

	if user.InactiveStatus {
		http.Error(w, "Account is inactive", http.StatusForbidden)
	}

	// Verify the password
	parts := strings.Split(user.Password, ".")
	if len(parts) != 2 {
		utils.ErrorHandler(errors.New("Invalid encoded hash format"), "Invalid encoded hash format")
		http.Error(w, "Invalid encoded hash format", http.StatusForbidden)
		return
	}
	saltBase64 := parts[0]
	passwordHashBase64 := parts[1]

	salt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		utils.ErrorHandler(err, "Failed to decode salt")
		http.Error(w, "Failed to decode salt", http.StatusInternalServerError)
		return
	}
	passwordHash, err := base64.StdEncoding.DecodeString(passwordHashBase64)
	if err != nil {
		utils.ErrorHandler(err, "Failed to decode password")
		http.Error(w, "Failed to decode password", http.StatusInternalServerError)
		return
	}

	inputPasswordHash := argon2.IDKey([]byte(req.Password), salt, 1, 64*1024, 4, 32)

	if len(inputPasswordHash) != len(passwordHash) {
		utils.ErrorHandler(errors.New("Incorect password"), "Incorect password")
		http.Error(w, "Incorect password", http.StatusForbidden)
		return
	}

	if subtle.ConstantTimeCompare(inputPasswordHash, passwordHash) == 1 {
	} else {
		utils.ErrorHandler(errors.New("Incorect password"), "Incorect password")
		http.Error(w, "Incorect password", http.StatusForbidden)
		return
	}

	// Generate Token
	tokenString, err := utils.SignToken(user.ID, req.Username, user.Role)
	if err != nil {
		utils.ErrorHandler(err, "Failed to generate token")
		http.Error(w, "Count not create token", http.StatusInternalServerError)
		return
	}

	// Send token as a response or as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Bearer",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "test",
		Value:    "test",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Token string `json:"Token"`
	}{
		Token: tokenString,
	}
	json.NewEncoder(w).Encode(response)
}
