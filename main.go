package main

import (
	"fmt"
	"main/routes"
    _ "github.com/go-sql-driver/mysql"
	// routes "./routes"
)


func main() {
	fmt.Println("In Main")
	routes.HandleRequests()
}
