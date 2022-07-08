package service

import (
	"GinLearning/structs"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users = []structs.User{}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func PostUser(c *gin.Context) {
	user := structs.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	users = append(users, user)
	c.JSON(http.StatusOK, "Successfully posted.")
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	idx := -1

	for i, user := range users {
		log.Print(user)
		if user.Id == id {
			idx = i
			break
		}
	}

	if idx != -1 {
		users = append(users[:idx], users[idx+1:]...)
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
	for i, user := range users {
		log.Print(user)
		if user.Id == id {
			users[i] = origin
			log.Print(users[i])
			c.JSON(http.StatusOK, "Successfully.")
			return
		}
	}
	c.JSON(http.StatusNotFound, "Error")
}
