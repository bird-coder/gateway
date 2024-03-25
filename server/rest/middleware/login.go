package middleware

import (
	"bytes"
	"encoding/json"
	"gateway/core/auth"
	"gateway/core/container"
	"gateway/options"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		content, _ := io.ReadAll(ctx.Request.Response.Body)

		ctx.Request.Response.Body = io.NopCloser(bytes.NewBuffer(content))

		data := initData(content)
		if data == nil {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg":  "login failed",
				"code": 500,
			})
			return
		}

		authConfig := container.App.GetConfig("auth").(options.AuthConfig)
		tk, err := auth.Issue(data, []byte(authConfig.Secret))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg":  "login failed",
				"code": 500,
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"msg":  "success",
			"code": 200,
			"data": map[string]string{
				"access_token": tk,
			},
		})
	}
}

func initData(body []byte) map[string]interface{} {
	var res map[string]interface{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil
	}
	if res["code"] != 200 {
		return nil
	}
	data := res["data"].(map[string]interface{})
	authData := map[string]interface{}{
		"user_id": data["uid"],
	}
	return authData
}
