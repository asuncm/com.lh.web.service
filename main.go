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
<<<<<<< HEAD
	root := tools.GetPath("LHPATH", "")
	pathname := fmt.Sprintf("%s%s%s%s", root, "/config/", platform.Env, ".config.yaml")
=======
	pathname := tools.GetPath("LHPATH", fmt.Sprintf("%s%s%s", "config/", platform.Env, ".config.yaml"))
>>>>>>> d24e334 (更新：代码更新)
	configs, err := tools.Yaml(pathname)
	if err != nil {
		return tools.MiddleConf{}, err
	}
	devServe := configs.Services["webService"]
<<<<<<< HEAD
	database := tools.GetPath(configs.Database, "pebble/webService")
	return tools.MiddleConf{
		Platform: platform.Platform,
		Serve:    fmt.Sprintf("%s%s", root, "/com.lh.web.service"),
		Root:     root,
		Host:     devServe.Host,
		Port:     devServe.Port,
		DataDir:  database,
=======
	root := configs.Root
	database := fmt.Sprintf("%s%s", configs.Database, "/pebble")
	return tools.MiddleConf{
		Platform:  platform.Platform,
		Serve:     "webService",
		Root:      root,
		Host:      devServe.Host,
		Port:      devServe.Port,
		DataCache: database,
		DataPort:  devServe.DataPort,
>>>>>>> d24e334 (更新：代码更新)
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
