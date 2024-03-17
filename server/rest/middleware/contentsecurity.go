/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-10 17:49:09
 * @LastEditTime: 2023-12-10 23:42:57
 * @LastEditors: yujiajie
 */
package middleware

import (
	"gateway/core/codec"
	"gateway/server/rest/header"
	"gateway/server/rest/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UnsignedCallback func(ctx *gin.Context, strict bool, code int)

func ContentSecurityHandler(decrypters map[string]codec.RsaDecrypter, tolerance time.Duration,
	strict bool, callbacks ...UnsignedCallback) gin.HandlerFunc {
	return LimitContentSecurityHandler(maxBytes, decrypters, tolerance, strict, callbacks...)
}

func LimitContentSecurityHandler(limitBytes int64, decrypters map[string]codec.RsaDecrypter,
	tolerance time.Duration, strict bool, callbacks ...UnsignedCallback) gin.HandlerFunc {
	if len(callbacks) == 0 {
		callbacks = append(callbacks, handleVerificationFailure)
	}

	return func(ctx *gin.Context) {
		switch ctx.Request.Method {
		case http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPut:
			securityHeader, err := security.ParseContentSecurity(ctx, decrypters)
			if err != nil {
				executeCallbacks(ctx, strict, header.CodeSignatureInvalidHeader, callbacks)
			} else if code := security.VerifySignature(ctx, securityHeader, tolerance); code != header.CodeSignaturePass {
				executeCallbacks(ctx, strict, code, callbacks)
			} else if ctx.Request.ContentLength > 0 && securityHeader.Encrypted() {
				LimitCryption(ctx, limitBytes, securityHeader.Key)
			} else {
				ctx.Next()
			}
		default:
			ctx.Next()
		}
	}
}

func executeCallbacks(ctx *gin.Context, strict bool, code int, callbacks []UnsignedCallback) {
	for _, callback := range callbacks {
		callback(ctx, strict, code)
	}
}

func handleVerificationFailure(ctx *gin.Context, strict bool, code int) {
	if strict {
		ctx.AbortWithStatus(http.StatusForbidden)
	} else {
		ctx.Next()
	}
}
