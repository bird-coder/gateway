/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2024-03-17 14:07:26
 * @LastEditTime: 2024-03-18 15:05:29
 * @LastEditors: yujiajie
 */
package options

import (
	"github.com/spf13/viper"
)

var (
	App = new(AppConfig)
)

type AppConfig struct {
	Logger  *LoggerConfig
	Gateway *GatewayConf
	Auth    *AuthConfig
}

func (app *AppConfig) LoadConfig(configFile string) (err error) {
	viper.SetConfigFile(configFile)
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	if err = viper.Unmarshal(&app); err != nil {
		return
	}
	return
}
