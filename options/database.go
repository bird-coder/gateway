/*
 * @Author: yujiajie
 * @Date: 2024-03-22 14:50:55
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-22 14:51:12
 * @FilePath: /gateway/options/database.go
 * @Description:
 */
package options

type MysqlConfig struct {
	Driver       string   `mapstructure:"driver"`
	IdleConns    int      `mapstructure:"idleConns"`
	OpenConns    int      `mapstructure:"openConns"`
	IdleTimeout  int64    `mapstructure:"idleTimeout"`
	AliveTimeout int64    `mapstructure:"aliveTimeout"`
	Cluster      bool     `mapstructure:"cluster"`
	Default      string   `mapstructure:"default"`
	Sources      []string `mapstructure:"sources"`
	Replicas     []string `mapstructure:"replicas"`
}
