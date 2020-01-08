package main

import (
	"fmt"
	 "github.com/67in/go-gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", setting2.HTTPPort),
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       setting.ReadTimeout,
		ReadHeaderTimeout: 0,
		WriteTimeout:      setting.WriteTimeout,
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	s.ListenAndServe()
}
