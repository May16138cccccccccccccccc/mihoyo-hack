package main

// 这个demo 设计用户发帖功能
import (
	"github.com/gin-gonic/gin"
	"go-article-demo/controller"
	"net/http"
)

func main() {
	//默认路由引擎实例化
	r := gin.Default()
	// 案例1：发送数据
	//r.GET("/ping", func(c *gin.Context) {
	//	//输出json结果给调用方
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	// 案例2：访问静态资源：图片
	r.StaticFS("/static", http.Dir("E:/qibu/gosoft/go/src/go-article-demo/web/static"))
	r.StaticFile("/test1.png", ("./web/static/test1.png"))

	// 导入前端模板
	r.LoadHTMLGlob("web/template/*")
	// 前端web分组，有一个路由接口
	v := r.Group("/")
	{
		//从包文件调用（类）go文件的方法，直接调用，因为go没有类的概念，所以直接调用跳过文件名controllerTest.go
		v.GET("/index.html", controller.Index)
	}

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	//# 运行 main.go 并且在浏览器中访问 HOST_IP:8080/ping
}
