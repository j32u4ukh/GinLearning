package internal

import (
	"GinLearning/middleware"
	"GinLearning/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users", middleware.SetSession())

	user.POST("/", service.PostUser)
	user.POST("/more", service.CreateUserList)
	user.GET("/", service.GetUsers)
	user.GET("/:id", service.GetUserById)

	user.PUT("/:id", service.UpdateUser)

	// Login
	user.POST("/login", service.LoginUser)

	// Check user session
	user.GET("/check", service.CheckUserSession)

	// 要求下列服務需要有 Session 才能使用(登入後會產生 Session)
	user.Use(middleware.AuthSession())
	{
		user.DELETE("/:id", service.DeleteUser)
		user.GET("/logout", service.LogoutUser)
	}
}
