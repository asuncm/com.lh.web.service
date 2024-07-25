package router

import (
	tools2 "com.lh.service/tools"
	"com.lh.web.service/tools"
	"github.com/gin-gonic/gin"
)

func AuthRouter(route *gin.RouterGroup) *gin.RouterGroup {
	config, _ := tools.GetConfig()
	services := config.Services
	auth := services["auth"]
	route.GET("/code/:name", func(c *gin.Context) {
		tools2.Proxy(c, auth)
	})
	return route
}

func CodeRouter(route *gin.RouterGroup) *gin.RouterGroup {
	config, _ := tools.GetConfig()
	services := config.Services
	auth := services["auth"]
	route.GET("/auth/:name", func(c *gin.Context) {
		tools2.Proxy(c, auth)
	})
	return route
}
