package config

import (
	"os"
	"strings"
)

//Config 扩展配置
type Config struct {
	CfgBase
	CupsCmd string `json:"cups_cmd" toml:"cups_cmd"` //cups命令
}

//GetCups cups命令
func GetCups() string {
	return _map.CupsCmd
}

//IsRunTesting 判断是否在测试环境下使用
func IsRunTesting() bool {
	if len(os.Args) > 1 {
		return strings.HasPrefix(os.Args[1], "-test")
	}
	return false
}
