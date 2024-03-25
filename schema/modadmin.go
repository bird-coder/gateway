/*
 * @Author: yujiajie
 * @Date: 2024-03-25 11:09:40
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-25 11:13:46
 * @FilePath: /gateway/schema/test.go
 * @Description:
 */
package schema

import "time"

type AppConfig struct {
	ID        int       `gorm:"primary_key" json:"-"`
	AppId     string    `gorm:"unique" json:"app_id"`
	AppSecret string    `json:"app_secret"`
	AppName   string    `json:"app_name"`
	CreatedAt time.Time `json:"created_at"`
}

type AppModule struct {
	ID        int       `gorm:"primary_key" json:"-"`
	AppId     string    `gorm:"unique_index:idx_app_module" json:"app_id"`
	ModuleId  int       `gorm:"unique_index:idx_app_module" json:"module_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ModuleInfo struct {
	ID        int       `gorm:"primary_key" json:"-"`
	Name      string    `gorm:"unique" json:"name"`
	Note      string    `json:"note"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
