package main

import (
	. "GinLearning/internal"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SayHello() {
	fmt.Println("Hello Gin.")
}

func main() {
	fmt.Println("Hello Gin.")
	router := gin.Default()
	v1 := router.Group("/v1")
	AddUserRouter(v1)
	router.Run(":8080")
}
