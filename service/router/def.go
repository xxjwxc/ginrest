package router

import (
	"github.com/gin-gonic/gin"
)

var apiRoot *gin.Engine
var route *gin.RouterGroup

//版本号
var version = "/:version"

func init() {
	//构造默认api
	apiRoot = gin.Default()
	//router = apiRoot.Group(config.Url_host + "/api")
	route = apiRoot.Group(version)
}

//基本路由
func GetRoot() *gin.Engine {
	return apiRoot
}

//路由群组
func GetRoute() *gin.RouterGroup {
	return route
}
