/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-26 17:55:56
 * @LastEditTime: 2024-03-18 17:55:45
 * @LastEditors: yujiajie
 */
package middleware

import (
	"gateway/core/stat"
	"gateway/options"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Init(r *gin.Engine) {
	middlewareConfig := options.App.Gateway.RestConf.Middlewares
	r.Use(NoCache)
	r.Use(Options)
	r.Use(Secure)
	if middlewareConfig.Auth {
		handler := Authorize(options.App.Auth.Secret, WithPrevSecret(options.App.Auth.PrevSecret))
		r.Use(handler)
	}
	if middlewareConfig.Trace {
		r.Use(otelgin.Middleware("my-server"))
	}
	if middlewareConfig.Log {
		r.Use(LogHandler())
	}
	if middlewareConfig.Prometheus {

	}
	if middlewareConfig.Breaker {
		r.Use(Sentinel())
	}
	if middlewareConfig.Shedding {

	}
	if middlewareConfig.Recover {
		r.Use(RecoverHandler())
	}
	if middlewareConfig.Metrics {
		r.Use(MetricHandle(stat.NewMetrics("test")))
	}
	if middlewareConfig.Gunzip {
		r.Use(gzip.Gzip(gzip.DefaultCompression))
	}
}
