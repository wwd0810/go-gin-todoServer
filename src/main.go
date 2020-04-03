package main

import (
	"fmt"
	"log"
	"net/http"
	"todoServer/src/pkg/database"
	"todoServer/src/pkg/setting"
	"todoServer/src/routers"
)

func init() {
	setting.Setup()
	database.Setup()
}

func main() {

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] server listening %s", endPoint)

	server.ListenAndServe()
}
