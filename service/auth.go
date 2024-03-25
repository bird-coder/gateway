/*
 * @Author: yujiajie
 * @Date: 2024-03-22 10:22:17
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-25 09:20:34
 * @FilePath: /gateway/service/auth.go
 * @Description:
 */
package service

import (
	"gateway/core/auth"
	"gateway/core/container"
	"gateway/options"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	data := map[string]interface{}{
		"user_id": 1,
	}
	authConfig := container.App.GetConfig("auth").(options.AuthConfig)
	tk, err := auth.Issue(data, []byte(authConfig.Secret))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"msg":  "login failed",
			"code": 500,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":  "success",
		"code": 200,
		"data": map[string]string{
			"access_token": tk,
		},
	})
}
