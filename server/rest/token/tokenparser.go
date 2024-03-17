/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-06 23:01:16
 * @LastEditTime: 2023-12-07 23:14:37
 * @LastEditors: yujiajie
 */
package token

import (
	"gateway/core/auth"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5/request"
)

const claimHistoryResetDuration = time.Hour * 24

type TokenParser struct {
	resetTime     time.Time
	resetDuration time.Duration
	history       sync.Map
}

func NewTokenParser() *TokenParser {
	parser := &TokenParser{
		resetTime:     time.Now(),
		resetDuration: claimHistoryResetDuration,
	}

	return parser
}

func (tp *TokenParser) ParseToken(ctx *gin.Context, secret, prevSecret string) (data map[string]interface{}, err error) {
	if len(prevSecret) > 0 {
		count := tp.loadCount(secret)
		prevCount := tp.loadCount(prevSecret)

		var first, second string
		if count > prevCount {
			first = secret
			second = prevSecret
		} else {
			first = prevSecret
			second = secret
		}

		data, err = tp.doParseToken(ctx, first)
		if err != nil {
			data, err = tp.doParseToken(ctx, second)
			if err != nil {
				return nil, err
			}

			tp.incrementCount(second)
		} else {
			tp.incrementCount(first)
		}
	} else {
		data, err = tp.doParseToken(ctx, secret)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (tp *TokenParser) doParseToken(ctx *gin.Context, secret string) (map[string]interface{}, error) {
	tokenString, err := request.AuthorizationHeaderExtractor.ExtractToken(ctx.Request)
	if err != nil {
		return nil, err
	}
	return auth.Auth(tokenString, []byte(secret))
}

func (tp *TokenParser) incrementCount(secret string) {
	if time.Since(tp.resetTime) > tp.resetDuration {
		tp.history.Range(func(key, value any) bool {
			tp.history.Delete(key)
			return true
		})
	}
	val, ok := tp.history.Load(secret)
	if ok {
		atomic.AddUint64(val.(*uint64), 1)
	} else {
		var count uint64 = 1
		tp.history.Store(secret, &count)
	}
}

func (tp *TokenParser) loadCount(secret string) uint64 {
	val, ok := tp.history.Load(secret)
	if ok {
		return *val.(*uint64)
	}
	return 0
}

func WithResetDuration(duration time.Duration) {

}
