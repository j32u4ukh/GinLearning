package internal

import (
	"GinLearning/middleware"
	"GinLearning/service"
	"GinLearning/structs"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users", middleware.SetSession())

	user.POST("/", service.PostUser)
	user.POST("/more", service.CreateUserList)
	user.GET("/", service.CacheAllUserDecorator(service.RedisAllUser, "users", structs.User{}))
	user.GET("/:id", service.CacheOneUserDecorator(service.RedisOneUser, "id", "user_%s", structs.User{}))

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

	// Mongo
	mgo := user.Group("/mongo")
	mgo.GET("/", service.MgoGetUsers)
	mgo.GET("/:id", service.MgoGetUserById)
	mgo.POST("/", service.MgoCreateUser)
	mgo.PUT("/:id", service.MgoUpdateUser)
	mgo.DELETE("/:id", service.MgoDeleteUser)

}
