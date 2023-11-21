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
	// 音视频模块
	video := app.Group("/api/video")
	VideoRouter(video)
	audio := app.Group("/api/audio")
	AudioRouter(audio)
	// 事务模块
	business := app.Group("/api/business")
	BusinessRouter(business)
	// 课程模块
	course := app.Group("/api/course")
	CourseRouter(course)
	// 设计模块
	design := app.Group("/api/design")
	DesignRouter(design)
	// 教务模块
	EA := app.Group("/api/ea")
	EARouter(EA)
	// 流程模块
	flow := app.Group("/api/flow")
	FlowRouter(flow)
	// 表单模块
	form := app.Group("/api/form")
	FormRouter(form)
	// 营销模块
	marketing := app.Group("/api/marketing")
	MarketingRouter(marketing)
	// 信息模块
	message := app.Group("/api/message")
	MessageRouter(message)
	// 权限模块
	rule := app.Group("/api/rule")
	RuleRouter(rule)
	// web模块
	web := app.Group("/api/web")
	WebRouter(web)
	return app
}
