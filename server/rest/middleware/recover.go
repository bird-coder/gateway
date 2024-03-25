/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-27 23:05:09
 * @LastEditTime: 2024-03-25 17:48:06
 * @LastEditors: yujiajie
 */
package middleware

import (
	"gateway/core/container"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoverHandler() gin.HandlerFunc {
	log := container.App.GetLogger("default")
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Error("系统错误: %v", err)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
