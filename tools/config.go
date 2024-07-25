package tools

import (
	"com.lh.service/tools"
	"os"
	"strings"
)

func GetConfig() (tools.YamlConf, error) {
	platform := tools.Platform("")
	dir, _ := os.Getwd()
	dir = strings.Replace(dir, "\\", "/", -1)
	paths := []string{dir, "config", platform.Env + ".config.go"}
	pathname := strings.Join(paths, "/")
	configs, err := tools.Yaml(pathname)
	if err != nil {
		return tools.YamlConf{}, err
	}
	return configs, err
}
