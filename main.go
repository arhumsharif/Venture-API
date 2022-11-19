package main

import (
	"fmt"
	"main/routes"
	"main/db"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
	// routes "./routes"
)


func main() {
	fmt.Println("In Main")
	var DB *sql.DB
	DB = db.ConnectDB()
	fmt.Println(DB)
	// insert, err := DB.Query("INSERT INTO user_details (user_guid, email, password, secret_key) VALUES ('3213232323', 'talha', '0123', '')")

    // // // if there is an error inserting, handle it
    // if err != nil {
    //     panic(err.Error())
    // }
    // // be careful deferring Queries if you are using transactions
    // defer insert.Close()
	routes.HandleRequests()
}