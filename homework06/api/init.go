package api

import (
	"awesomeProject1/homework06/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterGin() {
	r := gin.Default()

	// 公开路由
	r.POST("/register", Register)
	r.POST("/login", Login)
	r.POST("/refresh", RefreshToken)

	// 需要认证的路由
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.POST("/modify_password", ModifyPassword)
		auth.GET("/ping", Ping1)
	}

	r.Run(":8080")
}

func Ping1(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
