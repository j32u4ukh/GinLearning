package structs

import (
	"GinLearning/database"
	"log"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

type Users struct {
	UserList     []User `json:"UserList" binding:"required,gt=0,lt=3"`
	UserListSize int    `json:"UserListSize"`
}

// DB 命名要用 Users，似乎是 gorm 的限制
type User struct {
	// struct Id -> gorm db: id; struct UserId -> gorm db: user_id
	// binding:"required" 一定要輸入
	Id       int    `json:"UserId" binding:"required"`
	Name     string `json:"UserName" binding:"required,gt=5"`
	Password string `json:"UserPassword" binding:"min=4,max=20,userpasd"`
	Email    string `json:"UserEmail" binding:"email"`
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
	database.DB.Where("id = ?", id).First(&user)
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
	return result > 0
}

func CheckUserPassword(name string, password string) User {
	user := User{}
	database.DB.Where("name = ? and password = ?", name, password).First(&user)
	return user
}

////////////////////////////////////////////////////////////////
// Mongo
////////////////////////////////////////////////////////////////
func MgoCreateUser(user User) User {
	database.Mgo.Insert(user)
	return user
}

func MgoGetUsers() []User {
	var users []User
	database.Mgo.Find(nil).All(&users)
	return users
}

func MgoGetUserById(id string) User {
	userId, _ := strconv.Atoi(id)
	var user User
	database.Mgo.Find(bson.M{"id": userId}).One(&user)
	return user
}

func MgoUpdateUser(id string, user User) User {
	userId, _ := strconv.Atoi(id)
	updateUserId := bson.M{"id": userId}
	updateUserData := bson.M{"$set": user}
	err := database.Mgo.Update(updateUserId, updateUserData)
	if err != nil {
		log.Println(err.Error())
		return User{}
	}
	return user
}

func MgoDeleteUser(id string) bool {
	userId, _ := strconv.Atoi(id)
	err := database.Mgo.Remove(bson.M{"id": userId})
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
