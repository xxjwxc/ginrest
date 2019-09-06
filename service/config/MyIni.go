package config

//Config 扩展配置
type Config struct {
	CfgBase
	CupsCmd string `json:"cups_cmd" toml:"cups_cmd"` //cups命令
}

//GetCups cups命令
func GetCups() string {
	return _map.CupsCmd
}
