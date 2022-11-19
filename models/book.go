package models

type User struct{
	User_Guid string `json:"user_guid"`
	Email string `json:"email"`
	Password string `json:"password"`
	Secret_Key string `json:"secret_key"`
}