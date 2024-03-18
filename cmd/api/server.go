/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-14 00:08:06
 * @LastEditTime: 2024-03-18 11:33:13
 * @LastEditors: yujiajie
 */
package api

import (
	"fmt"
	"gateway/core/constant"
	zlog "gateway/core/logger"
	"gateway/options"
	"gateway/server"

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
	if err := options.App.LoadConfig(configYml); err != nil {
		panic(err)
	}
	fmt.Println(options.App.Gateway, options.App.Logger)
	fmt.Println("starting api server...")
}

func run() error {
	zlog.NewLogger(options.App.Logger, constant.Dev.String())
	defer zlog.Sync()
	zlog.Info("forager server start")

	server.Init()

	return nil
}
