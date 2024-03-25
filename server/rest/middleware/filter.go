/*
 * @Author: yujiajie
 * @Date: 2024-03-25 16:39:08
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-25 17:32:08
 * @FilePath: /gateway/server/rest/middleware/filter.go
 * @Description:
 */
package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	m sync.Map
)

func FilterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		key := path + ":" + method
		now := time.Now().Unix()
		if val, ok := m.Load(key); ok {
			timestamp := *val.(*int64)
			fmt.Println("路由不存在，拦截已过期", path, key)
			if timestamp+30 > now {
				fmt.Println("路由不存在，被拦截", path, key)
				ctx.AbortWithStatus(http.StatusNotFound)
				return
			}
		}
		ctx.Next()

		code := ctx.Writer.Status()
		if code == http.StatusNotFound {
			m.Store(key, &now)
		}
	}
}
