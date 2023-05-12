package api

import (
	"gin-demo/api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/register", register) // 注册

	r.POST("/login", login) // 登录

	r.POST("/login/changepassword", ChangePassword)

	r.POST("/find", FindPassword) //找回密码

	r.POST("/comments", Comments)

	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}

	r.Run(":8080")
}
