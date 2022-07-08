package structs

import "GinLearning/database"

type User struct {
	Id       int    `json:"UserId"`
	Name     string `json:"UserName"`
	Password string `json:"UserPassword"`
	Email    string `json:"UserEmail"`
}

func GetUsers() []User {
	var users []User
	database.DB.Find(&users)
	return users
}

func GetUserById(id string) User {
	var user User
	database.DB.Where("id = ?", id).Find(&user)
	return user
}
