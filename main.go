package main

import (
	"fmt"
	"github.com/CHYinYiHang/gocron/pkg/config"
	"github.com/CHYinYiHang/gocron/pkg/drivers"
	"github.com/CHYinYiHang/gocron/pkg/firlter"
	"github.com/CHYinYiHang/gocron/pkg/logging"
	"github.com/CHYinYiHang/gocron/pkg/utils"
	"github.com/CHYinYiHang/gocron/src/routers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func init() {
	config.LoadConf()
	drivers.LoadMysql()
	logging.LoadLogin()
	utils.LoadSnowflakeNode(config.Config.Server.ServerId)
}

/*
	此为gin的入口函数，编译此文件即可，main上层也是通过GOROUTINE协程处理的。
*/
func main() {

	/*
		程序停止是关闭数据连接
	*/
	defer func() {
		_ = drivers.MysqlDbPool.Close()
	}()

	//获取gin路由对象 使用默认的gin.Logger(),gin.Recovery() 插件
	r := gin.Default()

	//注册跨域组件
	r.Use(firlter.Cors())

	//初始化路由体
	routers.InitRouter(r)

	endPoint := fmt.Sprintf("0.0.0.0%s", config.Config.Server.Port)

	//配置服务对象
	s := &http.Server{
		Addr:           endPoint,         //监听的TCP地址
		Handler:        r,                //http句柄,用于处理程序响应HTTP请求
		ReadTimeout:    60 * time.Second, //允许读取的最大时间
		WriteTimeout:   60 * time.Second, //允许写入的最大时间
		MaxHeaderBytes: 1 << 20,          //请求头的最大字节数
	}

	logging.Info("============服务启动成功============")
	logging.Info("============服务启动加载配置成功============")
	logging.Info("============服务启动加载数据库连接成============")
	logging.Info("============服务启动加载日志成功============")
	logging.Info("============" + time.Now().Format("2006-01-02 15:04:05") + "============")
	//启动服务监听
	err := s.ListenAndServe()
	if err != nil {
		logging.Error("============服务启动监听失败============", err.Error())
	}
}
