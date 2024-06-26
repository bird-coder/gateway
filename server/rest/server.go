/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-19 20:28:11
 * @LastEditTime: 2024-03-28 15:59:14
 * @LastEditors: yujiajie
 */
package rest

import (
	"context"
	"fmt"
	"gateway/core/constant"
	"gateway/core/container"
	"gateway/options"
	"gateway/server/rest/middleware"
	"gateway/server/rest/router"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*http.Server
	cfg options.RestConf

	ctx context.Context
}

func NewServer(c options.RestConf) *Server {
	if container.App.GetEnv() == string(constant.Prod) {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	middleware.Init(g)
	router.Init(g)
	s := &http.Server{
		Addr:           c.Addr,
		Handler:        g,
		ReadTimeout:    time.Duration(c.Timeout) * time.Second,
		WriteTimeout:   time.Duration(c.Timeout) * time.Second,
		MaxHeaderBytes: int(c.MaxBytes),
	}
	server := &Server{
		Server: s,
		cfg:    c,
	}
	return server
}

func (s *Server) WithContext(ctx context.Context) {
	s.ctx = ctx
}

func (s *Server) Start() error {
	if len(s.cfg.CertFile) == 0 || len(s.cfg.KeyFile) == 0 {
		if err := s.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				fmt.Println("waiting for rest server finish...")
			}
			return err
		}
	} else {
		if err := s.ListenAndServeTLS(s.cfg.CertFile, s.cfg.KeyFile); err != nil {
			if err == http.ErrServerClosed {
				fmt.Println("waiting for rest server finish...")
			}
			return err
		}
	}

	return nil
}

func (s *Server) Stop() error {
	if err := s.Shutdown(s.ctx); err != nil {
		fmt.Println("rest server shutdown error:", err)
		return err
	}
	fmt.Println("rest server shutdown processed success")
	return nil
}

func (s *Server) AddRoute(route Route) {
	g := s.Handler.(*gin.Engine)
	group := g.Group(route.Group)
	group.Use(route.Middleware...)
	{
		group.Any(route.Path, route.Handler)
	}
}
