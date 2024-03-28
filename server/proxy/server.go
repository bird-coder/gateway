/*
 * @Author: yujiajie
 * @Date: 2024-03-18 15:35:38
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-28 16:00:59
 * @FilePath: /Gateway/server/proxy/server.go
 * @Description:
 */
package proxy

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"gateway/options"
	"gateway/server/rest/header"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	filterHeaders = []string{
		"Trailer", "Upgrade", "Proxy-Authorization", "Proxy-Authenticate",
		"Accept-Encoding",
	}
)

func NewServer(p options.Proxy) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director:       director(p),
		ModifyResponse: modifyResponse(),
		Transport:      transport(),
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
		r.Header.Set("X-Forwarded-For", target.Host)
		r.Header.Set("X-Forwarded-Proto", r.URL.Scheme)
		r.Header.Set("Accept", header.ApplicationJson)
		r.Host = target.Host
	}
}

func modifyResponse() func(*http.Response) error {
	return func(res *http.Response) error {
		if strings.Contains(res.Header.Get(header.ContentEncoding), header.GzipEncoding) && res.Body != nil {
			reader, err := gzip.NewReader(res.Body)
			if err != nil {
				return err
			}
			defer reader.Close()

			content, err := io.ReadAll(reader)
			if err != nil {
				return err
			}
			fmt.Println(string(content))

			res.Body = io.NopCloser(bytes.NewBuffer(content))

			res.Header.Del(header.ContentEncoding)
			res.Header.Del(header.ContentLength)
			res.Header.Set(header.ContentLength, strconv.Itoa(len(content)))
		}

		return nil
	}
}

func transport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          300,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

func filterHeader(r *http.Request) {
	for _, header := range filterHeaders {
		r.Header.Del(header)
	}
}
