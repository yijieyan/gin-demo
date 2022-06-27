package router

import (
	"gin_demo/controller"
	"gin_demo/middleware"
	"github.com/gin-gonic/gin"
)

func initUserRouter(router *gin.Engine) {
	user := router.Group("/user")

	userControl := controller.UserController{}
	user.POST("/register", userControl.Register)
	user.POST("/login", userControl.Login)
	user.GET("/getUserInfo", middleware.JWTAuthMiddleware(), userControl.GetUserInfo)
}
