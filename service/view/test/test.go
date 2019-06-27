package test

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xxjwxc/ginrest/service/router/api"
)

//
var T TTT

type TTT struct {
}

//TestFun1 gin 默认的函数回调地址
func TestFun1(c *gin.Context) {
	fmt.Println(c.Params)
}

func (t *TTT) TestFun11(c *gin.Context) {
	fmt.Println(c.Params)
}

//TestFun2 自定义context的函数回调地址
func (t *TTT) TestFun2(c *api.Context) {
	fmt.Println(c.Params)
}

//TestFun3 带自定义context跟已解析的req参数回调方式
func TestFun3(c *api.Context, req *ReqTest) {
	fmt.Println(c.Params)
	fmt.Println(req)
}

//TestFun3 带自定义context跟已解析的req参数回调方式
func TestFun4(c *gin.Context, req ReqTest) {
	fmt.Println(c.Params)
	fmt.Println(req)
}
