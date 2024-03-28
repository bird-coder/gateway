/*
 * @Author: yujiajie
 * @Date: 2024-03-22 16:10:37
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-28 16:02:04
 * @FilePath: /Gateway/core/container/container.go
 * @Description:
 */
package container

import (
	"gateway/core/logger"
	"gateway/core/stores/cache"
	"sync"

	"gorm.io/gorm"
)

var App = NewKernel()

type Kernel struct {
	dbs     map[string]*gorm.DB
	configs map[string]any
	cache   cache.Cache
	logs    map[string]logger.Logger
	env     string

	mu sync.RWMutex
}

func NewKernel() *Kernel {
	return &Kernel{
		dbs:     make(map[string]*gorm.DB),
		configs: make(map[string]any),
		logs:    make(map[string]logger.Logger),
	}
}

func (k *Kernel) SetDb(key string, db *gorm.DB) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.dbs[key] = db
}

func (k *Kernel) GetDb(key string) *gorm.DB {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.dbs[key]
}

func (k *Kernel) GetAllDb() map[string]*gorm.DB {
	return k.dbs
}

func (k *Kernel) SetConfig(key string, config any) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.configs[key] = config
}

func (k *Kernel) GetConfig(key string) any {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.configs[key]
}

func (k *Kernel) SetCache(c cache.Cache) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.cache = c
}

func (k *Kernel) GetCache() cache.Cache {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.cache
}

func (k *Kernel) SetLogger(key string, log logger.Logger) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.logs[key] = log
}

func (k *Kernel) GetLogger(key string) logger.Logger {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.logs[key]
}

func (k *Kernel) SyncLogger() {
	k.mu.Lock()
	defer k.mu.Unlock()
	for _, log := range k.logs {
		log.Sync()
	}
}

func (k *Kernel) SetEnv(env string) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.env = env
}

func (k *Kernel) GetEnv() string {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.env
}
