/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-19 19:26:34
 * @LastEditTime: 2024-03-18 16:06:33
 * @LastEditors: yujiajie
 */
package gateway

import (
	"gateway/options"
	"gateway/server/proxy"
	"gateway/server/rest"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

var (
	pathName = "path"
)

type Server struct {
	*rest.Server
	upstreams []options.Upstream
	proxys    []options.Proxy
}

func NewServer(c options.GatewayConf) *Server {
	s := &Server{
		upstreams: c.Upstreams,
		proxys:    c.Proxys,
		Server:    rest.NewServer(c.RestConf),
	}
	return s
}

func (s *Server) Start() error {
	if err := s.build(); err != nil {
		return err
	}
	if err := s.Server.Start(); err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	return s.Server.Stop()
}

func (s *Server) build() error {
	for _, p := range s.proxys {
		pro := proxy.NewServer(p)
		s.Server.AddRoute(rest.Route{
			Path:    p.Name + "/*" + pathName,
			Handler: s.buildHandler(pro),
		})
	}
	return nil
}

func (s *Server) buildHandler(proxy *httputil.ReverseProxy) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}