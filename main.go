package main

import (
	"fmt"
	"ginserver/bootstrap"
	"ginserver/global"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	bootstrap.InitializeConfig()

	global.AppConfig.Log = bootstrap.InitializeLog()
	global.AppConfig.Log.Info("log init success!")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	if err := r.Run(fmt.Sprintf(":%s", global.AppConfig.Config.App.Port)); err != nil {
		log.Fatalf("Run gin failed, err: %+v", err)
	}
}
