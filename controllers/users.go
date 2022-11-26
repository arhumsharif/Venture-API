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
	"github.com/dgrijalva/jwt-go"
    "database/sql"
	"time"
)

// ------------------------------------Authentication --------------------------------

// Authenticate
var jwtKey = []byte("secret_key")
func Authenticate(w http.ResponseWriter, r *http.Request) {
	// make an object of credentials
	var credentials models.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// run query to check for credentials
	DB := db.ConnectDB()
	rows, queryerr:= DB.Query("SELECT user_guid, email, password, secret_key FROM user_details WHERE email=? AND password=?",credentials.Email, credentials.Password)
	if queryerr != nil {
		fmt.Println("Error:", queryerr)
	}

	var myuser models.User
	for rows.Next() {

        err = rows.Scan(&myuser.User_Guid, &myuser.Email, &myuser.Password, &myuser.Secret_Key)
        if err != nil {
            fmt.Println("err:", err) // proper error handling instead of panic in your app
        }
    }
	if myuser.User_Guid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Secret Key
	myKey := utils.GenerateSecretKey()
	keyJWT := []byte(myKey)

	// Data is in myuser
	expirationTime := time.Now().Add(time.Minute * 60)

	claims := &models.Claims{
		Username: myuser.Email,
		User_Guid: myuser.User_Guid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, tokenerr := token.SignedString(keyJWT)

	if tokenerr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Update the Query to store secret_key
	update, updateerr := DB.Query("UPDATE user_details set secret_key = ? WHERE user_guid = ? ", myKey, myuser.User_Guid)

    // // if there is an error inserting, handle it
    if updateerr != nil {
        fmt.Println("Error:", updateerr)
    }
	defer update.Close()


	fmt.Println(tokenString)
	tokenString = tokenString + " " + myuser.User_Guid
	// Send Response
	var response models.Response
	response.Message = tokenString
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}

	defer DB.Close()
	w.Write(jsonResponse)
	// verify(tokenString, w, r)
}




// ---------------------------------------------------------------------

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
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Success"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}

func InsertEducation(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)
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
	insert, err := DB.Query("INSERT INTO user_education (education_guid, user_guid, school, user_from, user_to, degree) VALUES (" + "'" +  education_id + "'" + "," + "'" + userGuid + "'" + "," + "'" + myEducation.School + "'" + "," + "'" + myEducation.User_From + "'"  + "," + "'" + myEducation.User_To + "'" + "," + "'" + myEducation.Degree + "'" + ");")

    // // if there is an error inserting, handle it
    if err != nil {
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Success"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}

func InsertExperience(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)
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
	insert, err := DB.Query("INSERT INTO user_experience (experience_guid, user_guid, job_title, job_description, company_name, user_from, user_to) VALUES (" + "'" +  experience_id + "'" + "," + "'" + userGuid + "'" + "," + "'" + myExperience.Job_Title + "'" + "," + "'" + myExperience.Job_Description + "'"  + "," + "'" + myExperience.Company_Name + "'" + "," + "'" + myExperience.User_From + "'" + "," + "'" + myExperience.User_To + "'" + ");")

    // // if there is an error inserting, handle it
    if err != nil {
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Success"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}

func InsertProject(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)
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
	insert, err := DB.Query("INSERT INTO user_projects (project_guid, user_guid, title, description, technologies) VALUES (" + "'" +  project_id + "'" + "," + "'" + userGuid + "'" + "," + "'" + myProject.Title + "'" + "," + "'" + myProject.Description + "'"  + "," + "'" + myProject.Technologies + "'"  + ");")

    // // if there is an error inserting, handle it
    if err != nil {
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Success"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}


func InsertJob(w http.ResponseWriter, r *http.Request) {
	// Get Body
	guid := uuid.New() // job guid
	job_id := guid.String()	
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myJob models.Job

	json.Unmarshal(reqBody, &myJob)
	json.NewEncoder(w).Encode(myJob)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("INSERT INTO job_type (job_type_guid, job_title) VALUES (?, ?)", job_id, myJob.Job_Title)

    // // if there is an error inserting, handle it
    if err != nil {
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Success"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}


func InsertSkill(w http.ResponseWriter, r *http.Request) {
	// Get Body
	guid := uuid.New() // job guid
	skill_id := guid.String()	
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var mySkill models.Skill

	json.Unmarshal(reqBody, &mySkill)
	json.NewEncoder(w).Encode(mySkill)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("INSERT INTO skills (skill_guid, job_type_guid, skill_title) VALUES (?, ?, ?)", skill_id, mySkill.Job_Type_Guid, mySkill.Skill_Title)

    // // if there is an error inserting, handle it
    if err != nil {
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Success"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}


func InsertUserJob(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)
	// Get Body
	guid := uuid.New() // job guid
	user_job_id := guid.String()	
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myUserJob models.UserJob

	json.Unmarshal(reqBody, &myUserJob)
	json.NewEncoder(w).Encode(myUserJob)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("INSERT INTO user_job (user_job_guid, user_guid, job_type_guid, experience) VALUES (?, ?, ?, ?)", user_job_id, userGuid, myUserJob.Job_Type_Guid, myUserJob.Experience)

    // // if there is an error inserting, handle it
    if err != nil {
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Success"
	response.Id = user_job_id
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}


func InsertJobSkill(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)
	// Get Body
	guid := uuid.New() // job guid
	job_skill_id := guid.String()	
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myJobSkill models.JobSkill

	json.Unmarshal(reqBody, &myJobSkill)
	json.NewEncoder(w).Encode(myJobSkill)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("INSERT INTO job_skill (job_skill_guid, user_job_guid, skill_guid, experience) VALUES (?, ?, ?, ?)", job_skill_id, myJobSkill.User_Job_Guid, myJobSkill.Skill_Guid, myJobSkill.Experience)

    // // if there is an error inserting, handle it
    if err != nil {
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Success"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}


// ----------------------------Get Routes------------------------------



func GetUsers(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)
	// get data against guid
	DB := db.ConnectDB()
	rows, err:= DB.Query("SELECT * FROM user_details WHERE user_guid=?",userGuid)
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
	fmt.Println("In Experience")
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)
	DB := db.ConnectDB()
	rows, err:= DB.Query("SELECT * FROM user_experience WHERE user_guid=? AND is_deleted = '0'",userGuid)
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
	fmt.Println("In Education")
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)
	DB := db.ConnectDB()
	rows, err:= DB.Query("SELECT * FROM user_education WHERE user_guid=? AND is_deleted = '0'",userGuid)
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
	fmt.Println("In Projects")
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)
	DB := db.ConnectDB()
	rows, err:= DB.Query("SELECT * FROM user_projects WHERE user_guid=? AND is_deleted = '0'",userGuid)
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

func GetJobs(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)

	DB := db.ConnectDB()
	rows, err:= DB.Query("SELECT * FROM job_type")
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

func GetSkills(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)

	DB := db.ConnectDB()
	rows, err:= DB.Query("SELECT * FROM skills")
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

func GetUserJobs(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)

	vars := mux.Vars(r)
	guid := vars["guid"]
	fmt.Println("guid", guid)
	// get data against guid
	DB := db.ConnectDB()
	rows, err:= DB.Query("SELECT * FROM user_job WHERE user_guid=?",guid)
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

func GetJobSkills(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)

	vars := mux.Vars(r)
	guid := vars["guid"]
	fmt.Println("guid", guid)
	// get data against guid
	DB := db.ConnectDB()
	rows, err:= DB.Query("SELECT * FROM job_skill WHERE user_job_guid=?",guid)
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
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)
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
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Deleted"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}

func DeleteEducation(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)

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
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Deleted"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}

func DeleteExperience(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)

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
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Deleted"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)

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
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Deleted"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}

// ------------------- Put Routes ---------------------

func UpdateUserDetail(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)

	// Get Body
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myUser models.User

	json.Unmarshal(reqBody, &myUser)
	json.NewEncoder(w).Encode(myUser)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("UPDATE user_details set name = ?, phone = ?, address = ?, city = ?, country = ?, dob = ? WHERE user_guid = ? ", myUser.Name, myUser.Phone, myUser.Address, myUser.City, myUser.Country, myUser.DOB ,userGuid)

    // // if there is an error inserting, handle it
    if err != nil {
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Updated"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get Body
	var DB *sql.DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var myUser models.User

	json.Unmarshal(reqBody, &myUser)
	json.NewEncoder(w).Encode(myUser)

	// Perform Query
	DB = db.ConnectDB()
	insert, err := DB.Query("UPDATE user_details set email = ?, password = ? WHERE user_guid = ? ", myUser.Email, myUser.Password, myUser.User_Guid)

    // // if there is an error inserting, handle it
    if err != nil {
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Updated"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}

func UpdateEducation(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)

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
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Updated"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}

func UpdateExperience(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)
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
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Updated"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	// get headers
	authorization := r.Header.Get("Authorization")
	// Roles Defined
	rolesToCheck := [] string {"user", "admin"}
	// -------------
	userGuid := ""
	role := ""
	// CheckAuth 
	userGuid, role, status := utils.CheckAuth(authorization, w, r)
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		roleStatus := utils.ValidateRole(role, rolesToCheck)
		
		if !roleStatus {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Println("User Guid in Main: ", userGuid)

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
        fmt.Println("Error:", err)
    }
    // be careful deferring Queries if you are using transactions
	// Send Response
	var response models.Response
	response.Message = "Updated"
	var jsonResponse []byte
	jsonResponse, resErr := json.Marshal(response)

	if resErr != nil {
		fmt.Println("Error:", resErr)
	}
	w.Write(jsonResponse)
    defer insert.Close()
}




