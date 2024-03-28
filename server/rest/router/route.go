/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-26 17:36:46
 * @LastEditTime: 2024-03-28 15:59:25
 * @LastEditors: yujiajie
 */
package router

import (
	"gateway/core/constant"
	"gateway/core/container"
	"gateway/service"
	"net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init(r *gin.Engine) {
	//prometheus采集指标的接口
	r.GET("/metrics", func(ctx *gin.Context) {
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	})
	r.GET("/auth/login", service.Login)

	if container.App.GetEnv() == string(constant.Dev) {
		handlePprof(r)
	}
}

func handlePprof(r *gin.Engine) {
	pprofGroup := r.Group("/debug/pprof")
	{
		pprofGroup.GET("/", gin.WrapF(pprof.Index))
		pprofGroup.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		pprofGroup.GET("/profile", gin.WrapF(pprof.Profile))
		pprofGroup.POST("/symbol", gin.WrapF(pprof.Symbol))
		pprofGroup.GET("/symbol", gin.WrapF(pprof.Symbol))
		pprofGroup.GET("/trace", gin.WrapF(pprof.Trace))
		pprofGroup.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
		pprofGroup.GET("/block", gin.WrapH(pprof.Handler("block")))
		pprofGroup.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
		pprofGroup.GET("/heap", gin.WrapH(pprof.Handler("heap")))
		pprofGroup.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))
		pprofGroup.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
	}
}
