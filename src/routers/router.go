package routers

import (
	"github.com/CHYinYiHang/gocron/src/api/serviceOk"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	/*
		路由初始化方法
		在这里配置APi，url对应的处理函数

		此为单条路由示例
		r.POST("/url/", fun)
		r.GET("/url/", fun)
		r.StaticFile("/", config.Config.WebServer.IndexHtmlPath) //前端入口
		r.Static("/static", config.Config.WebServer.StaticPath)  // 静态资源路径
		//图片访问路径
		r.StaticFS("/fileData", http.Dir(config.Config.WebServer.FileDataPath))
	*/
	server := r.Group("/server")
	{
		server.GET("/cat/status", serviceOk.Ok)
	}


	return r
}
