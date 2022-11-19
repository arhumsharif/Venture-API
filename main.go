package main

import (
	"fmt"
	"time"
	"main/routes"
    _ "github.com/go-sql-driver/mysql"
	// routes "./routes"
)


func main() {
	fmt.Println(time.Now().Second())
	fmt.Println("In Main")
	routes.HandleRequests()
}
