package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SayHello() {
	fmt.Println("Hello Gin.")
}

func main() {
	fmt.Println("Hello Gin.")
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		// gin.H: head of gin
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/ping/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, gin.H{
			"id": id,
		})
	})
	router.Run(":8080")
}
