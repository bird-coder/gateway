/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2024-03-17 13:53:12
 * @LastEditTime: 2024-03-17 13:53:50
 * @LastEditors: yujiajie
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
