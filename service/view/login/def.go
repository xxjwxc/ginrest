package login

import "github.com/xie1xiao1jun/public/message"

// type ThirdLoginParam struct {
// 	Type     string                  `json:"type"`      //第三方标识
// 	OpenId   string                  `json:"openid"`    //code
// 	UserInfo weixin.User_wx_info_tbl `json:"user_info"` //用户信息
// }
//
type Req_changepwd struct {
	Access_token string `json:"access_token"` //access_token
	UserName     string `json:"username"`     //用户名
	Password     string `json:"password"`     //新密码
}

//
type LoginToken_s struct {
	message.MessageBody
	Data LoginToken `json:"data,omitempty"`
}

//
type LoginToken struct {
	Access_token  string `json:"access_token"`  //access_token
	Expire_time   string `json:"expire_time"`   //
	Refresh_token string `json:"refresh_token"` //
}
