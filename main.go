package main

import (
	"gin_demo/router"
	"gin_demo/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	utils.ReadConfigFile("./conf/dev.json")

	utils.Connect()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL,syscall.SIGQUIT,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	router.HttpServerStop()
}
