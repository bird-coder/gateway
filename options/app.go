/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2024-03-17 14:07:26
 * @LastEditTime: 2024-03-17 14:09:39
 * @LastEditors: yujiajie
 */
package options

import "github.com/spf13/viper"

type AppConfig struct {
	Http   *HttpConfig
	Logger *LoggerConfig
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
