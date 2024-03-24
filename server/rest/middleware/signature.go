package middleware

import "github.com/gin-gonic/gin"

func SignHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
