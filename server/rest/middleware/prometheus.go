/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-29 23:56:37
 * @LastEditTime: 2024-03-22 18:15:11
 * @LastEditors: yujiajie
 */
package middleware

import (
	"gateway/core/metric"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const serverNamespace = "http_server"

var (
	metricServerReqDur = metric.NewHistogramVec(&metric.HistogramVecOpts{
		Namespace: serverNamespace,
		Subsystem: "requests",
		Name:      "duration_ms",
		Help:      "http server requests duration(ms).",
		Labels:    []string{"path", "method"},
		Buckets:   []float64{5, 10, 25, 50, 100, 250, 500, 1000},
	})

	metricServerReqCodeTotal = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: serverNamespace,
		Subsystem: "requests",
		Name:      "code_total",
		Help:      "http server requests error count.",
		Labels:    []string{"path", "code", "method"},
	})
)

func PrometheusHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		startTime := time.Now()
		defer func() {
			metricServerReqDur.Observe(time.Since(startTime).Milliseconds(), path, method)
			metricServerReqCodeTotal.Inc(path, strconv.Itoa(ctx.Writer.Status()), method)
		}()
		ctx.Next()
	}
}
