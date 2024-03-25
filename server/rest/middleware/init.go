/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-26 17:55:56
 * @LastEditTime: 2024-03-25 17:12:02
 * @LastEditors: yujiajie
 */
package middleware

import (
	"gateway/core/container"
	"gateway/core/stat"
	"gateway/core/trace"
	"gateway/options"

	"github.com/alibaba/sentinel-golang/api"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Init(r *gin.Engine) {
	gatewayConfig := container.App.GetConfig("gateway").(*options.GatewayConf)
	middlewareConfig := gatewayConfig.RestConf.Middlewares
	r.Use(NoCache)
	r.Use(Options)
	r.Use(Secure)
	if middlewareConfig.Gunzip {
		r.Use(gzip.Gzip(gzip.DefaultCompression))
	}
	if middlewareConfig.Recover {
		r.Use(RecoverHandler())
	}
	if middlewareConfig.Trace {
		trace.NewOsExporter(trace.DefaultService)
		r.Use(otelgin.Middleware(trace.DefaultService))
	}
	if middlewareConfig.Log {
		r.Use(LogHandler())
	}
	if middlewareConfig.Prometheus {
		r.Use(PrometheusHandler())
	}
	api.InitWithConfigFile("config/sentinel.yaml")
	r.Use(Sentinel())
	if middlewareConfig.Metrics {
		r.Use(MetricHandle(stat.NewMetrics("test")))
	}
	if middlewareConfig.BlackList {
		r.Use(IpForbid())
		r.Use(UserForbid())
	}
	r.Use(RequestId())
	if middlewareConfig.Filter {
		r.Use(FilterHandler())
	}
}
