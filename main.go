package main

import (
	"GinLearning/database"
	. "GinLearning/internal"
	"GinLearning/middleware"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func setupLogging() {
	f, _ := os.Create("middleware/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogging()
	router := gin.Default()
	router.Use(gin.BasicAuth(gin.Accounts{"Tom": "123456"}), middleware.Logger())
	v1 := router.Group("/v1")
	AddUserRouter(v1)
	go database.Connect()
	router.Run(":8080")
}
