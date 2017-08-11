package main

import (
	"common/appconfig"
	"common/appconstant"
	"common/dao/mysqlconn"
	"common/notification"
	"fmt"
	"github.com/jabong/floRest/src/service"
	"hello"
	"newApp/getSelfhelpDetail"
	model "newApp/model/sqldb"
	"newApp/second"
	"newApp/third"
)

//main is the entry point of the florest web service
func main() {
	fmt.Println("APPLICATION BEGIN")
	webserver := new(service.Webserver)
	registerConfig()
	registerErrors()
	registerAllApis()
	mysqlconn.Initialize()
	notification.InitNotifpool()
	model.SetCounter("counter", 0)
	model.SetMaxValue()
	webserver.Start()

}

func registerAllApis() {
	service.RegisterApi(new(hello.HelloApi))
	service.RegisterApi(new(getSelfhelpDetail.CategoryDetails))
	service.RegisterApi(new(second.Second))
	service.RegisterApi(new(third.Third))
}
func registerConfig() {
	service.RegisterConfig(new(appconfig.AppConfig))
}

func registerErrors() {
	service.RegisterHttpErrors(appconstant.AppErrorCodeToHttpCodeMap)
}
