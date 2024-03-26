/*
 * @Author: yujiajie
 * @Date: 2024-03-25 17:38:15
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-25 17:39:32
 * @FilePath: /gateway/core/initialize/mysq.go
 * @Description:
 */
package initialize

import (
	"gateway/core/container"
	"gateway/core/stores/database"
	"gateway/options"
)

func SetupDB() error {
	dbConfigs := container.App.GetConfig("databases").(map[string]*options.MysqlConfig)
	for k, cfg := range dbConfigs {
		db, err := database.NewDB(cfg)
		if err != nil {
			return err
		}
		container.App.SetDb(k, db)
	}
	return nil
}
