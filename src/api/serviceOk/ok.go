package serviceOk

import (
	"github.com/CHYinYiHang/gocron/pkg/app"
	"github.com/CHYinYiHang/gocron/pkg/config"
	"github.com/CHYinYiHang/gocron/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Ok(ctx *gin.Context) {
	var appG = app.Gin{C: ctx}
	a := make(map[string]interface{})

	a["NowTime"] = time.Now().Format("2006年01月02日 15:04")
	a["ServerId"] = config.Config.Server.ServerId

	appG.Response(http.StatusOK, e.SUCCESS, "ok", a)
	return
}
