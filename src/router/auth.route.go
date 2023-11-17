package router

import (
	"com.lh.service/src/config/proxy"
	"com.lh.web.service/src/tools"
	"github.com/gin-gonic/gin"
)

func AuthRouter(route *gin.RouterGroup) *gin.RouterGroup {
	config, _ := tools.GetConfig()
	services := config.Services
	auth := services["auth"]
	route.Any("/code/:name", func(c *gin.Context) {
		proxy.Proxy(c, auth)
	})
	return route
}

func CodeRouter(route *gin.RouterGroup) *gin.RouterGroup {
	config, _ := tools.GetConfig()
	services := config.Services
	auth := services["auth"]
	route.Any("/code/:name", func(c *gin.Context) {
		proxy.Proxy(c, auth)
	})
	return route
}
