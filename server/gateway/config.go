package gateway

import "gateway/server/rest"

type GatewayConf struct {
	rest.RestConf
	Upstreams []Upstream
}

// type RouteMapping struct {
// 	Method  string
// 	Path    string
// 	RpcPath string
// }

// type Upstream struct {
// 	Name      string
// 	Grpc      string
// 	ProtoSets []string
// 	Mappings  []RouteMapping
// }

type RouteMapping struct {
	Method     string
	Path       string
	RemotePath string
}

type Upstream struct {
	Name     string
	Target   string
	Mappings []RouteMapping
}
