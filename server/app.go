/*
 * @Author: yujiajie
 * @Date: 2024-03-18 11:04:02
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-18 11:27:00
 * @FilePath: /gateway/server/app.go
 * @Description:
 */
package server

import (
	"context"
	"fmt"
	"gateway/core/rungroup"
	"gateway/options"
	"gateway/server/gateway"
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
		cfg := options.App.Gateway
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
	if err := g.Run(); err != nil {
		fmt.Println(err)
	}
	return nil
}
