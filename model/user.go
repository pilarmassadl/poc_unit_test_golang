package model

type User struct {
	Document string `json:"document"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Age int `json:"age"`
}

type UserList []User
