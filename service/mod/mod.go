/*
 * @Author: yujiajie
 * @Date: 2024-03-25 10:06:12
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-25 14:01:41
 * @FilePath: /gateway/service/mod/mod.go
 * @Description:
 */
package mod

import (
	"gateway/core/container"
	"gateway/schema"
	"sync"
)

var (
	appModules = make(map[string]AppModule)
	mux        sync.Mutex
)

type AppModule struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	AppName   string `json:"app_name"`
	Modules   map[string]string
}

func Sync() error {
	var apps []schema.AppConfig
	var modules []schema.ModuleInfo
	db := container.App.GetDb("alimatch/modadmin")
	if res := db.Find(&apps); res.Error != nil {
		return res.Error
	}
	for _, app := range apps {
		if res := db.Table("app_module").Joins("JOIN module_info ON app_module.module_id=module_info.id").
			Where("app_id = ?", app.AppId).Find(&modules); res.Error != nil {
			return res.Error
		}
		appModule := AppModule{
			AppId:     app.AppId,
			AppSecret: app.AppSecret,
			AppName:   app.AppName,
			Modules:   make(map[string]string),
		}
		for _, module := range modules {
			appModule.Modules[module.Name] = module.Url
		}
		SetAppModule(app.AppId, appModule)
	}
	return nil
}

func GetAppModule(appid string) AppModule {
	mux.Lock()
	defer mux.Unlock()
	return appModules[appid]
}

func SetAppModule(appid string, appModule AppModule) {
	mux.Lock()
	defer mux.Unlock()
	appModules[appid] = appModule
}
