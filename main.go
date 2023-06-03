package main

import (
	"ginserver/bootstrap"
	"ginserver/global"
)

func main() {
	bootstrap.InitializeConfig()

	global.AppConfig.Log = bootstrap.InitializeLog()
	global.AppConfig.Log.Info("log init success!")

	bootstrap.RunServer()
}
