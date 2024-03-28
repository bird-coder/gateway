/*
 * @Author: yujiajie
 * @Date: 2024-03-22 16:10:59
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-28 16:02:10
 * @FilePath: /Gateway/core/container/type.go
 * @Description:
 */
package container

import (
	"gateway/core/logger"
	"gateway/core/stores/cache"

	"gorm.io/gorm"
)

type Container interface {
	SetDb(key string, db *gorm.DB)
	GetDb(key string) *gorm.DB
	GetAllDb() map[string]*gorm.DB

	SetConfig(key string, config any)
	GetConfig(key string) any

	SetCache(cache.Cache)
	GetCache() cache.Cache

	SetLogger(key string, log logger.Logger)
	GetLogger(key string) logger.Logger
	SyncLogger()

	SetEnv(env string)
	GetEnv() string
}
