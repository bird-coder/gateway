/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-15 23:15:21
 * @LastEditTime: 2024-03-17 14:02:50
 * @LastEditors: yujiajie
 */
package server

import (
	"context"
	"fmt"
	"gateway/options"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	*http.Server
	ctx context.Context
}

func NewHttp(ctx context.Context, c *options.HttpConfig) *HttpServer {
	g := gin.New()
	s := &http.Server{
		Addr:           c.Addr,
		Handler:        g,
		ReadTimeout:    time.Duration(c.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(c.WriteTimeout) * time.Second,
		MaxHeaderBytes: c.MaxHeaderBytes,
	}
	return &HttpServer{Server: s, ctx: ctx}
}

func (server *HttpServer) Run() error {
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			fmt.Println("waiting for http server finish...")
		}
		return err
	}
	return nil
}

func (server *HttpServer) Close() error {
	if err := server.Shutdown(server.ctx); err != nil {
		fmt.Println("http server shutdown error:", err)
		return err
	}
	fmt.Println("http server shutdown processed success")
	return nil
}
