/*
	路由表
*/
package api

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/xie1xiao1jun/go-restful/data/view/login"
)

//
var DefaultRouler = []*rest.Route{
	//管理员登录授权 相关
	Post("/login", login.M_login.OnLogin), //用户登录
	// Post("/go/verify", verify.M_verify.GetVerify),      //获取验证码
	Post("/change_pwd", login.ChangePwd), //修改密码
}

/*
	默认初始化
*/
func OnInitRoter(api *rest.Api) (router rest.App, err error) {
	OnCheckRoter(DefaultRouler...)
	router, err = rest.MakeRouter(
		DefaultRouler...,
	)

	return
}

/*
	检测重复
*/
func OnCheckRoter(routes ...*rest.Route) {
	mp := make(map[string]int)
	for _, v := range routes {
		if mp[v.PathExp] == 1 {
			panic("============" + v.PathExp + " : is repeat in rooter.go.================")
		} else {
			mp[v.PathExp] = 1
		}
	}
}
