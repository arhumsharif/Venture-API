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
}

func HandleRequests() {
	fmt.Println("In Routes")
	Router := mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/", controllers.PrintHello)
	Router.HandleFunc("/user", controllers.InsertUser).Methods("POST")
	Router.HandleFunc("/education", controllers.InsertEducation).Methods("POST")
	Router.HandleFunc("/experience", controllers.InsertExperience).Methods("POST")
	Router.HandleFunc("/projects", controllers.InsertProject).Methods("POST")
	Router.HandleFunc("/user/{guid}", controllers.GetUsers).Methods("GET")
	Router.HandleFunc("/projects/{guid}", controllers.GetProjects).Methods("GET")
	Router.HandleFunc("/education/{guid}", controllers.GetEducations).Methods("GET")
	Router.HandleFunc("/experience/{guid}", controllers.GetExperiences).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", Router))
}
