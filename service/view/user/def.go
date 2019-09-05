package user

//GetAllParam ...
type GetAllParam struct {
	AccessToken string `json:"access_token" gorm:"-"` //access_token
	PageNum     int    `json:"page_num"`              //第几页
}

//UserInfo 用户信息
type UserInfo struct {
	Username   string //用户名
	ExpireTime int    //超时时间
}
