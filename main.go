package main

import (
	"com.lh.basic/config"
	"com.lh.service/pgx"
	"com.lh.service/tools"
	"com.lh.web.service/locales"
	"com.lh.web.service/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	config.InitConfig("com.lh.web.service")
	configs := config.GetConfig("webService")
	pgx.InitConfig()
	app.Use(tools.Cors())
	//app.Use(tools.MiddleWare(configs))
	locales.Init()
	router.Router(app)
	app.Run(fmt.Sprintf("%s:%s", configs.Host, configs.Port))
}
