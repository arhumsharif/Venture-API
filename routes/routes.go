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
	"github.com/gorilla/handlers"

)



func Test() {
	// code to write
	fmt.Println("In routes")
}

func HandleRequests() {
	fmt.Println("In Routes")
	Router := mux.NewRouter().StrictSlash(true)

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	
	// Router.HandleFunc("/", controllers.PrintHello)
	Router.HandleFunc("/user", controllers.InsertUser).Methods("POST")
	Router.HandleFunc("/education", controllers.InsertEducation).Methods("POST")
	Router.HandleFunc("/experience", controllers.InsertExperience).Methods("POST")
	Router.HandleFunc("/projects", controllers.InsertProject).Methods("POST")

	// Get Routes
	Router.HandleFunc("/user/{guid}", controllers.GetUsers).Methods("GET")
	Router.HandleFunc("/projects/{guid}", controllers.GetProjects).Methods("GET")
	Router.HandleFunc("/education/{guid}", controllers.GetEducations).Methods("GET")
	Router.HandleFunc("/experience/{guid}", controllers.GetExperiences).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", Router))

	// Delete Routes
	Router.HandleFunc("/user", controllers.DeleteUser).Methods("DELETE")
	Router.HandleFunc("/education", controllers.DeleteEducation).Methods("DELETE")
	Router.HandleFunc("/experience", controllers.DeleteExperience).Methods("DELETE")
	Router.HandleFunc("/projects", controllers.DeleteProject).Methods("DELETE")

	// Put Routes
	Router.HandleFunc("/user", controllers.UpdateUser).Methods("PUT")
	Router.HandleFunc("/education", controllers.UpdateEducation).Methods("PUT")
	Router.HandleFunc("/experience", controllers.UpdateExperience).Methods("PUT")
	Router.HandleFunc("/projects", controllers.UpdateProject).Methods("PUT")	
	// log.Fatal(http.ListenAndServe(":3000", Router))
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(Router)))
}
