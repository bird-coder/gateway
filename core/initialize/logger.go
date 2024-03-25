/*
 * @Author: yujiajie
 * @Date: 2024-03-25 17:38:45
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-25 17:41:55
 * @FilePath: /gateway/core/initialize/logger.go
 * @Description:
 */
package initialize

import (
	"gateway/core/constant"
	"gateway/core/container"
	"gateway/core/logger"
	"gateway/options"
)

func SetupLog() {
	logConfigs := container.App.GetConfig("loggers").(map[string]*options.LoggerConfig)
	for k, cfg := range logConfigs {
		log := logger.NewLogger(cfg, constant.Dev.String())
		container.App.SetLogger(k, log)
	}
}
