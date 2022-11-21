package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
	"net/http"
	"main/models"
	"main/utils"
	// "strings"
	// "sync"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"main/db"
    _ "github.com/go-sql-driver/mysql"
    "github.com/google/uuid"
    "database/sql"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	// Get Body
	guid := uuid.New() // user guid
	id := guid.String()	
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myUser models.User

	json.Unmarshal(reqBody, &myUser)
	json.NewEncoder(w).Encode(myUser)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("INSERT INTO user_details (user_guid, email, password, secret_key) VALUES (?, ?, ?, ?)", id, myUser.Email, myUser.Password, myUser.Secret_Key)

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Success"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		panic(resErr.Error())
	}
	w.Write(jsonResponse)
    defer insert.Close()
}

func InsertEducation(w http.ResponseWriter, r *http.Request) {
	// Get Body
	guid := uuid.New() // education guid
	education_id := guid.String()	
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myEducation models.Education

	json.Unmarshal(reqBody, &myEducation)
	json.NewEncoder(w).Encode(myEducation)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("INSERT INTO user_education (education_guid, user_guid, school, user_from, user_to, degree) VALUES (" + "'" +  education_id + "'" + "," + "'" + myEducation.User_Guid + "'" + "," + "'" + myEducation.School + "'" + "," + "'" + myEducation.User_From + "'"  + "," + "'" + myEducation.User_To + "'" + "," + "'" + myEducation.Degree + "'" + ");")

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Success"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		panic(resErr.Error())
	}
	w.Write(jsonResponse)
    defer insert.Close()
}

func InsertExperience(w http.ResponseWriter, r *http.Request) {
	// Get Body
	guid := uuid.New() // education guid
	experience_id := guid.String()	
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myExperience models.Experience

	json.Unmarshal(reqBody, &myExperience)
	json.NewEncoder(w).Encode(myExperience)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("INSERT INTO user_experience (experience_guid, user_guid, job_title, job_description, company_name, user_from, user_to) VALUES (" + "'" +  experience_id + "'" + "," + "'" + myExperience.User_Guid + "'" + "," + "'" + myExperience.Job_Title + "'" + "," + "'" + myExperience.Job_Description + "'"  + "," + "'" + myExperience.Company_Name + "'" + "," + "'" + myExperience.User_From + "'" + "," + "'" + myExperience.User_To + "'" + ");")

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Success"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		panic(resErr.Error())
	}
	w.Write(jsonResponse)
    defer insert.Close()
}

func InsertProject(w http.ResponseWriter, r *http.Request) {
	// Get Body
	guid := uuid.New() // education guid
	project_id := guid.String()	
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myProject models.Project

	json.Unmarshal(reqBody, &myProject)
	json.NewEncoder(w).Encode(myProject)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("INSERT INTO user_projects (project_guid, user_guid, title, description, technologies) VALUES (" + "'" +  project_id + "'" + "," + "'" + myProject.User_Guid + "'" + "," + "'" + myProject.Title + "'" + "," + "'" + myProject.Description + "'"  + "," + "'" + myProject.Technologies + "'"  + ");")

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Success"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		panic(resErr.Error())
	}
	w.Write(jsonResponse)
    defer insert.Close()
}


func GetUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	fmt.Println("guid", guid)
	// get data against guid
	DB := db.ConnectDB()
	rows, err:= DB.Query("SELECT * FROM user_details WHERE user_guid=?",guid)
	if err != nil {
		fmt.Println("Error:", err)
	}
	finalData, err := utils.SQLToJSON(rows)
	if err != nil{
		fmt.Println("Error:", err)
	}
	w.Write(finalData)
	defer DB.Close()
}



func GetExperiences(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	fmt.Println("guid", guid)
	// get data against guid
	DB := db.ConnectDB()
	rows, err:= DB.Query("SELECT * FROM user_experience WHERE user_guid=?",guid)
	if err != nil {
		fmt.Println("Error:", err)
	}
	finalData, err := utils.SQLToJSON(rows)
	if err != nil{
		fmt.Println("Error:", err)
	}
	w.Write(finalData)
	defer DB.Close()
}

func GetEducations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	fmt.Println("guid", guid)
	// get data against guid
	DB := db.ConnectDB()
	rows, err:= DB.Query("SELECT * FROM user_education WHERE user_guid=?",guid)
	if err != nil {
		fmt.Println("Error:", err)
	}
	finalData, err := utils.SQLToJSON(rows)
	if err != nil{
		fmt.Println("Error:", err)
	}
	w.Write(finalData)
	defer DB.Close()
}


func GetProjects(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	fmt.Println("guid", guid)
	// get data against guid
	DB := db.ConnectDB()
	rows, err:= DB.Query("SELECT * FROM user_projects WHERE user_guid=?",guid)
	if err != nil {
		fmt.Println("Error:", err)
	}
	finalData, err := utils.SQLToJSON(rows)
	if err != nil{
		fmt.Println("Error:", err)
	}
	w.Write(finalData)
	defer DB.Close()
}




// ------------------------ Delete Routes ----------------------------

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get Body
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myUser models.User

	json.Unmarshal(reqBody, &myUser)
	json.NewEncoder(w).Encode(myUser)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("UPDATE user_details set is_deleted = '1' WHERE user_guid =" + "'" + myUser.User_Guid + "'" + ";")

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
}

func DeleteEducation(w http.ResponseWriter, r *http.Request) {
	// Get Body
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myEducation models.Education

	json.Unmarshal(reqBody, &myEducation)
	json.NewEncoder(w).Encode(myEducation)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("UPDATE user_education set is_deleted = '1' WHERE education_guid =" + "'" + myEducation.Education_Guid + "'" + ";")

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
}

func DeleteExperience(w http.ResponseWriter, r *http.Request) {
	// Get Body
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myExperience models.Experience

	json.Unmarshal(reqBody, &myExperience)
	json.NewEncoder(w).Encode(myExperience)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("UPDATE user_experience set is_deleted = '1' WHERE experience_guid =" + "'" + myExperience.Experience_Guid + "'" + ";")

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	// Get Body
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myProject models.Project

	json.Unmarshal(reqBody, &myProject)
	json.NewEncoder(w).Encode(myProject)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("UPDATE user_projects set is_deleted = '1' WHERE project_guid =" + "'" + myProject.Project_Guid + "'" + ";")

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
}

// ------------------- Put Routes ---------------------

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get Body
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myUser models.User

	json.Unmarshal(reqBody, &myUser)
	json.NewEncoder(w).Encode(myUser)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("UPDATE user_details set email =" + "'" + myUser.Email  + "'" + ", password = " + "'" + myUser.Password  + "'" +  " WHERE user_guid =" + "'" + myUser.User_Guid + "'" + ";")

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
}

func UpdateEducation(w http.ResponseWriter, r *http.Request) {
	// Get Body
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myEducation models.Education

	json.Unmarshal(reqBody, &myEducation)
	json.NewEncoder(w).Encode(myEducation)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("UPDATE user_education set school = ? , user_from = ? , user_to = ? , degree = ? WHERE education_guid = ?;", myEducation.School, myEducation.User_From, myEducation.User_To, myEducation.Degree, myEducation.Education_Guid)

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
}

func UpdateExperience(w http.ResponseWriter, r *http.Request) {
	// Get Body
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myExperience models.Experience

	json.Unmarshal(reqBody, &myExperience)
	json.NewEncoder(w).Encode(myExperience)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("UPDATE user_experience set job_title = ? , job_description = ? , company_name = ? , user_from = ? , user_to = ? WHERE experience_guid = ? ", myExperience.Job_Title, myExperience.Job_Description, myExperience.Company_Name, myExperience.User_From, myExperience.User_To, myExperience.Experience_Guid)

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	// Get Body
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myProject models.Project

	json.Unmarshal(reqBody, &myProject)
	json.NewEncoder(w).Encode(myProject)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("UPDATE user_projects set title = ? , description = ? , technologies = ? WHERE project_guid = ?", myProject.Title, myProject.Description, myProject.Technologies, myProject.Project_Guid)

    // // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
}




