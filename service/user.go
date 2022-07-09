package service

import (
	"GinLearning/middleware"
	"GinLearning/structs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	user := structs.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	newUser := structs.CreateUser(user)
	c.JSON(http.StatusOK, newUser)
}

func CreateUserList(c *gin.Context) {
	users := structs.Users{}
	err := c.BindJSON(&users)
	if err != nil {
		// c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		c.String(400, "Error:%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUsers(c *gin.Context) {
	// c.JSON(http.StatusOK, users)
	users := structs.GetUsers()
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	user := structs.GetUserById(c.Param("id"))

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
	} else {
		log.Println("User ->", user)
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	user := structs.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	user = structs.UpdateUser(c.Param("id"), user)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	result := structs.DeleteUser(c.Param("id"))

	if result {
		c.JSON(http.StatusOK, "Successfully deleted.")
	} else {
		c.JSON(http.StatusNotFound, "Error")
	}
}

func LoginUser(c *gin.Context) {
	// NOTE: Postman 需使用 x-www-form-urlencoded，PostForm 才有辦法收到數據
	name := c.PostForm("name")
	password := c.PostForm("password")
	log.Printf("name: %s\n", name)
	log.Printf("password: %s\n", password)
	user := structs.CheckUserPassword(name, password)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
	} else {
		middleware.SaveSession(c, user.Id)
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successfully",
			"User":    user,
			"Session": middleware.GetSessionId(c),
		})
	}
}

// Logout user
func LogoutUser(c *gin.Context) {
	middleware.ClearSession(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successfully",
	})
}

// Check user session
func CheckUserSession(c *gin.Context) {
	sessionId := middleware.GetSessionId(c)
	if sessionId == 0 {
		c.JSON(http.StatusUnauthorized, "Error")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Check Session successfully",
			"User":    middleware.GetSessionId(c),
		})
	}
}

// Redis user
func RedisOneUser(c *gin.Context) {
	id := c.Param("id")
	if id == "0" {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	user := structs.GetUserById(id)
	c.Set("dbResult", user)
}

// Redis all users
func RedisAllUser(c *gin.Context) {
	users := structs.GetUsers()
	c.Set("dbAllUser", users)
}
