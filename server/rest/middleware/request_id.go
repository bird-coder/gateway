/*
 * @Author: yujiajie
 * @Date: 2024-03-22 11:31:45
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-28 16:06:40
 * @FilePath: /Gateway/server/rest/middleware/request_id.go
 * @Description:
 */
package middleware

import (
	"gateway/core/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequestId() gin.HandlerFunc {
	node, _ := uuid.NewNode(1)
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {
			ctx.Next()
			return
		}
		requestId := node.Generate()
		ctx.Set("RequestId", requestId.Int64())
		ctx.Next()
	}
}
