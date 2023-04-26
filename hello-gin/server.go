package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var signals = []os.Signal{
	syscall.SIGINT,  // kill -2，等同于 CTRL + c
	syscall.SIGTERM, // kill pid
	syscall.SIGQUIT, // ctrl + \
}



func main() {
	router := gin.New()

	router.Use(Middlewares{}.Logger())

	// 会将指定目录下的文件加载好，相对目录
	router.LoadHTMLGlob("templates/**/*")  // 二级子目录

	router.GET("/index", Index)

	v2 := router.Group("/v2")
	{
		v2.GET("/videos", ListVideo) 
		v2.GET("/video/:id", ShowVideo)
		v2.POST("/videos", CreateVideo)
	}

	httpSrv := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, signals...)
	<-quit
	log.Println("即将关闭服务器 ...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Printf("强制关机，%s", err)
	}

	log.Println("服务器关闭")
}