package bootstrap

import (
	"fmt"
	"ginserver/global"
	"ginserver/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// 注册分组路由
	apiGroup := router.Group("api")
	routes.SetApiGroupRoutes(apiGroup)

	return router
}

func RunServer() {
	r := setupRouter()
	if err := r.Run(fmt.Sprintf(":%s", global.AppConfig.Config.App.Port)); err != nil {
		panic(fmt.Sprintf("Run server failed, err: %+v", err))
	}
}
