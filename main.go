package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/xxjwxc/ginrest/service"
	"github.com/xxjwxc/ginrest/service/config"
	"github.com/xxjwxc/public/dev"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/server"
)

type handlerFunc1 func(*gin.Context)

//
func CallBack() {
	service.ApiRoot{}.OnCreat()

	//https 支持(单开一个线程)
	//	go func() {
	//		log.Println("https is running at " + config.GetServerHttpsPort() + " port.")
	//		log.Fatal(http.ListenAndServeTLS(":"+config.GetServerHttpsPort(), tools.GetModelPath()+"/pem/cacert.pem", tools.GetModelPath()+"/pem/privatekey.pem", nil))
	//	}()
	//---------------------------end

	//启动http
	mylog.Println("http is running at " + config.GetServerPort() + " ===>" + "service:" + dev.GetService())
	mylog.Fatal(http.ListenAndServe(":"+config.GetServerPort(), nil))
}

func main() {
	if config.OnIsDev() || len(os.Args) == 0 {
		CallBack()
	} else {
		server.On(config.GetServiceConfig()).Start(CallBack)
	}
}
