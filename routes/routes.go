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
	"github.com/rs/cors"

)



func Test() {
	// code to write
	fmt.Println("In routes")
}

func HandleRequests() {
	fmt.Println("In Routes")
	Router := mux.NewRouter().StrictSlash(true)


	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: false,
	})

	handler := c.Handler(Router)
	// with error handling
	
	// Router.HandleFunc("/", controllers.PrintHello)
	Router.HandleFunc("/authenticate", controllers.Authenticate).Methods("POST")
	Router.HandleFunc("/user", controllers.InsertUser).Methods("POST")
	Router.HandleFunc("/education", controllers.InsertEducation).Methods("POST")
	Router.HandleFunc("/experience", controllers.InsertExperience).Methods("POST")
	Router.HandleFunc("/projects", controllers.InsertProject).Methods("POST")
	Router.HandleFunc("/job", controllers.InsertJob).Methods("POST")
	Router.HandleFunc("/skill", controllers.InsertSkill).Methods("POST")
	Router.HandleFunc("/technology", controllers.InsertTechnology).Methods("POST")
	Router.HandleFunc("/user/job", controllers.InsertUserJob).Methods("POST")
	Router.HandleFunc("/job/skill", controllers.InsertJobSkill).Methods("POST")

	// Get Routes
	Router.HandleFunc("/user", controllers.GetUsers).Methods("GET")
	Router.HandleFunc("/projects", controllers.GetProjects).Methods("GET")
	Router.HandleFunc("/education", controllers.GetEducations).Methods("GET")
	Router.HandleFunc("/experience", controllers.GetExperiences).Methods("GET")
	Router.HandleFunc("/job", controllers.GetJobs).Methods("GET")
	Router.HandleFunc("/skill", controllers.GetSkills).Methods("GET")
	Router.HandleFunc("/admin/job/skill/{guid}", controllers.GetSpecificSkills).Methods("GET")
	Router.HandleFunc("/technology", controllers.GetTechnologies).Methods("GET")
	Router.HandleFunc("/user/job/{guid}", controllers.GetUserJobs).Methods("GET")
	Router.HandleFunc("/job/skill/{guid}", controllers.GetJobSkills).Methods("GET")
	// log.Fatal(http.ListenAndServe(":3000", Router))

	// Delete Routes
	Router.HandleFunc("/user", controllers.DeleteUser).Methods("DELETE")
	Router.HandleFunc("/education", controllers.DeleteEducation).Methods("DELETE")
	Router.HandleFunc("/experience", controllers.DeleteExperience).Methods("DELETE")
	Router.HandleFunc("/projects", controllers.DeleteProject).Methods("DELETE")

	// Put Routes
	Router.HandleFunc("/user", controllers.UpdateUser).Methods("PUT")
	Router.HandleFunc("/user/edit", controllers.UpdateUserDetail).Methods("PUT")
	Router.HandleFunc("/education", controllers.UpdateEducation).Methods("PUT")
	Router.HandleFunc("/experience", controllers.UpdateExperience).Methods("PUT")
	Router.HandleFunc("/projects", controllers.UpdateProject).Methods("PUT")	
	
	// start server listen
	log.Fatal(http.ListenAndServe(":8000", handler))
}
