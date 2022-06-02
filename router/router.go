package router

import (
	"gin_demo/controller"
	"gin_demo/middleware"
	"gin_demo/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.RecoveryMiddleware())

	utils.InitSwagger(router)
	user := router.Group("/user")

	userControl := controller.UserController{}
	user.POST("/register", userControl.Register)
	user.POST("/login", userControl.Login)
	user.GET("/getUserInfo",middleware.JWTAuthMiddleware(), userControl.GetUserInfo)


	return router
}
