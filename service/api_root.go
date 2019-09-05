/*
	构造restful基础框架类
*/
package service

import (
	"net/http"

	"ginrest/service/config"
	"ginrest/service/router"
	"ginrest/service/view/file"

	"github.com/xxjwxc/public/dev"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
)

// var api *rest.Api = nil

type ApiRoot struct {
}

func (ApiRoot) OnCreat() {

	router.OnInitRouter()

	//创建模版解析 ,创建文件
	//http.HandleFunc("/apiserver/upload/", upload.M_upload.UpLoadOne) //添加post form
	//微信相关接口
	//	http.HandleFunc("/hotelserver/index.do", weixin.Index)
	http.HandleFunc("/"+config.GetServiceName()+"/api/v1/file/upload", file.O_file.Upload) //统一上传文件

	mylog.Debug("file upload POST -->" + "/" + config.GetServiceName() + "/api/v1/file/upload")
	//http.HandleFunc(config.Url_host+"/api/v1/file/upload", file.O_file.UploadToCos) //统一上传文件cos
	//-------------------end

	buildStatic()

	apGroupiName := "/" + config.GetServiceName() + "/api"
	http.Handle(apGroupiName+"/", http.StripPrefix(apGroupiName, router.GetRoot())) //指定api默认路由
	mylog.Debug("group host --> " + apGroupiName)
}

func buildStatic() {
	if len(dev.GetFileHost()) > 0 {
		// 设置静态目录
		pattern := "/" + dev.GetService() + "/" + dev.GetFileHost()
		fsh := http.FileServer(http.Dir(tools.GetCurrentDirectory() + "/" + dev.GetFileHost()))
		http.Handle(pattern+"/", http.StripPrefix(pattern, fsh))
		mylog.Debug("static file host -->" + pattern)
	}
}
