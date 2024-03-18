/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-06 00:01:57
 * @LastEditTime: 2024-03-18 14:51:17
 * @LastEditors: yujiajie
 */
package middleware

import (
	"errors"
	"gateway/server/rest/token"
	"net/http"
	"net/http/httputil"

	zlog "gateway/core/logger"

	"github.com/gin-gonic/gin"
)

var (
	noDetailReason = "no detail reason"
	errNoClaims    = errors.New("no auth params")
)

type AuthorizeOptions struct {
	PrevSecret string
}

type AuthorizeOption func(opts *AuthorizeOptions)

func Authorize(secret string, opts ...AuthorizeOption) gin.HandlerFunc {
	var authOpts AuthorizeOptions
	for _, opt := range opts {
		opt(&authOpts)
	}

	parser := token.NewTokenParser()
	return func(ctx *gin.Context) {
		data, err := parser.ParseToken(ctx, secret, authOpts.PrevSecret)
		if err != nil {
			unauthorized(ctx, err)
			return
		}
		if data == nil {
			unauthorized(ctx, errNoClaims)
			return
		}

		for k, v := range data {
			ctx.Set(k, v)
		}
		ctx.Next()
	}
}

func WithPrevSecret(secret string) AuthorizeOption {
	return func(opts *AuthorizeOptions) {
		opts.PrevSecret = secret
	}
}

func unauthorized(ctx *gin.Context, err error) {
	details, _ := httputil.DumpRequest(ctx.Request, true)
	if err != nil {
		zlog.Error("authorize failed: %s\n=> %+v", err.Error(), string(details))
	} else {
		zlog.Error("authorize failed: %s\n=> %+v", noDetailReason, string(details))
	}
	ctx.AbortWithStatus(http.StatusUnauthorized)
}
