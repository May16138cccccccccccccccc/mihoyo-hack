package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	// 控制层根目录下创建了controller文件夹并创建go文件，主要用来存放逻辑层的操作
	// 下面这段代码是gin框架上下文实例化到调用html代码之后，使用tab命令直接生成
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "用户发帖功能测试",
	})
}
