package swagger

import (
	"gin_demo/docs"
	"gin_demo/pkg/conf"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitSwagger(router *gin.Engine) {
	docs.SwaggerInfo.Title = conf.GetString("swagger.title")
	docs.SwaggerInfo.Description = conf.GetString("swagger.desc")
	docs.SwaggerInfo.Version = conf.GetString("swagger.version")
	docs.SwaggerInfo.Host = conf.GetString("swagger.host")
	docs.SwaggerInfo.BasePath = conf.GetString("swagger.base_path")
	docs.SwaggerInfo.Schemes = []string{"http"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
