/*
 * @Author: yujiajie
 * @Date: 2024-03-22 14:43:02
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-22 14:43:54
 * @FilePath: /gateway/core/stores/cache/type.go
 * @Description:
 */
package cache

import "time"

type Cache interface {
	String() string
	Get(key string) (string, error)
	Set(key string, val interface{}, expire int) error
	Del(key string) error
	HGet(hk, key string) (string, error)
	HDel(hk, key string) error
	Increase(key string) error
	Decrease(key string) error
	Expire(key string, dur time.Duration) error
}
