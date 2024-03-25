/*
 * @Author: yujiajie
 * @Date: 2024-03-25 17:38:27
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-25 17:50:06
 * @FilePath: /gateway/core/initialize/redis.go
 * @Description:
 */
package initialize

import (
	"fmt"
	"gateway/core/container"
	"gateway/core/stores/cache"
	"gateway/options"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	_redis *redis.Client
)

func SetupCache() error {
	cacheConfig := container.App.GetConfig("cache").(*options.CacheConfig)
	if cacheConfig != nil {
		if err := setupCache(cacheConfig); err != nil {
			return err
		}
	}
	return nil
}

func setupCache(cfg *options.CacheConfig) error {
	if cfg.Redis != nil {
		option := &redis.Options{
			Addr:         cfg.Redis.Addr,
			Password:     cfg.Redis.Password,
			DB:           cfg.Redis.Db,
			Protocol:     cfg.Redis.Protocol,
			DialTimeout:  time.Duration(cfg.Redis.DialTimeout) * time.Second,
			ReadTimeout:  time.Duration(cfg.Redis.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(cfg.Redis.WriteTimeout) * time.Second,
			PoolSize:     cfg.Redis.PoolSize,
			MinIdleConns: cfg.Redis.IdleConns,
		}
		r, err := cache.NewRedis(nil, option)
		if err != nil {
			fmt.Println("cache setup error", err)
			return err
		}
		_redis = r.GetClient()
		container.App.SetCache(r)
	}
	return nil
}
