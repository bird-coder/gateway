/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-14 00:08:06
 * @LastEditTime: 2024-03-17 14:11:02
 * @LastEditors: yujiajie
 */
package api

import (
	"fmt"
	zlog "gateway/core/logger"
	"gateway/options"

	"github.com/spf13/cobra"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API Server",
		Example:      "gateway server -c config/server.yml",
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
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func setup() {
	fmt.Println("starting api server...")
}

func run() error {
	app := &options.AppConfig{}
	if err := app.LoadConfig(configYml); err != nil {
		panic(err)
	}
	zlog.NewLogger(app.Logger, "develop")
	defer zlog.Sync()
	zlog.Info("forager server start")
	return nil
}
