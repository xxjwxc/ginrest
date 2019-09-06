package config

import (
	"fmt"

	"github.com/xxjwxc/public/mylog"

	"github.com/BurntSushi/toml"
	"github.com/xxjwxc/public/dev"
	"github.com/xxjwxc/public/tools"
)

//CfgBase 基础配置
type CfgBase struct {
	SerialNumber       string `json:"serial_number" toml:"serial_number"`             //版本号
	ServiceName        string `json:"service_name" toml:"service_name"`               //service名字
	ServiceDisplayname string `json:"service_displayname" toml:"service_displayname"` //显示名
	SerciceDesc        string `json:"sercice_desc" toml:"sercice_desc"`               //service描述
	IsDev              bool   `json:"is_dev" toml:"is_dev"`                           //是否是开发版本
	DomainName         string `json:"domain_name" toml:"domain_name"`                 //域名
	HTTPPort           string `json:"http_port" toml:"http_port"`                     //端口号
	LeveldbDir         string `json:"leveldb_dir" toml:"leveldb_dir"`                 //leveldb数据库地址
	Oauth2URL          string `json:"oauth2_url" toml:"oauth2_url"`                   //oauth授权地址入口
}

const (
	//FileHost 文件地址
	FileHost = "file"
)

//StaticHost 静态文件存放地址
var StaticHost = [2]string{"/static", ""}

var _map = Config{}

func init() {
	onInit()

	//配置公共配置到public
	dev.OnSetDev(OnIsDev())
	dev.SetService(_map.ServiceName)
	dev.SetFileHost(FileHost)
}

func onInit() {
	path := tools.GetModelPath()
	err := initFile(path + "/config.toml")
	if err != nil {
		fmt.Println("initFile error : ", err.Error())
		mylog.Fatal(err)
		return
	}
}

//OnIsDev 是否是开发版本
func OnIsDev() bool {
	return _map.IsDev
}

//GetServerPort 获取端口号
func GetServerPort() (strPort string) {
	strPort = _map.HTTPPort
	return
}

//GetLevelDbDir 获取Leveldb的文件地址
func GetLevelDbDir() string {
	return tools.GetModelPath() + "/" + _map.LeveldbDir
}

//初始化文件
func initFile(filename string) error {

	if IsRunTesting() {
		if _, err := toml.Decode(test_file, &_map); err != nil {
			fmt.Println("read toml error: ", err.Error())
			return err
		}

		return nil
	}

	if _, err := toml.DecodeFile(filename, &_map); err != nil {
		fmt.Println("read toml error: ", err.Error())
		return err
	}

	return nil
}

//GetServiceName 服务名字
func GetServiceName() string {
	return _map.ServiceName
}

//GetServiceConfig 获取service配置信息
func GetServiceConfig() (name, displayName, desc string) {
	name = _map.ServiceName
	displayName = _map.ServiceDisplayname
	desc = _map.SerciceDesc
	return
}

//GetDomainName 获取服务器地址
func GetDomainName() string {
	return _map.DomainName
}

//GetLoginURL 登录地址
func GetLoginURL() string {
	return _map.Oauth2URL + "/authorize"
}

//GetLoginNoPwdURL 密码登录
func GetLoginNoPwdURL() string {
	return _map.Oauth2URL + "/authorize_nopwd"
}

//GetCheckTokenURL token 检测
func GetCheckTokenURL() string {
	return _map.Oauth2URL + "/check_token"
}

//GetRefreshTokenURL token刷新
func GetRefreshTokenURL() string {
	return _map.Oauth2URL + "/refresh_token"
}
