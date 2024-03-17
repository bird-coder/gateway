/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-26 17:55:56
 * @LastEditTime: 2023-11-29 00:05:18
 * @LastEditors: yujiajie
 */
package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Init(r *gin.Engine) {
	r.Use(otelgin.Middleware("my-server"))
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(Sentinel())
}
