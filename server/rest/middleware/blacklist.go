/*
 * @Author: yujiajie
 * @Date: 2024-03-21 14:52:57
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-22 10:49:31
 * @FilePath: /gateway/server/rest/middleware/blacklist.go
 * @Description:
 */
package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

var (
	blackListIP = []string{}

	blackListUser = []int{}
)

func IpForbid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		if len(blackListIP) > 0 && slices.Contains[[]string, string](blackListIP, ip) {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		ctx.Next()
	}
}

func UserForbid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if authUser, exists := ctx.Get(Auth_User); exists {
			if user_id, ok := authUser.(int); ok {
				if len(blackListUser) > 0 && slices.Contains[[]int, int](blackListUser, user_id) {
					ctx.AbortWithStatus(http.StatusForbidden)
					return
				}
			}
		}
		ctx.Next()
	}
}
