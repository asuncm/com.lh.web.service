package router

import (
	"com.lh.service/config/proxy"
	"com.lh.web.service/tools"
	"github.com/gin-gonic/gin"
)

func AuthRouter(route *gin.RouterGroup) *gin.RouterGroup {
	config, _ := tools.GetConfig()
	services := config.Services
	auth := services["auth"]
	route.GET("/code/:name", func(c *gin.Context) {
		proxy.Proxy(c, auth)
	})
	return route
}

func CodeRouter(route *gin.RouterGroup) *gin.RouterGroup {
	config, _ := tools.GetConfig()
	services := config.Services
	auth := services["auth"]
	route.GET("/auth/:name", func(c *gin.Context) {
		proxy.Proxy(c, auth)
	})
	return route
}
