/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-26 15:10:54
 * @LastEditTime: 2023-12-10 23:44:43
 * @LastEditors: yujiajie
 */
package rest

import "time"

type MiddlewaresConf struct {
	Trace      bool `json:",default=true"`
	Log        bool `json:",default=true"`
	Prometheus bool `json:",default=true"`
	Breaker    bool `json:",default=true"`
	Shedding   bool `json:",default=true"`
	Timeout    bool `json:",default=true"`
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
