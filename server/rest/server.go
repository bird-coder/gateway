/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-19 20:28:11
 * @LastEditTime: 2024-03-18 11:57:54
 * @LastEditors: yujiajie
 */
package rest

import (
	"context"
	"fmt"
	"gateway/options"
	"gateway/server/rest/middleware"
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
	g := gin.New()
	middleware.Init(g)
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
	if err := s.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			fmt.Println("waiting for rest server finish...")
		}
		return err
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
	g.Any(route.Path, route.Handler)
}
