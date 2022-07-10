package main

import (
	"GinLearning/database"
	"GinLearning/internal"
	"GinLearning/middleware"
	"GinLearning/structs"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DbConfig database.Config `yaml:"database"`
}

func setupLogging() {
	f, _ := os.Create("middleware/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogging()

	config := Config{}
	b, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(b, &config)
	dc := config.DbConfig

	router := gin.Default()

	// 註冊 Validator 的 Func
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("userpasd", middleware.IsValidPassword)
		v.RegisterStructValidation(middleware.UserList, structs.Users{})
	}

	router.Use(gin.Recovery(), middleware.Logger())
	v1 := router.Group("/v1")
	internal.AddUserRouter(v1)
	go database.ConnectMySQL(dc.UserName, dc.Password, dc.Server, dc.Port, dc.DatabaseName)
	go database.ConnectMongo()

	router.Run(":8080")
}
