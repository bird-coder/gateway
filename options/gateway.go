/*
 * @Author: yujiajie
 * @Date: 2024-03-18 09:21:33
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-18 15:56:40
 * @FilePath: /gateway/options/gateway.go
 * @Description:
 */
package options

type GatewayConf struct {
	RestConf
	Upstreams []Upstream `json:",optional"`
	Proxys    []Proxy    `json:",optional"`
}

type RouteMapping struct {
	Method  string
	Path    string
	RpcPath string
}

type Upstream struct {
	Name      string
	Grpc      string
	ProtoSets []string
	Mappings  []RouteMapping
}

type ProxyMapping struct {
	Method     string
	Path       string
	TargetPath string
}

type Proxy struct {
	Name      string
	EndPoints []string
	Current   int
	Mappings  []ProxyMapping
}
