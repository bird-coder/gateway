/*
 * @Author: yujiajie
 * @Date: 2024-03-18 10:42:00
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-18 10:42:09
 * @FilePath: /gateway/options/logger.go
 * @Description:
 */
package options

type LoggerConfig struct {
	LogLevel   string `json:"level"`
	LogPath    string `json:"logpath"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"age"`
	MaxBackups int    `json:"backups"`
	Compress   string `json:"compress"`
}
