/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-19 20:28:11
 * @LastEditTime: 2024-03-17 23:14:22
 * @LastEditors: yujiajie
 */
package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*http.Server
}

func NewServer(c RestConf) *Server {
	g := gin.New()
	s := &http.Server{
		Addr:           c.Addr,
		Handler:        g,
		ReadTimeout:    time.Duration(c.Timeout),
		WriteTimeout:   time.Duration(c.Timeout),
		MaxHeaderBytes: int(c.MaxBytes),
	}
	server := &Server{
		Server: s,
	}
	return server
}

func (s *Server) Start() {

}

func (s *Server) Stop() {

}

func (s *Server) AddRoute(route Route) {
	g := s.Handler.(*gin.Engine)
	g.Any(route.Path, route.Handler)
}
