package router

import (
	"context"
	"gin_demo/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var HttpSrvHandle *http.Server

func HttpServerRun() {
	config := utils.Config
	gin.SetMode(config.DebugMode)
	r := InitRouter()

	HttpSrvHandle = &http.Server{
		Addr: config.Addr,
		Handler: r,
		ReadTimeout: time.Duration(config.ReadTimeout)  * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << config.MaxHeaderBytes,
	}

	go func() {
		log.Printf("[INFO] HttpServerRun:%s\n",config.Addr)
		if err := HttpSrvHandle.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n",config.Addr,err)
		}
	}()
}

func HttpServerStop() {
	ctx,cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	if err := HttpSrvHandle.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n",err)
	} else {
		log.Printf("[INFO] HttpServerStop stopped\n")
	}
}
