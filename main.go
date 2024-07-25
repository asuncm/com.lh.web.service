package main

import (
	"com.lh.service/tools"
	"com.lh.web.service/locales"
	router2 "com.lh.web.service/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func Config() (tools.MiddleConf, error) {
	platform := tools.Platform("")
	root := tools.GetPath("LHPATH", "")
	pathname := fmt.Sprintf("%s%s%s%s", root, "/config/", platform.Env, ".config.yaml")
	configs, err := tools.Yaml(pathname)
	if err != nil {
		return tools.MiddleConf{}, err
	}
	devServe := configs.Services["webService"]
	database := tools.GetPath(configs.Database, "pebble/webService")
	return tools.MiddleConf{
		Platform: platform.Platform,
		Serve:    fmt.Sprintf("%s%s", root, "/com.lh.web.service"),
		Root:     root,
		Host:     devServe.Host,
		Port:     devServe.Port,
		DataDir:  database,
	}, err
}

func main() {
	router := gin.Default()
	configs, _ := Config()
	router.Use(tools.Cors())
	locales.Init()
	router2.Router(router)
	address := []string{configs.Host, configs.Port}
	router.Run(strings.Join(address, ":"))
}
