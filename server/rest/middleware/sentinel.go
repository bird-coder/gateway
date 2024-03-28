/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-26 17:42:29
 * @LastEditTime: 2024-03-28 16:00:34
 * @LastEditors: yujiajie
 */
package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/core/system"
	sentinel "github.com/alibaba/sentinel-golang/pkg/adapters/gin"
	"github.com/gin-gonic/gin"
)

// 根据cpu使用率限制请求，自适应保护系统
func Sentinel() gin.HandlerFunc {
	resName := "gateway"
	cores := runtime.NumCPU()
	if _, err := system.LoadRules([]*system.Rule{
		{
			MetricType:   system.Load,
			TriggerCount: float64(cores) * 2.5,
			Strategy:     system.BBR,
		},
	}); err != nil {
		fmt.Println(err)
	}
	return sentinel.SentinelMiddleware(
		sentinel.WithBlockFallback(func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg":  "too high system load; the quota used up!",
				"code": 500,
			})
		}),
		sentinel.WithResourceExtractor(func(ctx *gin.Context) string {
			return resName
		}),
	)
}

// 流量控制
func FlowHandler(resName string, threshold int) gin.HandlerFunc {
	if _, err := flow.LoadRulesOfResource(resName, []*flow.Rule{
		{
			Resource:               resName,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			Threshold:              float64(threshold),
			StatIntervalInMs:       1000,
		},
	}); err != nil {
		fmt.Println(err)
	}

	return sentinel.SentinelMiddleware(
		sentinel.WithBlockFallback(func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg":  "too many request; the quota used up!",
				"code": 500,
			})
		}),
		sentinel.WithResourceExtractor(func(ctx *gin.Context) string {
			return resName
		}),
	)
}

// 根据错误率熔断
func ErrorBreakerHandler(resName string) gin.HandlerFunc {
	if _, err := circuitbreaker.LoadRulesOfResource(resName, []*circuitbreaker.Rule{
		{
			Resource:         resName,
			Strategy:         circuitbreaker.ErrorRatio,
			RetryTimeoutMs:   3000,
			MinRequestAmount: 10,
			StatIntervalMs:   10000,
			Threshold:        0.3,
		},
	}); err != nil {
		fmt.Println(err)
	}

	return func(ctx *gin.Context) {
		entry, err := api.Entry(
			resName,
			api.WithResourceType(base.ResTypeAPIGateway),
			api.WithTrafficType(base.Inbound),
		)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg":  "too many error; break request!",
				"code": 500,
			})
			return
		}
		defer entry.Exit()
		ctx.Next()

		code := ctx.Writer.Status()
		if code >= http.StatusBadRequest {
			api.TraceError(entry, errors.New("service unavailable"))
		}
	}
}

// 根据慢请求率熔断
func SlowBreakerHandler(resName string, threshold int) gin.HandlerFunc {
	if _, err := circuitbreaker.LoadRules([]*circuitbreaker.Rule{
		{
			Resource:         resName,
			Strategy:         circuitbreaker.SlowRequestRatio,
			RetryTimeoutMs:   3000,
			MinRequestAmount: 10,
			StatIntervalMs:   10000,
			MaxAllowedRtMs:   uint64(threshold),
			Threshold:        0.3,
		},
	}); err != nil {
		fmt.Println(err)
	}

	return sentinel.SentinelMiddleware(
		sentinel.WithBlockFallback(func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg":  "too many request; the quota used up!",
				"code": 500,
			})
		}),
		sentinel.WithResourceExtractor(func(ctx *gin.Context) string {
			return resName
		}),
	)
}
