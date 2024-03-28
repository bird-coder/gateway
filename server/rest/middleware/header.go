/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-26 18:11:29
 * @LastEditTime: 2024-03-28 16:06:12
 * @LastEditors: yujiajie
 */
package middleware

import (
	"gateway/server/rest/header"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	origins = []string{}
)

func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

func Options(c *gin.Context) {
	checkAndSetHeader(c)
	if c.Request.Method != http.MethodOptions {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusNoContent)
	}
}

func Secure(c *gin.Context) {
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}
}

func checkAndSetHeader(c *gin.Context) {
	setVaryHeaders(c)
	if len(origins) == 0 {
		setHeader(c, header.AllOrigins)
		return
	}
	origin := c.GetHeader(header.OriginHeader)
	if isOriginAllowed(origins, origin) {
		setHeader(c, origin)
	}
}

func isOriginAllowed(allows []string, origin string) bool {
	origin = strings.ToLower(origin)

	for _, allow := range allows {
		if allow == header.AllOrigins {
			return true
		}

		allow = strings.ToLower(allow)
		if origin == allow {
			return true
		}

		if strings.HasSuffix(origin, "."+allow) {
			return true
		}
	}

	return false
}

func setHeader(c *gin.Context, origin string) {
	c.Header(header.AllowOrigin, origin)
	c.Header(header.AllowMethods, header.Methods)
	c.Header(header.AllowHeaders, header.AllowHeadersVal)
	c.Header(header.ExposeHeaders, header.ExposeHeadersVal)
	if origin != header.AllOrigins {
		c.Header(header.AllowCredentials, header.AllowTrue)
	}
	c.Header(header.MaxAgeHeader, header.MaxAgeHeaderVal)
}

func setVaryHeaders(c *gin.Context) {
	headers := make([]string, 0, 3)
	headers = append(headers, header.OriginHeader)
	if c.Request.Method == http.MethodOptions {
		headers = append(headers, header.RequestMethod, header.RequestHeaders)
	}
	c.Header(header.VaryHeader, strings.Join(headers, ", "))
}
