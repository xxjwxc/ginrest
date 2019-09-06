package config

import (
	"os"
	"strings"
)

const (
	test_file = `
	serial_number = "1.0"  #版本号
	service_name = "xxj"   #服务名
	service_displayname = "xxj name" #服务显示名
	sercice_desc = "xxj desc"    #服务描述
	is_dev = true  # 是否开发者模式
	http_port = "7011" # 监听端口号
	leveldb_dir = "database" # leveldb数据库地址
	lg = "cn" #语言
	domain_name = "http://www.xxjwxc.cn:6001/dzmhotpot/api/v1" #服务器地址	
`
)

//判断是否在测试环境下使用
func IsRunTesting() bool {
	if len(os.Args) > 1 {
		return strings.HasPrefix(os.Args[1], "-test")
	}
	return false
}
