/*
	路由表
*/
package router

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/xie1xiao1jun/ginrest/service/view/test"
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

/*
	检测重复
*/
func OnCheckRouter(routes ...*rest.Route) {
	mp := make(map[string]int)
	for _, v := range routes {
		if mp[v.PathExp] == 1 {
			panic("============" + v.PathExp + " : is repeat in rooter.go.================")
		} else {
			mp[v.PathExp] = 1
		}
	}
}
