package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xxjwxc/public/dev"
)

var apiRoot *gin.Engine
var route *gin.RouterGroup

//版本号
var version = "/:version"

func init() {
	if dev.OnIsDev() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
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
