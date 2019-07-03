package user

//
type GetAllParam struct {
	Access_token string `json:"access_token" gorm:"-"` //access_token
	Page_num     int    `json:"page_num"`              //第几页
}

type UserInfo struct {
	Username   string //用户名
	ExpireTime int    //超时时间
}
