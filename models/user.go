package models

type User struct{
	User_Guid string `json:"user_guid"`
	Email string `json:"email"`
	Password string `json:"password"`
	Secret_Key string `json:"secret_key"`
}

type Education struct {
	Education_Guid string `json:"education_guid"`
	User_Guid  string `json:"user_guid"`
	School string `json:"school"`
	User_From string `json:"user_from"`
	User_To string `json:"user_to"`
	Degree string `json:"degree"`
}

type Experience struct {
	Experience_Guid string `json:"experience_guid"`
	User_Guid  string `json:"user_guid"`
	Job_Title string `json:"job_title"`
	Job_Description string `json:"job_description"`
	Company_Name string `json:"company_name"`
	User_To string `json:"user_to"`
	User_From string `json:"user_from"`
}

type Project struct {
	Project_Guid string `json:"project_guid"`
	User_Guid  string `json:"user_guid"`
	Title string `json:"title"`
	Description string `json:"description"`
	Technologies string `json:"technologies"`
}