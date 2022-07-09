package main

import (
	"GinLearning/database"
	"GinLearning/internal"
	"GinLearning/middleware"
	"GinLearning/structs"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func setupLogging() {
	f, _ := os.Create("middleware/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogging()
	router := gin.Default()

	// 註冊 Validator 的 Func
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("userpasd", middleware.IsValidPassword)
		v.RegisterStructValidation(middleware.UserList, structs.Users{})
	}

	router.Use(gin.Recovery(), middleware.Logger())
	v1 := router.Group("/v1")
	internal.AddUserRouter(v1)
	go database.ConnectMySQL()
	go database.ConnectMongo()
	router.Run(":8080")
}
