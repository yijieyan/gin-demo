package router

import (
	"gin_demo/middleware"
	"gin_demo/pkg/swagger"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	//中间件
	router.Use(middleware.RecoveryMiddleware(),middleware.Cors())

	//注册路由分组模块
	initUserRouter(router)

	//swagger文档初始化
	swagger.InitSwagger(router)
}
