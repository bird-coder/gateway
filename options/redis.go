/*
 * @Author: yujiajie
 * @Date: 2024-03-22 14:51:17
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-22 14:51:32
 * @FilePath: /gateway/options/redis.go
 * @Description:
 */
package options

type CacheConfig struct {
	Redis *RedisDailConfig
}

type RedisDailConfig struct {
	DialTimeout  int64
	ReadTimeout  int64
	WriteTimeout int64
	Protocol     int
	Addr         string
	Db           int
	Password     string
	PoolSize     int
	IdleConns    int
	MaxRetry     int
}
