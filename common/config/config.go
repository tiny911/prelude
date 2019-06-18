package config

import "github.com/tiny911/doraemon/config"

// Cfg 配置变量，统一从YAML文件解析
var Cfg = struct {
	FrontPort int `yaml:"FrontPort"`
}{}

func init() {
	config.Parse(&Cfg)
}
