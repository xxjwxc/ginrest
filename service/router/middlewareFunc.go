/*
	为各种回调机制的默认事件
*/
package router

import (
	"github.com/ant0ine/go-json-rest/rest"
)

//
type MiddlewareFunc struct {
	//在此添加变量
}

//从服务器取出资源（一项或多项）。
func Get(pathExp string, handlerFunc interface{}) {
	route.GET(pathExp, getHandlerFunc(handlerFunc))
}

//在服务器新建一个资源。
func Post(pathExp string, handlerFunc interface{}) {
	route.POST(pathExp, getHandlerFunc(handlerFunc))
}

//在服务器更新资源（客户端提供改变后的完整资源）。
func Put(pathExp string, handlerFunc interface{}) {
	route.PUT(pathExp, getHandlerFunc(handlerFunc))
}

//在服务器更新资源（客户端提供改变的属性）。
func Patch(pathExp string, handlerFunc interface{}) {
	route.PATCH(pathExp, getHandlerFunc(handlerFunc))
}

//从服务器删除资源。
func Delete(pathExp string, handlerFunc interface{}) {
	route.DELETE(pathExp, getHandlerFunc(handlerFunc))
}

//-----------------------------end

/*
 Token校验，并返回client_id
*/
func TokenMwFunc(handler rest.HandlerFunc) (b bool, client_id string) {

	return
}

/*
 版本号检测,并返回版本号
*/
func VersionMwFunc(handler rest.HandlerFunc) (b bool, version string) {
	b = false
	version = ""
	return
}
