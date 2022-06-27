package main

import (
	"flag"
	"gin_demo/pkg/conf"
	"gin_demo/pkg/gorm"
	rds "gin_demo/pkg/redis"
	"gin_demo/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	configFile := flag.String("config", "./conf/dev.json", "Config file path")
	if err := conf.Init(*configFile); err != nil {
		log.Panicf("load config file failed:%v", err)
	}
	if err := gorm.CheckStatus(); err != nil {
		log.Panicf("check mysql failed. error:%v", err)
	}

	if err := rds.CheckStatus(); err != nil {
		log.Panicf("check redis failed. error:%v", err)
	}
	r := gin.Default()
	gin.SetMode(conf.GetString("mode"))
	router.InitRouter(r)
	r.Run(conf.GetString("addr"))
}
