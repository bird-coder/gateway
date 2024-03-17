/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-26 17:42:29
 * @LastEditTime: 2023-11-26 17:54:49
 * @LastEditors: yujiajie
 */
package middleware

import (
	"net/http"

	"github.com/alibaba/sentinel-golang/core/system"
	sentinel "github.com/alibaba/sentinel-golang/pkg/adapters/gin"
	"github.com/gin-gonic/gin"
)

func Sentinel() gin.HandlerFunc {
	if _, err := system.LoadRules([]*system.Rule{
		{
			MetricType:   system.InboundQPS,
			TriggerCount: 200,
			Strategy:     system.BBR,
		},
	}); err != nil {

	}
	return sentinel.SentinelMiddleware(
		sentinel.WithBlockFallback(func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg":  "too many request; the quota used up!",
				"code": 500,
			})
		}),
	)
}
