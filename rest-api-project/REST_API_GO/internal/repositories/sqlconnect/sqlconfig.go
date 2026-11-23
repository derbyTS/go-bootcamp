// Package sqlconnect
package sqlconnect

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
	// fmt.Println("Trying to connect to Database")

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DB_PORT")
	// connectionString := "temp:temp@tcp(127.0.0.1:3306)/" + dbName
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, dbPort, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		// panic(err)
		return nil, err
	}

	// THIS actually tests the connection.
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	// Configure pool limits so connections stay open and are reused under load.
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)

	// fmt.Println("Connected to DB..", dbName)
	DB = db
	return db, nil
}
