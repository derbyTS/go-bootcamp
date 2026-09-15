// Package sqlconnect
package sqlconnect

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(dbName string) (*sql.DB, error) {
	fmt.Println("Trying to connect to Database")
	connectionString := "temp:temp@tcp(127.0.0.1:3306)/" + dbName
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		// panic(err)
		return nil, err
	}

	fmt.Println("Connected to DB..", dbName)
	return db, nil
}
