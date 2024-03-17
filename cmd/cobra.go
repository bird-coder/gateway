/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-14 00:02:27
 * @LastEditTime: 2024-03-17 11:59:10
 * @LastEditors: yujiajie
 */
package cmd

import (
	"errors"
	"fmt"
	"gateway/cmd/api"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "gateway",
	Short:        "gateway",
	SilenceUsage: true,
	Long:         `gateway`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one args")
		}
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	fmt.Println("启动gateway")
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
