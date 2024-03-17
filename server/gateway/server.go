/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-19 19:26:34
 * @LastEditTime: 2024-03-17 23:13:51
 * @LastEditors: yujiajie
 */
package gateway

import (
	"gateway/server/rest"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*rest.Server
	upstreams []Upstream
}

func NewServer(c GatewayConf) *Server {
	s := &Server{
		upstreams: c.Upstreams,
		Server:    rest.NewServer(c.RestConf),
	}
	return s
}

func (s *Server) Start() {
	s.build()
	s.Server.Start()
}

func (s *Server) Stop() {
	s.Server.Stop()
}

func (s *Server) build() error {
	for _, up := range s.upstreams {
		s.Server.AddRoute(rest.Route{
			Path:    up.Name + "/*path",
			Handler: s.buildHandler(up),
		})
	}
	return nil
}

func (s *Server) buildHandler(up Upstream) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		proxy := &httputil.ReverseProxy{
			Director: func(r *http.Request) {
				target, _ := url.Parse(up.Target)
				r.URL.Scheme = target.Scheme
				r.URL.Host = target.Host
				r.URL.Path = ctx.Param("path")
				r.Header.Set("X-Forwarded-Host", ctx.Request.Host)
				r.Header.Set("X-Forwarded-Proto", ctx.Request.URL.Scheme)
				r.Host = target.Host
			},
		}
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
