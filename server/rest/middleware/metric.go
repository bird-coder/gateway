/*
 * @Author: yujiajie
 * @Date: 2024-03-18 16:18:09
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-18 17:48:31
 * @FilePath: /gateway/server/rest/middleware/metric.go
 * @Description:
 */
package middleware

import (
	"gateway/core/stat"
	"time"

	"github.com/gin-gonic/gin"
)

func MetricHandle(metrics *stat.Metrics) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		defer func() {
			metrics.Add(stat.Task{
				Duration: time.Since(start),
			})
		}()
		ctx.Next()
	}
}
