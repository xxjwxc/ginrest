package view

//基础类型（accesstoken）
type BaseAccessToken struct {
	Access_token string `json:"access_token"` //用户登陆凭证
}

//openid
type BaseOpenidBody struct {
	Openid string `gorm:"unique_index" json:"openid"` //微信用户唯一标识符
}
