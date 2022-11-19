package db

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// var DB *sql.DB

// Connect DB
func ConnectDB() *sql.DB {
	DB, err := sql.Open("mysql", "venture:somePass@tcp(127.0.0.1:3306)/venture?parseTime=true")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB Connected")
	}
	return DB
}
