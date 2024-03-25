/*
 * @Author: yujiajie
 * @Date: 2024-03-18 11:04:02
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-25 15:36:40
 * @FilePath: /gateway/server/app.go
 * @Description:
 */
package server

import (
	"context"
	"fmt"
	"gateway/core/container"
	"gateway/core/rungroup"
	"gateway/options"
	"gateway/server/gateway"
	"gateway/service/mod"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Init() error {
	var g rungroup.Group
	{
		term := make(chan os.Signal, 1)
		signal.Notify(term, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		ticker := time.NewTicker(time.Second * 60)
		closeChan := make(chan struct{})

		g.Add(
			func() error {
				for {
					select {
					case s := <-term:
						fmt.Println("get a signal:", s.String())
						return nil
					case <-closeChan:
						return nil
					case <-ticker.C:
						fmt.Println("running...")
					}
				}
			},
			func(err error) {
				close(closeChan)
			},
		)
	}
	{
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		cfg := container.App.GetConfig("gateway").(*options.GatewayConf)
		server := gateway.NewServer(*cfg)
		server.Server.WithContext(ctx)
		g.Add(
			func() error {
				err := server.Start()
				return err
			},
			func(err error) {
				server.Stop()
				cancel()
			},
		)
	}
	{
		gateConfig := container.App.GetConfig("gateway").(*options.GatewayConf)
		middleConfig := gateConfig.RestConf.Middlewares
		if middleConfig.Sign {
			ticker := time.NewTicker(time.Minute * 5)
			closeChan := make(chan struct{})
			g.Add(
				func() error {
					for {
						select {
						case <-ticker.C:
							mod.Sync()
						case <-closeChan:
							return nil
						}
					}
				},
				func(err error) {
					close(closeChan)
				},
			)
		}
	}
	if err := g.Run(); err != nil {
		fmt.Println(err)
	}
	return nil
}
