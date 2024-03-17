/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-27 23:05:09
 * @LastEditTime: 2023-11-27 23:07:52
 * @LastEditors: yujiajie
 */
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoverHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
