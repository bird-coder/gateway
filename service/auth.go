/*
 * @Author: yujiajie
 * @Date: 2024-03-22 10:22:17
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-22 10:40:51
 * @FilePath: /gateway/service/auth.go
 * @Description:
 */
package service

import (
	"gateway/core/auth"
	"gateway/options"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	data := map[string]interface{}{
		"user_id": 1,
	}
	tk, err := auth.Issue(data, []byte(options.App.Auth.Secret))
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
