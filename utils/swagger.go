package utils

import (
	"gin_demo/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitSwagger(router *gin.Engine) {
	swaggerConf := Config.Swagger
	docs.SwaggerInfo.Title = swaggerConf.Title
	docs.SwaggerInfo.Description = swaggerConf.Desc
	docs.SwaggerInfo.Version = swaggerConf.Version
	docs.SwaggerInfo.Host = swaggerConf.Host
	docs.SwaggerInfo.BasePath = swaggerConf.BasePath
	docs.SwaggerInfo.Schemes = []string{"http"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
