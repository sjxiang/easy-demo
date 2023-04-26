package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)



func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home/index.tmpl", gin.H{
		"content": "欢迎",
	})
}


func ListVideo(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")

	ctx.HTML(http.StatusOK, "video/list.tmpl", gin.H{
		"content": fmt.Sprintf("视频列表页(%s条/页) - 第 %s 页", offset, limit),
	})
}

func ShowVideo(ctx *gin.Context) {
	// 获取动态路由参数
	id := ctx.Param("id")
	ctx.String(http.StatusOK, "id = %s", id)
}

func CreateVideo(ctx *gin.Context) {
	var request CreateVideoRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}

type CreateVideoRequest struct {
	Title  string `json:"title"  binding:"required,min=3,max=100"`
	Info   string `json:"info"   binding:"required"`
	URL    string `json:"url"    binding:"required"`
	Avatar string `json:"avatar" binding:"required"`
}
