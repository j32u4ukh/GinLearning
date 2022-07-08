package main

import (
	"GinLearning/database"
	. "GinLearning/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	AddUserRouter(v1)
	go database.Connect()
	router.Run(":8080")
}
