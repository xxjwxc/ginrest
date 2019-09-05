/*
	路由表
*/
package router

import "ginrest/service/view/test"

/*
	默认初始化
*/
func OnInitRouter() {
	Get("ttt", test.TestFun1)
	Get("ttt1", test.T.TestFun2)
	Post("ttt2", test.TestFun3)
	Post("ttt4", test.TestFun4)
}
