/*
 * @Author: yujiajie
 * @Date: 2024-03-18 15:35:38
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-18 16:15:37
 * @FilePath: /gateway/server/proxy/server.go
 * @Description:
 */
package proxy

import (
	"gateway/options"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

var (
	filterHeaders = []string{
		"Trailer", "Upgrade", "Proxy-Authorization", "Proxy-Authenticate",
	}
)

func NewServer(p options.Proxy) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: director(p),
	}
}

func director(p options.Proxy) func(*http.Request) {
	return func(r *http.Request) {
		endpoint := p.EndPoints[p.Current]
		p.Current = (p.Current + 1) % len(p.EndPoints)
		target, _ := url.Parse(endpoint)
		targetQuery := target.RawQuery

		re := regexp.MustCompile("^/" + p.Name + "(.*)")
		r.URL.Path = re.ReplaceAllString(r.URL.Path, "$1")

		r.URL.Scheme = target.Scheme
		r.URL.Host = target.Host

		for _, m := range p.Mappings {
			if m.Path == r.URL.Path {
				r.URL.Path = m.TargetPath
			}
		}

		if targetQuery == "" || r.URL.RawQuery == "" {
			r.URL.RawQuery = targetQuery + r.URL.RawQuery
		} else {
			r.URL.RawQuery = targetQuery + "&" + r.URL.RawQuery
		}

		filterHeader(r)
		r.Header.Set("X-Forwarded-Host", r.Host)
		r.Header.Set("X-Forwarded-Proto", r.URL.Scheme)
		r.Host = target.Host
	}
}

func filterHeader(r *http.Request) {
	for _, header := range filterHeaders {
		r.Header.Del(header)
	}
}
