/*
	路由表
*/
package router

import (
	"github.com/xxjwxc/ginrest/service/view/test"
)

//
func DefaultRouler() {
	Get("ttt", test.TestFun1)
	Get("ttt1", test.T.TestFun2)
	Post("ttt2", test.TestFun3)
	Post("ttt4", test.TestFun4)
}

/*
	默认初始化
*/
func OnInitRouter() (err error) {
	DefaultRouler()
	return nil
}
