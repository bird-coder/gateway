/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2024-03-17 14:07:26
 * @LastEditTime: 2024-03-25 16:00:41
 * @LastEditors: yujiajie
 */
package options

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Gateway   *GatewayConf
	Auth      *AuthConfig
	Cache     *CacheConfig
	Databases map[string]*MysqlConfig  `mapstructure:"databases"`
	Loggers   map[string]*LoggerConfig `mapstructure:"loggers"`
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
