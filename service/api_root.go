/*
	构造restful基础框架类
*/
package data

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/xie1xiao1jun/ginrestful/service/api/file"
	"github.com/xie1xiao1jun/ginrestful/service/config"
	"github.com/xie1xiao1jun/ginrestful/service/router"
)

var api *rest.Api = nil

type ApiRoot struct {
}

func (*ApiRoot) OnCreat() {
	api = rest.NewApi()
	if config.OnIsDev() {
		api.Use(rest.DefaultDevStack...) //DefaultProdStack
	} else {
		api.Use(rest.DefaultProdStack...)
	}
	router, err := router.OnInitRoter(api)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	api.SetApp(router)

	//创建模版解析 ,创建文件

	//http.HandleFunc("/apiserver/upload/", upload.M_upload.UpLoadOne) //添加post form
	//微信相关接口
	//	http.HandleFunc("/hotelserver/index.do", weixin.Index)
	http.HandleFunc(config.Url_host+"/api/v1/file/upload", file.O_file.Upload) //统一上传文件
	//http.HandleFunc(config.Url_host+"/api/v1/file/upload", file.O_file.UploadToCos) //统一上传文件cos

	//-------------------end

	http.Handle(config.Url_host+"/api/", http.StripPrefix(config.Url_host+"/api", api.MakeHandler())) //指定api默认路由
}
