/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-06 00:01:57
 * @LastEditTime: 2023-12-07 23:25:07
 * @LastEditors: yujiajie
 */
package middleware

import (
	"gateway/server/rest/token"
	"net/http"

	"github.com/gin-gonic/gin"
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
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if data == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
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
