package main

import (
	"fmt"
	"github.com/67in/go-gin-example/routers"
)

func main() {
	r := routers.InitRouter()
	_ = r.Run(fmt.Sprintf(":#{Port}"))
	//router.GET("")
	//
	//s := &http.Server{
	//	Addr:              fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:           router,
	//	TLSConfig:         nil,
	//	ReadTimeout:       setting.ReadTimeout,
	//	ReadHeaderTimeout: 0,
	//	WriteTimeout:      setting.WriteTimeout,
	//	IdleTimeout:       0,
	//	MaxHeaderBytes:    1 << 20,
	//	TLSNextProto:      nil,
	//	ConnState:         nil,
	//	ErrorLog:          nil,
	//	BaseContext:       nil,
	//	ConnContext:       nil,
	//}
	//s.ListenAndServe()
}
