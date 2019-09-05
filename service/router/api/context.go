package api

/*
* 基础类目 下一个版本将支持 单个struct 自动解析。
 */

import (
	"ginrest/service/view/user"

	"github.com/gin-gonic/gin"
	"github.com/xxjwxc/public/mylog"
)

const (
	//UserToken 用户token
	UserToken = "user_token"
)

//Context 上下文
type Context struct {
	*gin.Context
}

//Newctx 新建
func Newctx(c *gin.Context) *Context {
	return &Context{c}
}

//GetVersion 获取版本号
func (c *Context) GetVersion() string {
	return c.Param("version")
}

//WriteJSON 写入json对象
func (c *Context) WriteJSON(obj interface{}) {
	c.JSON(200, obj)
}

//GetUserInfo 获取用户信息
func (c *Context) GetUserInfo() (u *user.UserInfo, b bool) {
	accessToken, err := c.Cookie(UserToken)
	if err == nil {
		u = &user.UserInfo{}
		u.Username, u.ExpireTime, b = user.GetUserFromToken(accessToken)
		return
	}

	mylog.ErrorString(err.Error())
	return nil, false
}

//GetDid 获取设备did
func (c *Context) GetDid() string {
	//先获取用户
	accessToken, err := c.Cookie(UserToken)
	if err == nil {
		username, _, b := user.GetUserFromToken(accessToken)
		if b {
			return username
		}
	}

	return c.GetHeader("did")
}

//GetLanguage 获取语言
func (c *Context) GetLanguage() string {
	return c.GetHeader("lg")
}

// //获取用户ip
// func (c *Context) GetClientIp() string {
// 	return c.ClientIP()
// }
