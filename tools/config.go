package tools

import (
	"com.lh.service/config/yaml"
	"com.lh.service/tools"
	"os"
	"strings"
)

func GetConfig() (yaml.YamlConf, error) {
	platform := tools.Platform("")
	dir, _ := os.Getwd()
	dir = strings.Replace(dir, "\\", "/", -1)
	paths := []string{dir, "config", platform.Env + ".config.yaml"}
	pathname := strings.Join(paths, "/")
	configs, err := yaml.Yaml(pathname)
	if err != nil {
		return yaml.YamlConf{}, err
	}
	return configs, err
}
