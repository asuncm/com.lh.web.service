package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func EARouter(route *gin.RouterGroup) *gin.RouterGroup {
	route.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "asaaa===")
	})
	return route
}
