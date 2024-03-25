/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-14 00:08:06
 * @LastEditTime: 2024-03-25 17:42:47
 * @LastEditors: yujiajie
 */
package api

import (
	"fmt"
	"gateway/core/container"
	"gateway/core/initialize"
	"gateway/options"
	"gateway/server"
	"gateway/service/mod"

	"github.com/spf13/cobra"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API Server",
		Example:      "gateway server -c config/server.yaml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/server.yaml", "Start server with provided configuration file")
}

func setup() {
	appConfig := new(options.AppConfig)
	if err := appConfig.LoadConfig(configYml); err != nil {
		panic(err)
	}
	container.App.SetConfig("gateway", appConfig.Gateway)
	container.App.SetConfig("auth", appConfig.Auth)
	container.App.SetConfig("loggers", appConfig.Loggers)
	container.App.SetConfig("cache", appConfig.Cache)
	container.App.SetConfig("databases", appConfig.Databases)
	initialize.SetupDB()
	initialize.SetupCache()
	initialize.SetupLog()
	mod.Sync()
	fmt.Println("starting api server...")
}

func run() error {
	defer container.App.SyncLogger()
	log := container.App.GetLogger("default")
	log.Info("gateway server start")

	server.Init()

	return nil
}
