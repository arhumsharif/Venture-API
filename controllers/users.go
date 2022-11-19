package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
	"net/http"
	"main/models"
	// "strings"
	// "sync"

	// "github.com/dgrijalva/jwt-go"
	// "github.com/gorilla/mux"
	"main/db"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
)

type Book struct{
	Name string `json:"name"`
	Author string `json:"author"`
}

func PrintHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
	fmt.Println(r)
	// w.Header().Set("Content-Type", "application/json")
	// var myBook Book
	// myBook.Name = "Witcher"
	// myBook.Author = "Arhum"
	// // json.NewEncoder(w).Encode(myBook)
	// jData, err := json.Marshal(myBook)
	// if err != nil {
	// 	// handle error
	// }
	// fmt.Println(jData)
	// w.write(jData)
}

func TryPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var mybook Book 
	json.Unmarshal(reqBody, &mybook)


	json.NewEncoder(w).Encode(mybook)

	newData, err := json.Marshal(mybook)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newData))
	}

	fmt.Println(mybook.Name)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	// Get Body

	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myUser models.User

	json.Unmarshal(reqBody, &myUser)
	json.NewEncoder(w).Encode(myUser)

	// Perform Query
	fmt.Println("INSERT INTO user_details (user_guid, email, password, secret_key) VALUES (" + "'" +  myUser.User_Guid + "'" + "," + "'" + myUser.Email + "'" + "," + "'" + myUser.Password + "'" + ", '')")
	DB = db.ConnectDB()
	insert, err := DB.Query("INSERT INTO user_details (user_guid, email, password, secret_key) VALUES (" + "'" +  myUser.User_Guid + "'" + "," + "'" + myUser.Email + "'" + "," + "'" + myUser.Password + "'" + ", '')")

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
}