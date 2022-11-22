package models

type User struct{
	User_Guid string `json:"user_guid"`
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	City string `json:"city"`
	Country string `json:"country"`
	DOB string `json:"dob"`
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

type Job struct {
	Job_Type_Guid string `json:"job_type_guid"`
	Job_Title string `json:"job_title"`
}

type Skill struct {
	Skill_Guid string `json:"skill_guid"`
	Job_Type_Guid string `json:"job_type_guid"`
	Skill_Title string `json:"skill_title"`
}

type UserJob struct {
	User_Job_Guid string `json:"user_job_guid"`
	User_Guid string `json:"user_guid"`
	Job_Type_Guid string `json:"job_type_guid"`
	Experience int `json:"experience"`
}

type JobSkill struct {
	Job_Skill_Guid string `json:"job_skill_guid"`
	User_Job_Guid string `json:"user_job_guid"`
	Skill_Guid string `json:"skill_guid"`
	Experience int `json:"experience"`
}

type Response struct {
	Message string `json:"message"`
}
