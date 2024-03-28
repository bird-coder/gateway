/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-26 17:55:56
 * @LastEditTime: 2024-03-28 16:01:14
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
	// r.Use(NoCache)
	r.Use(Options)
	r.Use(Secure)
	api.InitWithConfigFile("config/sentinel.yaml")
}

func Common() []gin.HandlerFunc {
	var commonMiddle []gin.HandlerFunc
	gatewayConfig := container.App.GetConfig("gateway").(*options.GatewayConf)
	middlewareConfig := gatewayConfig.RestConf.Middlewares

	commonMiddle = append(commonMiddle, RequestId())
	if middlewareConfig.Gunzip {
		commonMiddle = append(commonMiddle, gzip.Gzip(gzip.DefaultCompression))
	}
	if middlewareConfig.Recover {
		commonMiddle = append(commonMiddle, RecoverHandler())
	}
	if middlewareConfig.Trace {
		trace.NewOsExporter(trace.DefaultService)
		commonMiddle = append(commonMiddle, otelgin.Middleware(trace.DefaultService))
	}
	if middlewareConfig.Log {
		commonMiddle = append(commonMiddle, LogHandler())
	}
	if middlewareConfig.Prometheus {
		commonMiddle = append(commonMiddle, PrometheusHandler())
	}
	commonMiddle = append(commonMiddle, Sentinel())
	if middlewareConfig.Metrics {
		commonMiddle = append(commonMiddle, MetricHandle(stat.NewMetrics("test")))
	}
	if middlewareConfig.BlackList {
		commonMiddle = append(commonMiddle, IpForbid())
		commonMiddle = append(commonMiddle, UserForbid())
	}
	if middlewareConfig.Filter {
		commonMiddle = append(commonMiddle, FilterHandler())
	}
	if middlewareConfig.Auth {
		authConfig := container.App.GetConfig("auth").(*options.AuthConfig)
		authHandle := Authorize(authConfig.Secret, WithPrevSecret(authConfig.PrevSecret))
		commonMiddle = append(commonMiddle, authHandle)
	}
	if middlewareConfig.Sign {
		commonMiddle = append(commonMiddle, SignHandler())
	}

	return commonMiddle
}
