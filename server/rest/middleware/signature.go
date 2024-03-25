package middleware

import (
	"fmt"
	"gateway/core/codec"
	"gateway/server/rest/header"
	"gateway/service/mod"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SignHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		contentSignature := ctx.GetHeader(header.ContentSignature)
		attrs := header.ParseHeader(contentSignature)
		appid := attrs[header.AppIdField]
		timestamp := attrs[header.TimeField]
		nonce := attrs[header.NonceField]
		signature := attrs[header.SignatureField]

		reqTime, err := strconv.ParseInt(timestamp, 10, 64)
		if err != nil {
			fmt.Println("时间戳格式不对")
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		now := time.Now().Unix()
		if reqTime+10 < now {
			fmt.Println("请求已过期")
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		appModule := mod.GetAppModule(appid)
		app_secret := appModule.AppSecret
		if len(app_secret) == 0 {
			fmt.Println("appid不存在")
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		signContent := strings.Join([]string{
			"appid=" + appid,
			"nonce=" + nonce,
			"time=" + timestamp,
			app_secret,
		}, "&")
		actualSignature := codec.Md5(signContent)
		if actualSignature != signature {
			fmt.Println("签名有误")
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		path := ctx.Request.URL.Path
		modName := strings.Split(path, "/")[0]
		if _, exists := appModule.Modules[modName]; !exists {
			fmt.Println("没有该模块权限")
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}
