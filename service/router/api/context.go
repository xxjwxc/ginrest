package api

/*
* 基础类目 下一个版本将支持 单个struct 自动解析。
 */

import (
	"github.com/gin-gonic/gin"
	"github.com/xxjwxc/ginrest/service/view/user"
	"github.com/xxjwxc/public/mylog"
)

type Context struct {
	*gin.Context
}

//
func Newctx(c *gin.Context) *Context {
	return &Context{c}
}

//获取版本号
func (c *Context) GetVersion() string {
	return c.Param("version")
}

//写入json对象
func (c *Context) WriteJson(obj interface{}) {
	c.JSON(200, obj)
}

//获取用户信息
func (c *Context) GetUserInfo() (u *user.UserInfo, b bool) {
	access_token, err := c.Cookie("user_token")
	if err != nil {
		u = &user.UserInfo{}
		u.Username, u.ExpireTime, b = user.GetUserFromToken(access_token)
		return
	}

	mylog.Error(err)
	return nil, false
}
