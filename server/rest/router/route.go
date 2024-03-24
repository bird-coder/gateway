/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-26 17:36:46
 * @LastEditTime: 2024-03-22 10:39:31
 * @LastEditors: yujiajie
 */
package router

import (
	"gateway/service"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init(r *gin.Engine) {
	r.GET("/metrics", func(ctx *gin.Context) {
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	})
	r.GET("/auth/login", service.Login)
}
