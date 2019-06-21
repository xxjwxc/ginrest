/*
	构造restful基础框架类
*/
package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xie1xiao1jun/ginrest/service/config"
	"github.com/xie1xiao1jun/ginrest/service/router"
	"github.com/xie1xiao1jun/ginrest/service/view/file"
)

// var api *rest.Api = nil

type ApiRoot struct {
}

func (*ApiRoot) OnCreat() {

	if config.OnIsDev() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router.OnInitRouter()

	// api = rest.NewApi()
	// if config.OnIsDev() {
	// 	api.Use(rest.DefaultDevStack...) //DefaultProdStack
	// } else {
	// 	api.Use(rest.DefaultProdStack...)
	// }
	// router, err := router.OnInitRoter(api)
	// if err != nil {
	// 	log.Fatal(err)
	// 	panic(err)
	// }
	// api.SetApp(router)

	//创建模版解析 ,创建文件

	//http.HandleFunc("/apiserver/upload/", upload.M_upload.UpLoadOne) //添加post form
	//微信相关接口
	//	http.HandleFunc("/hotelserver/index.do", weixin.Index)
	http.HandleFunc("/"+config.GetServiceName()+"/api/v1/file/upload", file.O_file.Upload) //统一上传文件
	//http.HandleFunc(config.Url_host+"/api/v1/file/upload", file.O_file.UploadToCos) //统一上传文件cos
	//-------------------end

	http.Handle("/"+config.GetServiceName()+"/api/", http.StripPrefix("/"+config.GetServiceName()+"/api", router.GetRoot())) //指定api默认路由
}
