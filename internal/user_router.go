package internal

import (
	"GinLearning/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")
	user.POST("/", service.PostUser)
	user.POST("/more", service.CreateUserList)
	user.GET("/", service.GetUsers)
	user.GET("/:id", service.GetUserById)
	user.DELETE("/:id", service.DeleteUser)
	user.PUT("/:id", service.UpdateUser)
}
