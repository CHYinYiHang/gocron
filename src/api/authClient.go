package api

import (
	"github.com/CHYinYiHang/gocron/pkg/app"
	"github.com/CHYinYiHang/gocron/pkg/e"
	"github.com/CHYinYiHang/gocron/src/service/AuthService"
	"github.com/gin-gonic/gin"
	"net/http"
)

//创建客户端并生成token
func CreateClient(ctx *gin.Context) {
	var (
		appG = app.Gin{C: ctx}
	)

	name := ctx.PostForm("client_name")
	client, _ := AuthService.GenerateClient(name)
	token, err := client.GenerateToken()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err.Error(), "")
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, "ok", token)
}
