package service

import (
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
