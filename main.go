package main

import (
	"log"
	"net/http"
	"os"

	"github.com/xie1xiao1jun/go-restful/data"
	"github.com/xie1xiao1jun/go-restful/data/config"
	"github.com/xie1xiao1jun/public/server"
)

//
func CallBack() {
	var apiroot data.ApiRoot
	apiroot.OnCreat()

	//https 支持(单开一个线程)
	//	go func() {
	//		log.Println("https is running at " + config.GetServerHttpsPort() + " port.")
	//		log.Fatal(http.ListenAndServeTLS(":"+config.GetServerHttpsPort(), tools.GetModelPath()+"/pem/cacert.pem", tools.GetModelPath()+"/pem/privatekey.pem", nil))
	//	}()
	//---------------------------end
	//启动http
	log.Println("http is running at " + config.GetServerPort() + " port.")
	log.Fatal(http.ListenAndServe(":"+config.GetServerPort(), nil))
}

func main() {
	if config.OnIsDev() || len(os.Args) == 0 {
		CallBack()
	} else {
		server.OnStart(CallBack)
	}
}
