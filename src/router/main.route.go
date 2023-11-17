package router

import (
	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine) *gin.Engine {
	// 授权模块
	auth := app.Group("/api/auth")
	AuthRouter(auth)
	// 验证码模块
	code := app.Group("/api/code")
	CodeRouter(code)
	// 基础模块
	basic := app.Group("/api/basic")
	BasicRouter(basic)
	// 用户模块
	user := app.Group("/api/user")
	UserRouter(user)
	return app
}
