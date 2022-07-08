package service

import (
	"GinLearning/structs"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userList = []structs.User{}

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

func PostUser(c *gin.Context) {
	user := structs.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	userList = append(userList, user)
	c.JSON(http.StatusOK, "Successfully posted.")
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	idx := -1

	for i, user := range userList {
		log.Print(user)
		if user.Id == id {
			idx = i
			break
		}
	}

	if idx != -1 {
		userList = append(userList[:idx], userList[idx+1:]...)
		c.JSON(http.StatusOK, "Successfully deleted.")
	} else {
		c.JSON(http.StatusNotFound, "Error")
	}
}

func PutUser(c *gin.Context) {
	origin := structs.User{}
	err := c.BindJSON(&origin)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	for i, user := range userList {
		log.Print(user)
		if user.Id == id {
			userList[i] = origin
			log.Print(userList[i])
			c.JSON(http.StatusOK, "Successfully.")
			return
		}
	}
	c.JSON(http.StatusNotFound, "Error")
}
