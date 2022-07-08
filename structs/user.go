package structs

import (
	"GinLearning/database"
	"log"
)

type User struct {
	Id       int    `json:"UserId"`
	Name     string `json:"UserName"`
	Password string `json:"UserPassword"`
	Email    string `json:"UserEmail"`
}

func CreateUser(user User) User {
	database.DB.Create(&user)
	return user
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

func UpdateUser(id string, user User) User {
	database.DB.Model(&user).Where("id = ?", id).Updates(user)
	return user
}

func DeleteUser(id string) bool {
	var user User
	result := database.DB.Where("id = ?", id).Delete(&user).RowsAffected
	log.Println("result:", result)
	return result != 0
}
