// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routing/v1alpha2/destination_rule.proto

/*
Package istio_routing_v1alpha2 is a generated protocol buffer package.

It is generated from these files:
	routing/v1alpha2/destination_rule.proto
	routing/v1alpha2/envoy_filters.proto
	routing/v1alpha2/gateway.proto
	routing/v1alpha2/route_rule.proto

It has these top-level messages:
	DestinationRule
	TrafficPolicy
	Subset
	ConnectionPoolSettings
	OutlierDetection
	EnvoyFilters
	Gateway
	Server
	RouteRule
	Destination
	HTTPRoute
	TCPRoute
	HTTPMatchRequest
	DestinationWeight
	L4MatchAttributes
	HTTPRedirect
	HTTPRewrite
	StringMatch
	HTTPRetry
	CorsPolicy
	HTTPFaultInjection
	PortSelector
*/
package istio_routing_v1alpha2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Load balancing algorithm to use when
// forwarding traffic. These settings directly correlate to [load
// balancer
// types](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/load_balancing)
// supported by Envoy.
type TrafficPolicy_LBPolicy int32

const (
	// Simple round robin policy. Default
	TrafficPolicy_ROUND_ROBIN TrafficPolicy_LBPolicy = 0
	// The least request load balancer uses an O(1) algorithm which selects
	// two random healthy hosts and picks the host which has fewer active
	// requests.
	TrafficPolicy_LEAST_CONN TrafficPolicy_LBPolicy = 1
	// The random load balancer selects a random healthy host. The random
	// load balancer generally performs better than round robin if no health
	// checking policy is configured.
	TrafficPolicy_RANDOM TrafficPolicy_LBPolicy = 2
)

var TrafficPolicy_LBPolicy_name = map[int32]string{
	0: "ROUND_ROBIN",
	1: "LEAST_CONN",
	2: "RANDOM",
}
var TrafficPolicy_LBPolicy_value = map[string]int32{
	"ROUND_ROBIN": 0,
	"LEAST_CONN":  1,
	"RANDOM":      2,
}

func (x TrafficPolicy_LBPolicy) String() string {
	return proto.EnumName(TrafficPolicy_LBPolicy_name, int32(x))
}
func (TrafficPolicy_LBPolicy) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// DestinationRule defines policies that apply to traffic intended for a
// service after routing has occurred. These rules specify configuration
// for load balancing, connection pool size from the sidecar, and outlier
// detection settings to detect and evict unhealthy hosts from the load
// balancing pool. For example, a simple load balancing policy for the
// ratings service would look as follows:
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: DestinationRule
//     metadata:
//       name: bookinfo-ratings
//     spec:
//       name: ratings
//       trafficPolicy:
//         lbPolicy: LEAST_CONN
//
// Version specific DestinationRule can be specified by defining a named
// subset and overriding the settings specified at the service level. The
// following rule uses a round robin load balancing policy for all traffic
// going to a subset named testversion that is composed of endpoints (e.g.,
// pods) with labels (version:v3).
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: DestinationRule
//     metadata:
//       name: bookinfo-ratings
//     spec:
//       name: ratings
//       trafficPolicy:
//         lbPolicy: LEAST_CONN
//       subsets:
//       - name: testversion
//         labels:
//           version: v3
//         trafficPolicy:
//           lbPolicy: ROUND_ROBIN
//
// Note that policies specified for subsets will not take effect until
// a route rule explicitly sends traffic to this subset.
type DestinationRule struct {
	// REQUIRED. The destination address for traffic captured by this
	// rule.  Could be a DNS name with wildcard prefix or a CIDR
	// prefix. Depending on the platform, short-names can also be used
	// instead of a FQDN (i.e. has no dots in the name). In such a scenario,
	// the FQDN of the host would be derived based on the underlying
	// platform.
	//
	// For example on Kubernetes, when hosts contains a short name, Istio
	// will interpret the short name based on the namespace of the client
	// where rules are being applied. Thus, when a client in the "default"
	// namespace applies a rule containing a name "reviews", Istio will setup
	// routes to the "reviews.default.svc.cluster.local" service. However, if
	// a different name such as "reviews.sales" is used, it would be treated
	// as a FQDN during virtual host matching.  In Consul, a plain service
	// name would be resolved to the FQDN "reviews.service.consul".
	//
	// Note that the hosts field applies to both HTTP and TCP
	// services. Service inside the mesh, i.e. those found in the service
	// registry, must always be referred to using their alphanumeric
	// names. IP addresses or CIDR prefixes are allowed only for services
	// defined via the Gateway.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Traffic policies to apply (load balancing policy, connection pool
	// sizes, outlier detection).
	TrafficPolicy *TrafficPolicy `protobuf:"bytes,2,opt,name=traffic_policy,json=trafficPolicy" json:"traffic_policy,omitempty"`
	// One or more named sets that represent individual versions of a
	// service. Traffic policies can be overridden at subset level.
	Subsets []*Subset `protobuf:"bytes,3,rep,name=subsets" json:"subsets,omitempty"`
}

func (m *DestinationRule) Reset()                    { *m = DestinationRule{} }
func (m *DestinationRule) String() string            { return proto.CompactTextString(m) }
func (*DestinationRule) ProtoMessage()               {}
func (*DestinationRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DestinationRule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DestinationRule) GetTrafficPolicy() *TrafficPolicy {
	if m != nil {
		return m.TrafficPolicy
	}
	return nil
}

func (m *DestinationRule) GetSubsets() []*Subset {
	if m != nil {
		return m.Subsets
	}
	return nil
}

// Traffic policies to apply for a specific destination. See
// DestinationRule for examples.
type TrafficPolicy struct {
	// Upstream load balancing policy
	LbPolicy TrafficPolicy_LBPolicy `protobuf:"varint,1,opt,name=lb_policy,json=lbPolicy,enum=istio.routing.v1alpha2.TrafficPolicy_LBPolicy" json:"lb_policy,omitempty"`
	// Settings controlling the volume of connections to an upstream service
	ConnectionPool *ConnectionPoolSettings `protobuf:"bytes,2,opt,name=connection_pool,json=connectionPool" json:"connection_pool,omitempty"`
	// Settings controlling eviction of unhealthy hosts from the load balancing pool
	OutlierDetection *OutlierDetection `protobuf:"bytes,3,opt,name=outlier_detection,json=outlierDetection" json:"outlier_detection,omitempty"`
}

func (m *TrafficPolicy) Reset()                    { *m = TrafficPolicy{} }
func (m *TrafficPolicy) String() string            { return proto.CompactTextString(m) }
func (*TrafficPolicy) ProtoMessage()               {}
func (*TrafficPolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TrafficPolicy) GetLbPolicy() TrafficPolicy_LBPolicy {
	if m != nil {
		return m.LbPolicy
	}
	return TrafficPolicy_ROUND_ROBIN
}

func (m *TrafficPolicy) GetConnectionPool() *ConnectionPoolSettings {
	if m != nil {
		return m.ConnectionPool
	}
	return nil
}

func (m *TrafficPolicy) GetOutlierDetection() *OutlierDetection {
	if m != nil {
		return m.OutlierDetection
	}
	return nil
}

// A subset of endpoints of a service. Subsets can be used for scenarios
// like A/B testing, or routing to a specific version of a service. Refer
// to Route Rules documentation for examples on using subsets in these
// scenarios. In addition, traffic policies defined at the service-level
// can be overridden at a subset-level. The following rule uses a round
// robin load balancing policy for all traffic going to a subset named
// testversion that is composed of endpoints (e.g., pods) with labels
// (version:v3).
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: DestinationRule
//     metadata:
//       name: bookinfo-ratings
//     spec:
//       name: ratings
//       trafficPolicy:
//         lbPolicy: LEAST_CONN
//       subsets:
//       - name: testversion
//         labels:
//           version: v3
//         trafficPolicy:
//           lbPolicy: ROUND_ROBIN
//
// Note that policies specified for subsets will not take effect until
// a route rule explicitly sends traffic to this subset.
type Subset struct {
	// REQUIRED. name of the subset. The service name and the subset name can
	// be used for traffic splitting in a route rule.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// REQUIRED. Labels apply a filter over the endpoints of a service in the
	// service registry. See route rules for examples of usage.
	SourceLabels map[string]string `protobuf:"bytes,2,rep,name=source_labels,json=sourceLabels" json:"source_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Traffic policies that apply to this subset. Subsets inherit the
	// traffic policies specified at the DestinationRule level. Settings
	// specified at the subset level will override the corresponding settings
	// specified at the DestinationRule level.
	TrafficPolicy *TrafficPolicy `protobuf:"bytes,3,opt,name=traffic_policy,json=trafficPolicy" json:"traffic_policy,omitempty"`
}

func (m *Subset) Reset()                    { *m = Subset{} }
func (m *Subset) String() string            { return proto.CompactTextString(m) }
func (*Subset) ProtoMessage()               {}
func (*Subset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Subset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Subset) GetSourceLabels() map[string]string {
	if m != nil {
		return m.SourceLabels
	}
	return nil
}

func (m *Subset) GetTrafficPolicy() *TrafficPolicy {
	if m != nil {
		return m.TrafficPolicy
	}
	return nil
}

// Connection pool settings for an upstream host. The settings apply to
// each individual host in the upstream service.  See Envoy's [circuit
// breaker](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/circuit_breaking)
// for more details. Connection pool settings can be applied at the TCP
// level as well as at HTTP level.
//
// For example, the following rule sets a limit of 100 connections to redis
// service called myredissrv with a connect timeout of 30ms
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: DestinationRule
//     metadata:
//       name: bookinfo-redis
//     spec:
//       destination:
//         name: myredissrv
//       connectionPool:
//         tcp:
//           maxConnections: 100
//           connectTimeout: 30ms
//
type ConnectionPoolSettings struct {
	// Settings common to both HTTP and TCP upstream connections.
	Tcp *ConnectionPoolSettings_TCPSettings `protobuf:"bytes,1,opt,name=tcp" json:"tcp,omitempty"`
	// HTTP connection pool settings.
	Http *ConnectionPoolSettings_HTTPSettings `protobuf:"bytes,2,opt,name=http" json:"http,omitempty"`
}

func (m *ConnectionPoolSettings) Reset()                    { *m = ConnectionPoolSettings{} }
func (m *ConnectionPoolSettings) String() string            { return proto.CompactTextString(m) }
func (*ConnectionPoolSettings) ProtoMessage()               {}
func (*ConnectionPoolSettings) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ConnectionPoolSettings) GetTcp() *ConnectionPoolSettings_TCPSettings {
	if m != nil {
		return m.Tcp
	}
	return nil
}

func (m *ConnectionPoolSettings) GetHttp() *ConnectionPoolSettings_HTTPSettings {
	if m != nil {
		return m.Http
	}
	return nil
}

// Settings common to both HTTP and TCP upstream connections.
type ConnectionPoolSettings_TCPSettings struct {
	// Maximum number of HTTP/TCP connections to a destination host.
	MaxConnections int32 `protobuf:"varint,1,opt,name=max_connections,json=maxConnections" json:"max_connections,omitempty"`
	// TCP connection timeout.
	ConnectTimeout *google_protobuf.Duration `protobuf:"bytes,2,opt,name=connect_timeout,json=connectTimeout" json:"connect_timeout,omitempty"`
}

func (m *ConnectionPoolSettings_TCPSettings) Reset()         { *m = ConnectionPoolSettings_TCPSettings{} }
func (m *ConnectionPoolSettings_TCPSettings) String() string { return proto.CompactTextString(m) }
func (*ConnectionPoolSettings_TCPSettings) ProtoMessage()    {}
func (*ConnectionPoolSettings_TCPSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

func (m *ConnectionPoolSettings_TCPSettings) GetMaxConnections() int32 {
	if m != nil {
		return m.MaxConnections
	}
	return 0
}

func (m *ConnectionPoolSettings_TCPSettings) GetConnectTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.ConnectTimeout
	}
	return nil
}

// Settings applicable to HTTP1.1/HTTP2/GRPC connections.
type ConnectionPoolSettings_HTTPSettings struct {
	// Maximum number of pending HTTP requests to a destination. Default 1024.
	MaxPendingRequests int32 `protobuf:"varint,1,opt,name=max_pending_requests,json=maxPendingRequests" json:"max_pending_requests,omitempty"`
	// Maximum number of requests to a backend. Default 1024.
	MaxRequests int32 `protobuf:"varint,2,opt,name=max_requests,json=maxRequests" json:"max_requests,omitempty"`
	// Maximum number of requests per connection to a backend. Setting this
	// parameter to 1 disables keep alive.
	MaxRequestsPerConnection int32 `protobuf:"varint,3,opt,name=max_requests_per_connection,json=maxRequestsPerConnection" json:"max_requests_per_connection,omitempty"`
}

func (m *ConnectionPoolSettings_HTTPSettings) Reset()         { *m = ConnectionPoolSettings_HTTPSettings{} }
func (m *ConnectionPoolSettings_HTTPSettings) String() string { return proto.CompactTextString(m) }
func (*ConnectionPoolSettings_HTTPSettings) ProtoMessage()    {}
func (*ConnectionPoolSettings_HTTPSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 1}
}

func (m *ConnectionPoolSettings_HTTPSettings) GetMaxPendingRequests() int32 {
	if m != nil {
		return m.MaxPendingRequests
	}
	return 0
}

func (m *ConnectionPoolSettings_HTTPSettings) GetMaxRequests() int32 {
	if m != nil {
		return m.MaxRequests
	}
	return 0
}

func (m *ConnectionPoolSettings_HTTPSettings) GetMaxRequestsPerConnection() int32 {
	if m != nil {
		return m.MaxRequestsPerConnection
	}
	return 0
}

// A Circuit breaker implementation that tracks the status of each
// individual host in the upstream service.  While currently applicable to
// only HTTP services, future versions will support opaque TCP services as
// well. For HTTP services, hosts that continually return errors for API
// calls are ejected from the pool for a pre-defined period of time. See
// Envoy's [outlier
// detection](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/outlier)
// for more details.
//
// The following rule sets a connection pool size of 100 connections and
// 1000 concurrent requests, with no more than 10 req/connection to
// "reviews" service. In addition, it configures upstream hosts to be
// scanned every 5 mins, such that any host that fails 7 consecutive times
// with 5XX error code will be ejected for 15 minutes.
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: DestinationRule
//     metadata:
//       name: reviews-cb-policy
//     spec:
//       destination:
//         name: reviews
//       connectionPool:
//         tcp:
//           maxConnections: 100
//         http:
//           maxRequests: 1000
//           maxRequestsPerConnection: 10
//       outlierDetection:
//         http:
//           consecutiveErrors: 7
//           interval: 5m
//           baseEjectionTime: 15m
//
type OutlierDetection struct {
	// Settings for HTTP1.1/HTTP2/GRPC connections.
	Http *OutlierDetection_HTTPSettings `protobuf:"bytes,1,opt,name=http" json:"http,omitempty"`
}

func (m *OutlierDetection) Reset()                    { *m = OutlierDetection{} }
func (m *OutlierDetection) String() string            { return proto.CompactTextString(m) }
func (*OutlierDetection) ProtoMessage()               {}
func (*OutlierDetection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OutlierDetection) GetHttp() *OutlierDetection_HTTPSettings {
	if m != nil {
		return m.Http
	}
	return nil
}

// Outlier detection settings for HTTP1.1/HTTP2/GRPC connections.
type OutlierDetection_HTTPSettings struct {
	// Number of 5XX errors before a host is ejected from the connection
	// pool. Defaults to 5.
	ConsecutiveErrors int32 `protobuf:"varint,1,opt,name=consecutive_errors,json=consecutiveErrors" json:"consecutive_errors,omitempty"`
	// Time interval between ejection sweep analysis. format:
	// 1h/1m/1s/1ms. MUST BE >=1ms. Default is 10s.
	Interval *google_protobuf.Duration `protobuf:"bytes,2,opt,name=interval" json:"interval,omitempty"`
	// Minimum ejection duration. A host will remain ejected for a period
	// equal to the product of minimum ejection duration and the number of
	// times the host has been ejected. This technique allows the system to
	// automatically increase the ejection period for unhealthy upstream
	// servers. format: 1h/1m/1s/1ms. MUST BE >=1ms. Default is 30s.
	BaseEjectionTime *google_protobuf.Duration `protobuf:"bytes,3,opt,name=base_ejection_time,json=baseEjectionTime" json:"base_ejection_time,omitempty"`
	// Maximum % of hosts in the load balancing pool for the upstream
	// service that can be ejected. Defaults to 10%.
	MaxEjectionPercent int32 `protobuf:"varint,4,opt,name=max_ejection_percent,json=maxEjectionPercent" json:"max_ejection_percent,omitempty"`
}

func (m *OutlierDetection_HTTPSettings) Reset()         { *m = OutlierDetection_HTTPSettings{} }
func (m *OutlierDetection_HTTPSettings) String() string { return proto.CompactTextString(m) }
func (*OutlierDetection_HTTPSettings) ProtoMessage()    {}
func (*OutlierDetection_HTTPSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

func (m *OutlierDetection_HTTPSettings) GetConsecutiveErrors() int32 {
	if m != nil {
		return m.ConsecutiveErrors
	}
	return 0
}

func (m *OutlierDetection_HTTPSettings) GetInterval() *google_protobuf.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *OutlierDetection_HTTPSettings) GetBaseEjectionTime() *google_protobuf.Duration {
	if m != nil {
		return m.BaseEjectionTime
	}
	return nil
}

func (m *OutlierDetection_HTTPSettings) GetMaxEjectionPercent() int32 {
	if m != nil {
		return m.MaxEjectionPercent
	}
	return 0
}

func init() {
	proto.RegisterType((*DestinationRule)(nil), "istio.routing.v1alpha2.DestinationRule")
	proto.RegisterType((*TrafficPolicy)(nil), "istio.routing.v1alpha2.TrafficPolicy")
	proto.RegisterType((*Subset)(nil), "istio.routing.v1alpha2.Subset")
	proto.RegisterType((*ConnectionPoolSettings)(nil), "istio.routing.v1alpha2.ConnectionPoolSettings")
	proto.RegisterType((*ConnectionPoolSettings_TCPSettings)(nil), "istio.routing.v1alpha2.ConnectionPoolSettings.TCPSettings")
	proto.RegisterType((*ConnectionPoolSettings_HTTPSettings)(nil), "istio.routing.v1alpha2.ConnectionPoolSettings.HTTPSettings")
	proto.RegisterType((*OutlierDetection)(nil), "istio.routing.v1alpha2.OutlierDetection")
	proto.RegisterType((*OutlierDetection_HTTPSettings)(nil), "istio.routing.v1alpha2.OutlierDetection.HTTPSettings")
	proto.RegisterEnum("istio.routing.v1alpha2.TrafficPolicy_LBPolicy", TrafficPolicy_LBPolicy_name, TrafficPolicy_LBPolicy_value)
}

func init() { proto.RegisterFile("routing/v1alpha2/destination_rule.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x4f, 0xdb, 0x4a,
	0x14, 0x7d, 0x71, 0x80, 0x07, 0x37, 0x90, 0x84, 0x11, 0x42, 0x79, 0x79, 0x12, 0xe2, 0x45, 0x7a,
	0x22, 0x9b, 0x3a, 0x34, 0x15, 0x2a, 0xa2, 0xaa, 0x2a, 0x20, 0x51, 0x8b, 0x9a, 0x26, 0x91, 0x13,
	0xd4, 0xa5, 0x65, 0x3b, 0x97, 0xe0, 0x76, 0xe2, 0x71, 0x67, 0xc6, 0x51, 0xe8, 0xb2, 0x3f, 0xa5,
	0xfb, 0xfe, 0x8c, 0xae, 0xbb, 0xed, 0xbf, 0xe8, 0x5f, 0xa8, 0x3c, 0xfe, 0x88, 0xa1, 0x81, 0x42,
	0x77, 0xf6, 0xdc, 0x73, 0xce, 0xdc, 0x3b, 0xe7, 0xcc, 0xc0, 0x1e, 0x67, 0x81, 0x74, 0xbd, 0x71,
	0x63, 0xfa, 0xd8, 0xa2, 0xfe, 0xa5, 0xd5, 0x6c, 0x8c, 0x50, 0x48, 0xd7, 0xb3, 0xa4, 0xcb, 0x3c,
	0x93, 0x07, 0x14, 0x75, 0x9f, 0x33, 0xc9, 0xc8, 0xb6, 0x2b, 0xa4, 0xcb, 0xf4, 0x18, 0xae, 0x27,
	0xf0, 0xea, 0xce, 0x98, 0xb1, 0x31, 0xc5, 0x86, 0x42, 0xd9, 0xc1, 0x45, 0x63, 0x14, 0x70, 0x45,
	0x8e, 0x78, 0xb5, 0x2f, 0x39, 0x28, 0xb5, 0xe6, 0x92, 0x46, 0x40, 0x91, 0x10, 0x58, 0xf2, 0xac,
	0x09, 0x56, 0x72, 0xbb, 0xb9, 0xfa, 0x9a, 0xa1, 0xbe, 0x49, 0x07, 0x8a, 0x92, 0x5b, 0x17, 0x17,
	0xae, 0x63, 0xfa, 0x8c, 0xba, 0xce, 0x55, 0x45, 0xdb, 0xcd, 0xd5, 0x0b, 0xcd, 0xff, 0xf5, 0xc5,
	0x1b, 0xeb, 0xc3, 0x08, 0xdd, 0x57, 0x60, 0x63, 0x43, 0x66, 0x7f, 0xc9, 0x21, 0xfc, 0x2d, 0x02,
	0x5b, 0xa0, 0x14, 0x95, 0xfc, 0x6e, 0xbe, 0x5e, 0x68, 0xee, 0xdc, 0x26, 0x33, 0x50, 0x30, 0x23,
	0x81, 0xd7, 0xbe, 0x6a, 0xb0, 0x71, 0x4d, 0x9a, 0xbc, 0x86, 0x35, 0x6a, 0x27, 0x4d, 0x85, 0x2d,
	0x17, 0x9b, 0xfa, 0xbd, 0x9a, 0xd2, 0x3b, 0x27, 0x71, 0x77, 0xab, 0xd4, 0x8e, 0xc5, 0xde, 0x42,
	0xc9, 0x61, 0x9e, 0x87, 0x8e, 0x3a, 0x5f, 0x9f, 0x31, 0x1a, 0xcf, 0x79, 0xab, 0xe4, 0x69, 0x0a,
	0xef, 0x33, 0x46, 0x07, 0x28, 0xc3, 0xb2, 0x30, 0x8a, 0xce, 0xb5, 0x75, 0x72, 0x0e, 0x9b, 0x2c,
	0x90, 0xd4, 0x45, 0x6e, 0x8e, 0x50, 0x46, 0x85, 0x4a, 0x5e, 0x49, 0xd7, 0x6f, 0x93, 0xee, 0x45,
	0x84, 0x56, 0x82, 0x37, 0xca, 0xec, 0xc6, 0x4a, 0xed, 0x29, 0xac, 0x26, 0x53, 0x90, 0x12, 0x14,
	0x8c, 0xde, 0x79, 0xb7, 0x65, 0x1a, 0xbd, 0x93, 0xb3, 0x6e, 0xf9, 0x2f, 0x52, 0x04, 0xe8, 0xb4,
	0x8f, 0x07, 0x43, 0xf3, 0xb4, 0xd7, 0xed, 0x96, 0x73, 0x04, 0x60, 0xc5, 0x38, 0xee, 0xb6, 0x7a,
	0x6f, 0xca, 0x5a, 0xed, 0x93, 0x06, 0x2b, 0xd1, 0xd9, 0x2e, 0xb4, 0xfb, 0x1c, 0x36, 0x04, 0x0b,
	0xb8, 0x83, 0x26, 0xb5, 0x6c, 0xa4, 0xa2, 0xa2, 0x29, 0x9b, 0xf6, 0xef, 0xb6, 0x49, 0x1f, 0x28,
	0x4e, 0x47, 0x51, 0xda, 0x9e, 0xe4, 0x57, 0xc6, 0xba, 0xc8, 0x2c, 0x2d, 0x48, 0x51, 0xfe, 0xcf,
	0x53, 0x54, 0x7d, 0x01, 0x9b, 0xbf, 0x6c, 0x48, 0xca, 0x90, 0x7f, 0x8f, 0x57, 0xf1, 0x30, 0xe1,
	0x27, 0xd9, 0x82, 0xe5, 0xa9, 0x45, 0x03, 0x54, 0x4e, 0xae, 0x19, 0xd1, 0xcf, 0x91, 0x76, 0x98,
	0xab, 0x7d, 0xcf, 0xc3, 0xf6, 0x62, 0xff, 0x48, 0x07, 0xf2, 0xd2, 0xf1, 0x95, 0x4c, 0xa1, 0x79,
	0xf4, 0x30, 0xf3, 0xf5, 0xe1, 0x69, 0x3f, 0x0d, 0x42, 0x28, 0x43, 0x7a, 0xb0, 0x74, 0x29, 0xa5,
	0x1f, 0x67, 0xe9, 0xd9, 0x03, 0xe5, 0x5e, 0x0d, 0x87, 0x73, 0x3d, 0x25, 0x54, 0xfd, 0x08, 0x85,
	0xcc, 0x26, 0x64, 0x0f, 0x4a, 0x13, 0x6b, 0x66, 0xce, 0x33, 0x27, 0x54, 0xe7, 0xcb, 0x46, 0x71,
	0x62, 0xcd, 0xe6, 0xaa, 0x82, 0x9c, 0xa4, 0xf9, 0x36, 0xa5, 0x3b, 0x41, 0x16, 0xc8, 0xb8, 0xa7,
	0x7f, 0xf4, 0xe8, 0xa1, 0xd0, 0x93, 0x87, 0x42, 0x6f, 0xc5, 0x0f, 0x45, 0x1a, 0xe5, 0x61, 0x44,
	0xa8, 0x7e, 0xce, 0xc1, 0x7a, 0xb6, 0x25, 0xb2, 0x0f, 0x5b, 0xe1, 0xee, 0x3e, 0x7a, 0x23, 0xd7,
	0x1b, 0x9b, 0x1c, 0x3f, 0x04, 0x28, 0x64, 0xd2, 0x02, 0x99, 0x58, 0xb3, 0x7e, 0x54, 0x32, 0xe2,
	0x0a, 0xf9, 0x0f, 0xd6, 0x43, 0x46, 0x8a, 0xd4, 0x14, 0xb2, 0x30, 0xb1, 0x66, 0x29, 0xe4, 0x39,
	0xfc, 0x9b, 0x85, 0x98, 0x3e, 0xf2, 0xcc, 0x7c, 0x2a, 0x37, 0xcb, 0x46, 0x25, 0xc3, 0xe8, 0x23,
	0x9f, 0x4f, 0x5a, 0xfb, 0xa6, 0x41, 0xf9, 0xe6, 0xfd, 0x21, 0x67, 0xb1, 0x0d, 0x91, 0xab, 0x07,
	0xf7, 0xbd, 0x77, 0x8b, 0x0c, 0xf8, 0x71, 0xf3, 0x10, 0x1e, 0x01, 0x71, 0x98, 0x27, 0xd0, 0x09,
	0xa4, 0x3b, 0x45, 0x13, 0x39, 0x67, 0x3c, 0x39, 0x82, 0xcd, 0x4c, 0xa5, 0xad, 0x0a, 0xe4, 0x00,
	0x56, 0x5d, 0x4f, 0x22, 0x9f, 0x5a, 0xf4, 0xf7, 0x0e, 0xa4, 0x50, 0xf2, 0x12, 0x88, 0x6d, 0x09,
	0x34, 0xf1, 0x5d, 0xfc, 0x44, 0x85, 0x2e, 0xc6, 0x97, 0xe8, 0x0e, 0x81, 0x72, 0x48, 0x6a, 0xc7,
	0x9c, 0xd0, 0xc7, 0xc4, 0xb3, 0x54, 0xc7, 0x47, 0xee, 0xa0, 0x27, 0x2b, 0x4b, 0xa9, 0x67, 0x09,
	0xbc, 0x1f, 0x55, 0xec, 0x15, 0x25, 0xfb, 0xe4, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x81,
	0x23, 0x8a, 0x93, 0x06, 0x00, 0x00,
}
