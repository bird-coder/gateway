/*
 * @Author: yujiajie
 * @Date: 2024-03-18 10:39:13
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-18 14:58:17
 * @FilePath: /gateway/options/rest.go
 * @Description:
 */
package options

import "time"

type MiddlewaresConf struct {
	Auth       bool `json:",default=true"`
	Trace      bool `json:",default=true"`
	Log        bool `json:",default=true"`
	Prometheus bool `json:",default=true"`
	Breaker    bool `json:",default=true"`
	Shedding   bool `json:",default=true"`
	Recover    bool `json:",default=true"`
	Metrics    bool `json:",default=true"`
	Gunzip     bool `json:",default=true"`
}

type PrivateKeyConf struct {
	Fingerprint string
	KeyFile     string
}

type SignatureConf struct {
	Strict      bool          `json:",default=false"`
	Expire      time.Duration `json:",default=1h"`
	PrivateKeys []PrivateKeyConf
}

type RestConf struct {
	Addr             string        `json:",default=0.0.0.0:8081"`
	CertFile         string        `json:",optional"`
	KeyFile          string        `json:",optional"`
	Verbose          bool          `json:",optional"`
	MaxConns         int           `json:",default=10000"`
	MaxBytes         int64         `json:",default=1048576"`
	Timeout          int64         `json:",default=3000"`
	CpuThreshold     int64         `json:",default=900,range=[0:1000]"`
	Signature        SignatureConf `json:",optional"`
	Middlewares      MiddlewaresConf
	TraceIgnorePaths []string `json:",optional"`
}
