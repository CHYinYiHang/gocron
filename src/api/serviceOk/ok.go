package serviceOk

import (
	"fmt"
	"github.com/CHYinYiHang/gocron/pkg/app"
	"github.com/CHYinYiHang/gocron/pkg/config"
	"github.com/CHYinYiHang/gocron/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"sync/atomic"
	"time"
)

func Ok(ctx *gin.Context) {
	var appG = app.Gin{C: ctx}
	a := make(map[string]interface{})

	a["NowTime"] = time.Now().Format("2006年01月02日 15:04")
	a["ServerId"] = config.Config.Server.ServerId
	a["NumGoroutine"] = runtime.NumGoroutine()

	var opts int64 = 0

	for i := 0; i < 1000; i++ {

		go func() {
			// 注意第一个参数必须是地址
			atomic.AddInt64(&opts, 1) //加操作
			//atomic.AddInt64(&opts, -1) 减操作
		}()

	}

	time.Sleep(time.Second * 1)

	fmt.Println("opts: ", atomic.LoadInt64(&opts))

	appG.Response(http.StatusOK, e.SUCCESS, "ok", a)
	return
}
