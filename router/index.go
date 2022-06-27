package router

import (
	"gin_demo/middleware"
	"gin_demo/pkg/swagger"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.Use(middleware.RecoveryMiddleware())
	router.Use(middleware.Cors())
	initUserRouter(router)
	swagger.InitSwagger(router)
}
