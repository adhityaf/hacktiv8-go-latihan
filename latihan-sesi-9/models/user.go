package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users []User

func GetUsers() *[]User {
	return &Users
}

func AddNewUser(user *User) {
	Users = append(Users, *user)
}
