package login

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/tools"
)

//
var M_login Login

//
type Login struct {
}

/*
	登录
*/
func (a *Login) OnLogin(w rest.ResponseWriter, r *rest.Request) {
	var req map[string]string
	tools.GetRequestJsonObj(r, &req)

	username := req["username"]
	password := req["password"]
	times := req["time"]

	if len(username) == 0 || len(password) == 0 || len(times) == 0 {
		w.WriteJson(message.GetErrorMsg(message.ParameterInvalid))
		return
	}
	//-------------------------end

	msg := message.GetSuccessMsg(message.NormalMessageId)
	msg.Data = password
	w.WriteJson(msg)

	return
}

/*
	修改用户密码
*/
func ChangePwd(w rest.ResponseWriter, r *rest.Request) {
	var req Req_changepwd
	tools.GetRequestJsonObj(r, &req)

	//参数检测
	if !tools.CheckParam(req.UserName, req.Password) {
		w.WriteJson(message.GetErrorMsg(message.ParameterInvalid, req))
		return
	}

	//验证token，并获取用户名
	// _, bfind := user.GetUserDetail(req.UserName)
	// if !bfind {
	// 	w.WriteJson(message.GetErrorMsg(message.NotFindError))
	// 	return
	// }

	w.WriteJson(message.GetSuccessMsg())
	//w.WriteJson(message.GetErrorMsg(message.UserNameDoNotExist))

	return
}
