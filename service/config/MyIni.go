package config

type Config struct {
	Cfg_Base
	Wx_AppID           string `json:"wx_appid"`
	Wx_apiKey          string `json:"wx_apiKey"`
	Wx_mchId           string `json:"wx_mchId"`
	Wx_AppSecret       string `json:"wx_secret"`
	Wx_Token           string `json:"wx_token"`
	Wx_EncodingAESKey  string `json:"wx_encodingAESKey"`
	Wx_NotifyUrl       string `json:"wx_notifyUrl"`
	Wx_authorize_url   string `json:"wx_authorize_url"`
	Sign_api_key       string `json:"sign_api_key,omitempty"` //验签key
	Sms_AppID          string `json:"sms_appid"`
	Sms_Secret         string `json:"sms_secret"`
	Paper_qrcode_url   string `json:"paper_qrcode_url"`
	Register_app_id    string `json:"register_app_id"`
	Getui_AppID        string `json:"getui_appid"`
	Getui_AppKey       string `json:"getui_appkey"`
	Getui_MasterSecret string `json:"getui_mastersecret"`
	Baidu_Key          string `json:"baidu_key"`
	Baidu_Secret       string `json:"baidu_secret"`
	Cos_bucket         string `json:"cos_bucket"`
	Cos_appid          string `json:"cos_appid"`
	Cos_secretid       string `json:"cos_secretid"`
	Cos_secretkey      string `json:"cos_secretkey"`
	Cos_region         string `json:"cos_region"`
	Cos_domain         string `json:"cos_domain"`     //cos域名
	Content_count      int    `json:"content_count"`  //每页行数
	Qrcode_wx_url      string `json:"qrcode_wx_url"`  //小程序码地址
	Qrcode_url         string `json:"qrcode_url"`     //小程序二维码地址
	Top_count          int    `json:"top_count"`      //销售排行前几
	Qrcode_width       int    `json:"qrcode_width"`   //小程序二维码宽度
	Wx_public_name     string `json:"wx_public_name"` //公众号名
	Cmd                string `json:"cmd"`            //cmd
}

type WxInfo struct {
	AppID          string
	AppSecret      string
	MchId          string
	Key            string
	Token          string
	EncodingAESKey string
	NotifyUrl      string
}
type SmsInfo struct {
	AppID     string
	AppSecret string
}
type GeTuiInfo struct {
	AppID        string
	AppKey       string
	MasterSecret string
}

type CosInfo struct {
	Bucket    string
	AppID     string
	SecretID  string
	Secretkey string
	Region    string
	Domain    string
}

func GetCosInfo() CosInfo {
	return CosInfo{
		Bucket:    _map.Cos_bucket,
		AppID:     _map.Cos_appid,
		SecretID:  _map.Cos_secretid,
		Secretkey: _map.Cos_secretkey,
		Region:    _map.Cos_region,
		Domain:    _map.Cos_domain}
}

func GetWxInfo() WxInfo {
	return WxInfo{_map.Wx_AppID,
		_map.Wx_AppSecret,
		_map.Wx_mchId,
		_map.Wx_apiKey,
		_map.Wx_Token,
		_map.Wx_EncodingAESKey,
		_map.Wx_NotifyUrl}
}
func GetSmsInfo() SmsInfo {
	return SmsInfo{_map.Sms_AppID, _map.Sms_Secret}
}

func GetGeTuiInfo() GeTuiInfo {
	return GeTuiInfo{_map.Getui_AppID,
		_map.Getui_AppKey,
		_map.Getui_MasterSecret}
}

type ApiServerUrl struct {
	GetUserInfo string
}

func GetSignKey() string {
	return _map.Sign_api_key
}

func GetPaperQrcodeUrl() string {
	return _map.Paper_qrcode_url
}

func GetRegisterId() string {
	return _map.Register_app_id
}

//百度身份证识别id
func GetBaiduInfo() (key, secret string) {
	key = _map.Baidu_Key
	secret = _map.Baidu_Secret
	return
}

func GetExternal_iPAddr() string {
	return ""
}

func GetContent_count() int {
	return _map.Content_count
}

func GetQrcode_wx_url() string { //小程序码地址
	return _map.Qrcode_wx_url
}

func GetQrcode_url() string { //小程序二维码地址
	return _map.Qrcode_url
}

func GetQrcode_width() int { //小程序二维码宽度
	return _map.Qrcode_width
}
func GetTop_count() int {
	return _map.Top_count
}

func GetWx_public_name() string {
	return _map.Wx_public_name
}

func GetCmd() string {
	return _map.Cmd
}
