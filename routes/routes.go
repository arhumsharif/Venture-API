package routes

import ( 
	"fmt"
"main/controllers"
// "encoding/json"
	// "io/ioutil"
	"log"
	"net/http"
	// "strings"
	// "sync"
	// "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

)

func Test() {
	// code to write
	fmt.Println("In routes")
	controllers.FetchBookName()
}

func HandleRequests() {
	fmt.Println("In Routes")
	Router := mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/", controllers.PrintHello)
	Router.HandleFunc("/try", controllers.TryPost).Methods("POST")
	Router.HandleFunc("/register", controllers.InsertUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", Router))
}