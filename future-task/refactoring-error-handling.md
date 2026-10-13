# Refactoring Go Error Handling: Preserving HTTP Status Codes

When following a beginner tutorial, instructors often simplify error handling by logging the error and returning a generic string message (`fmt.Errorf("%s", message)`). While this keeps things simple early on, it strips away the underlying error type (like `sql.ErrNoRows`), forcing your HTTP handlers to return a `500 Internal Server Error` even when a resource isn't found (`404 Not Found`).

This guide outlines how to seamlessly upgrade your code _after_ your course to support proper error wrapping and dynamic messages.

---

## Step 1: Update `utils.ErrorHandler` to Wrap Errors

Instead of discarding the original error with `%s`, use `%w` to preserve it. This allows functions higher up the call stack to inspect the root cause.

```go
package utils

import (
	"fmt"
	"log"
	"os"
)

func ErrorHandler(err error, message string) error {
	errorLogger := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger.Println(message, err)

	// %s prints your custom message, %w wraps the original database error
	return fmt.Errorf("%s: %w", message, err)
}
```

---

## Step 2: Use Dynamic Messages with `fmt.Sprintf` in the Database Layer

When an error occurs (such as a missing record ID), you can construct a descriptive dynamic message using `fmt.Sprintf` and pass it directly into your updated `ErrorHandler`.

```go
var existing models.Teacher
err = db.QueryRow("SELECT id, first_name, ... FROM teachers WHERE ID = ?", id).
    Scan(&existing.ID, &existing.FirstName, ...)

if err != nil {
    if err == sql.ErrNoRows {
        // Construct a dynamic message and pass it to ErrorHandler
        customMsg := fmt.Sprintf("ID: %d does not exist", id)
        return models.Teacher{}, utils.ErrorHandler(err, customMsg)
    }

    return models.Teacher{}, utils.ErrorHandler(err, "Unable to retrieve from database")
}
```

---

## Step 3: Handle Status Codes Gracefully in HTTP Handlers

In your HTTP router/handler, use `errors.Is` to check what type of error was returned. This lets you send the correct HTTP status code to the client.

```go
import (
	"errors"
	"database/sql"
	"net/http"
)

func PutTeacherHandler(w http.ResponseWriter, r *http.Request) {
    // ... parse ID and decode body ...

	updatedTeacherFromDB, err := sqlconnect.PutTeacherDBHandler(id, updateTeacher)
	if err != nil {
		log.Println(err)

		// Check if the wrapped error is a "not found" scenario
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusNotFound) // Returns 404
			return
		}

		// Fallback for general database/server errors
		http.Error(w, err.Error(), http.StatusInternalServerError) // Returns 500
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTeacherFromDB)
}
```

---

### Why this approach is powerful:

1. **Keeps logs descriptive:** Your server logs will still clearly show things like `ID: 5 does not exist: sql: no rows in result set`.
2. **Restores correct API behavior:** Clients get accurate status codes (`404` instead of `500`) when a resource is missing.
3. **Follows Go idioms:** Leverages native `errors.Is` error unwrapping instead of parsing raw error strings.
