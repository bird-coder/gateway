/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-06 00:01:57
 * @LastEditTime: 2024-03-22 09:49:32
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

const (
	Auth_User = "user_id"
)

var (
	noDetailReason = "no detail reason"
	errNoClaims    = errors.New("no auth params")

	errAuthData = errors.New("auth data not valid")
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

		if err = checkAuth(data); err != nil {
			unauthorized(ctx, err)
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

func checkAuth(data map[string]interface{}) error {
	if _, ok := data[Auth_User]; !ok {
		return errAuthData
	}
	return nil
}
