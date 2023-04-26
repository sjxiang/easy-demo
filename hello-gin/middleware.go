package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


type Middlewares struct {

}


func (mds Middlewares) Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		duration := time.Since(start)
		fmt.Printf("耗时：%v\n", duration)
	}
}



func (Middlewares) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		
		if token != "123456" {
			
			// 终止后续逻辑的执行
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "登录失败",
			})
		}

		ctx.Next()
	}

}