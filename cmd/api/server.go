/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-14 00:08:06
 * @LastEditTime: 2024-03-28 15:59:55
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
	container.App.SetEnv(appConfig.Environment)
	container.App.SetConfig("gateway", appConfig.Gateway)
	container.App.SetConfig("auth", appConfig.Auth)
	container.App.SetConfig("loggers", appConfig.Loggers)
	container.App.SetConfig("cache", appConfig.Cache)
	container.App.SetConfig("databases", appConfig.Databases)
	if err := initialize.SetupDB(); err != nil {
		panic(err)
	}
	if err := initialize.SetupCache(); err != nil {
		panic(err)
	}
	initialize.SetupLog()
	if err := mod.Sync(); err != nil {
		panic(err)
	}
	fmt.Println("starting api server...")
}

func run() error {
	defer container.App.SyncLogger()
	log := container.App.GetLogger("default")
	log.Info("gateway server start")

	server.Init()

	return nil
}
