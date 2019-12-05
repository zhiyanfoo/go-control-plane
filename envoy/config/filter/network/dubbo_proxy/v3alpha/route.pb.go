// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/dubbo_proxy/v3alpha/route.proto

package envoy_config_filter_network_dubbo_proxy_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/api/annotations"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/route"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3alpha"
	v3alpha1 "github.com/envoyproxy/go-control-plane/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RouteConfiguration struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Interface            string   `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	Group                string   `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Routes               []*Route `protobuf:"bytes,5,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cb7b66cb50bd113, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *RouteConfiguration) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *RouteConfiguration) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	Match                *RouteMatch  `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	Route                *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cb7b66cb50bd113, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Route) GetRoute() *RouteAction {
	if m != nil {
		return m.Route
	}
	return nil
}

type RouteMatch struct {
	Method               *MethodMatch           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Headers              []*route.HeaderMatcher `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cb7b66cb50bd113, []int{2}
}

func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatch.Unmarshal(m, b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
}
func (m *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(m, src)
}
func (m *RouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteMatch.Size(m)
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

func (m *RouteMatch) GetMethod() *MethodMatch {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *RouteMatch) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

type RouteAction struct {
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	ClusterSpecifier     isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cb7b66cb50bd113, []int{3}
}

func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction.Unmarshal(m, b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
}
func (m *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(m, src)
}
func (m *RouteAction) XXX_Size() int {
	return xxx_messageInfo_RouteAction.Size(m)
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	WeightedClusters *route.WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *RouteAction) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *RouteAction) GetWeightedClusters() *route.WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
	}
}

type MethodMatch struct {
	Name                 *v3alpha.StringMatcher                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParamsMatch          map[uint32]*MethodMatch_ParameterMatchSpecifier `protobuf:"bytes,2,rep,name=params_match,json=paramsMatch,proto3" json:"params_match,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *MethodMatch) Reset()         { *m = MethodMatch{} }
func (m *MethodMatch) String() string { return proto.CompactTextString(m) }
func (*MethodMatch) ProtoMessage()    {}
func (*MethodMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cb7b66cb50bd113, []int{4}
}

func (m *MethodMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodMatch.Unmarshal(m, b)
}
func (m *MethodMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodMatch.Marshal(b, m, deterministic)
}
func (m *MethodMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodMatch.Merge(m, src)
}
func (m *MethodMatch) XXX_Size() int {
	return xxx_messageInfo_MethodMatch.Size(m)
}
func (m *MethodMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodMatch.DiscardUnknown(m)
}

var xxx_messageInfo_MethodMatch proto.InternalMessageInfo

func (m *MethodMatch) GetName() *v3alpha.StringMatcher {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MethodMatch) GetParamsMatch() map[uint32]*MethodMatch_ParameterMatchSpecifier {
	if m != nil {
		return m.ParamsMatch
	}
	return nil
}

type MethodMatch_ParameterMatchSpecifier struct {
	// Types that are valid to be assigned to ParameterMatchSpecifier:
	//	*MethodMatch_ParameterMatchSpecifier_ExactMatch
	//	*MethodMatch_ParameterMatchSpecifier_RangeMatch
	ParameterMatchSpecifier isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier `protobuf_oneof:"parameter_match_specifier"`
	XXX_NoUnkeyedLiteral    struct{}                                                      `json:"-"`
	XXX_unrecognized        []byte                                                        `json:"-"`
	XXX_sizecache           int32                                                         `json:"-"`
}

func (m *MethodMatch_ParameterMatchSpecifier) Reset()         { *m = MethodMatch_ParameterMatchSpecifier{} }
func (m *MethodMatch_ParameterMatchSpecifier) String() string { return proto.CompactTextString(m) }
func (*MethodMatch_ParameterMatchSpecifier) ProtoMessage()    {}
func (*MethodMatch_ParameterMatchSpecifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cb7b66cb50bd113, []int{4, 0}
}

func (m *MethodMatch_ParameterMatchSpecifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Unmarshal(m, b)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Marshal(b, m, deterministic)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Merge(m, src)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Size() int {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Size(m)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.DiscardUnknown(m)
}

var xxx_messageInfo_MethodMatch_ParameterMatchSpecifier proto.InternalMessageInfo

type isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier interface {
	isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier()
}

type MethodMatch_ParameterMatchSpecifier_ExactMatch struct {
	ExactMatch string `protobuf:"bytes,3,opt,name=exact_match,json=exactMatch,proto3,oneof"`
}

type MethodMatch_ParameterMatchSpecifier_RangeMatch struct {
	RangeMatch *v3alpha1.Int64Range `protobuf:"bytes,4,opt,name=range_match,json=rangeMatch,proto3,oneof"`
}

func (*MethodMatch_ParameterMatchSpecifier_ExactMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

func (*MethodMatch_ParameterMatchSpecifier_RangeMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

func (m *MethodMatch_ParameterMatchSpecifier) GetParameterMatchSpecifier() isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier {
	if m != nil {
		return m.ParameterMatchSpecifier
	}
	return nil
}

func (m *MethodMatch_ParameterMatchSpecifier) GetExactMatch() string {
	if x, ok := m.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_ExactMatch); ok {
		return x.ExactMatch
	}
	return ""
}

func (m *MethodMatch_ParameterMatchSpecifier) GetRangeMatch() *v3alpha1.Int64Range {
	if x, ok := m.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_RangeMatch); ok {
		return x.RangeMatch
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MethodMatch_ParameterMatchSpecifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MethodMatch_ParameterMatchSpecifier_ExactMatch)(nil),
		(*MethodMatch_ParameterMatchSpecifier_RangeMatch)(nil),
	}
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.config.filter.network.dubbo_proxy.v3alpha.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.config.filter.network.dubbo_proxy.v3alpha.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.config.filter.network.dubbo_proxy.v3alpha.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.config.filter.network.dubbo_proxy.v3alpha.RouteAction")
	proto.RegisterType((*MethodMatch)(nil), "envoy.config.filter.network.dubbo_proxy.v3alpha.MethodMatch")
	proto.RegisterMapType((map[uint32]*MethodMatch_ParameterMatchSpecifier)(nil), "envoy.config.filter.network.dubbo_proxy.v3alpha.MethodMatch.ParamsMatchEntry")
	proto.RegisterType((*MethodMatch_ParameterMatchSpecifier)(nil), "envoy.config.filter.network.dubbo_proxy.v3alpha.MethodMatch.ParameterMatchSpecifier")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/dubbo_proxy/v3alpha/route.proto", fileDescriptor_1cb7b66cb50bd113)
}

var fileDescriptor_1cb7b66cb50bd113 = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x41, 0x6f, 0xd3, 0x4a,
	0x10, 0xae, 0x93, 0xba, 0x7d, 0x1d, 0xbf, 0x27, 0xa5, 0xab, 0x27, 0xd5, 0x04, 0x54, 0x95, 0x20,
	0xa0, 0x5c, 0x6c, 0x91, 0xa2, 0xaa, 0x4a, 0x1b, 0x44, 0x13, 0x55, 0x2a, 0x48, 0x41, 0x95, 0x5b,
	0xd1, 0x43, 0x91, 0xa2, 0x6d, 0xb2, 0x49, 0x4c, 0x13, 0xaf, 0xb5, 0x5e, 0xa7, 0xcd, 0x91, 0x1b,
	0xe2, 0x88, 0x10, 0x07, 0xfe, 0x0f, 0x3f, 0x84, 0x1b, 0x3f, 0x01, 0x38, 0x21, 0xcf, 0xae, 0x13,
	0x97, 0x08, 0x89, 0xa4, 0xb7, 0xf5, 0xec, 0x7c, 0xdf, 0x7e, 0xdf, 0xcc, 0xee, 0x18, 0x76, 0x59,
	0x30, 0xe4, 0x23, 0xb7, 0xc5, 0x83, 0x8e, 0xdf, 0x75, 0x3b, 0x7e, 0x5f, 0x32, 0xe1, 0x06, 0x4c,
	0x5e, 0x72, 0x71, 0xe1, 0xb6, 0xe3, 0xf3, 0x73, 0xde, 0x0c, 0x05, 0xbf, 0x1a, 0xb9, 0xc3, 0x2d,
	0xda, 0x0f, 0x7b, 0xd4, 0x15, 0x3c, 0x96, 0xcc, 0x09, 0x05, 0x97, 0x9c, 0xb8, 0x08, 0x76, 0x14,
	0xd8, 0x51, 0x60, 0x47, 0x83, 0x9d, 0x0c, 0xd8, 0xd1, 0xe0, 0xe2, 0x3d, 0x75, 0x1a, 0x0d, 0xfd,
	0xeb, 0x7c, 0x59, 0xd6, 0xe2, 0x43, 0x95, 0x24, 0x47, 0x21, 0x73, 0x07, 0x54, 0xb6, 0x7a, 0x4c,
	0x8c, 0xb3, 0x23, 0x29, 0xfc, 0xa0, 0xab, 0x13, 0xd7, 0x33, 0x89, 0x63, 0x3a, 0x1a, 0x74, 0x53,
	0xa2, 0xfb, 0x71, 0x3b, 0xa4, 0x78, 0x18, 0x0d, 0x02, 0x2e, 0xa9, 0xf4, 0x79, 0x10, 0xb9, 0x43,
	0x26, 0x22, 0x9f, 0x07, 0x13, 0x9a, 0xb5, 0x21, 0xed, 0xfb, 0x6d, 0x2a, 0x99, 0x9b, 0x2e, 0xd4,
	0x46, 0xe9, 0x43, 0x0e, 0x88, 0x97, 0x08, 0xab, 0xa3, 0xc1, 0x58, 0x20, 0x03, 0x21, 0xb0, 0x18,
	0xd0, 0x01, 0xb3, 0x8d, 0x0d, 0x63, 0x73, 0xc5, 0xc3, 0x35, 0xb9, 0x03, 0x2b, 0x7e, 0x20, 0x99,
	0xe8, 0xd0, 0x16, 0xb3, 0x73, 0xb8, 0x31, 0x09, 0x90, 0xff, 0xc1, 0xec, 0x0a, 0x1e, 0x87, 0x76,
	0x1e, 0x77, 0xd4, 0x07, 0xb1, 0x61, 0x59, 0x6b, 0xb1, 0x17, 0x31, 0x9e, 0x7e, 0x92, 0x97, 0xb0,
	0x84, 0x05, 0x89, 0x6c, 0x73, 0x23, 0xbf, 0x69, 0x95, 0xb7, 0x9d, 0x19, 0x0b, 0xed, 0xa0, 0x6c,
	0x4f, 0xb3, 0x54, 0x5e, 0x7c, 0xfe, 0xf2, 0x6e, 0xfd, 0x00, 0xea, 0x7f, 0xcd, 0x52, 0x46, 0x96,
	0xc7, 0xce, 0xb4, 0xfb, 0xd2, 0xdb, 0x1c, 0x98, 0x18, 0x26, 0x67, 0x60, 0x62, 0x7b, 0xb0, 0x10,
	0x56, 0x79, 0x77, 0x3e, 0x91, 0x8d, 0x84, 0xa2, 0xf6, 0xcf, 0xcf, 0x9a, 0xf9, 0xde, 0xc8, 0x15,
	0x0c, 0x4f, 0x71, 0x92, 0xd7, 0x60, 0xa2, 0x78, 0x2c, 0xa6, 0x55, 0xde, 0x9b, 0x8f, 0x7c, 0xbf,
	0x95, 0x68, 0xce, 0xb2, 0x23, 0x69, 0xa5, 0x9a, 0x14, 0x64, 0x07, 0xb6, 0xe7, 0x2b, 0x48, 0xe9,
	0x9b, 0x01, 0x30, 0x11, 0x4f, 0x4e, 0x60, 0x69, 0xc0, 0x64, 0x8f, 0xb7, 0x75, 0x25, 0x66, 0x17,
	0xdb, 0x40, 0x38, 0xb2, 0x79, 0x9a, 0x8b, 0x3c, 0x83, 0xe5, 0x1e, 0xa3, 0x6d, 0x26, 0x22, 0x3b,
	0x87, 0xb7, 0xe0, 0x81, 0xa6, 0xa5, 0xa1, 0x3f, 0x06, 0xaa, 0x77, 0x73, 0x88, 0x79, 0x0d, 0xf5,
	0x56, 0xbc, 0x14, 0x56, 0xa9, 0x25, 0x2e, 0xab, 0xfa, 0x89, 0xcf, 0xec, 0x12, 0xe9, 0x4a, 0x5f,
	0x0d, 0xb0, 0x32, 0xa5, 0x24, 0x45, 0x58, 0x6e, 0xf5, 0xe3, 0x48, 0x32, 0xa1, 0xee, 0xff, 0xe1,
	0x82, 0x97, 0x06, 0xc8, 0x29, 0xac, 0x5e, 0x32, 0xbf, 0xdb, 0x93, 0xac, 0xdd, 0xd4, 0xb1, 0x48,
	0xf7, 0x6f, 0xf3, 0x8f, 0xda, 0x4f, 0x35, 0xa2, 0xae, 0x00, 0x87, 0x0b, 0x5e, 0xe1, 0xf2, 0x7a,
	0x28, 0xaa, 0xd4, 0x13, 0x23, 0x4f, 0x61, 0x6f, 0x3e, 0x23, 0xfa, 0x12, 0xd8, 0xb0, 0xaa, 0x45,
	0x35, 0xa3, 0x90, 0xb5, 0xfc, 0x8e, 0xcf, 0x04, 0xc9, 0x7f, 0xaf, 0x19, 0xa5, 0x4f, 0x26, 0x58,
	0x99, 0x0e, 0x90, 0x6a, 0xe6, 0x81, 0x5b, 0xe5, 0x47, 0x5a, 0x7a, 0x32, 0x66, 0x1c, 0x3d, 0x8f,
	0xc6, 0x1e, 0x8e, 0x71, 0x1e, 0xa5, 0x95, 0x57, 0xb3, 0x20, 0x84, 0x7f, 0x43, 0x2a, 0xe8, 0x20,
	0x6a, 0xaa, 0xe7, 0xa1, 0xba, 0xd7, 0xb8, 0xc9, 0xa5, 0x70, 0x8e, 0x90, 0x10, 0xd7, 0x07, 0x81,
	0x14, 0x23, 0xcf, 0x0a, 0x27, 0x91, 0xe2, 0x0f, 0x03, 0xd6, 0x30, 0x83, 0x49, 0x7d, 0x0d, 0x8e,
	0xc7, 0x0e, 0xef, 0x82, 0xc5, 0xae, 0x68, 0x4b, 0x6a, 0x31, 0x79, 0xdd, 0x34, 0xc0, 0xa0, 0xf2,
	0xbb, 0x0f, 0x16, 0x8e, 0x4d, 0x9d, 0xb2, 0x88, 0xb6, 0xd7, 0xb3, 0xb6, 0x53, 0x49, 0xcf, 0x03,
	0xb9, 0xfd, 0xc4, 0x4b, 0x72, 0x13, 0x0a, 0x04, 0x21, 0x45, 0xe5, 0x2c, 0xe9, 0xd0, 0x2b, 0x38,
	0x99, 0xb9, 0x43, 0x53, 0x26, 0xa7, 0x2d, 0xd4, 0x6e, 0xc3, 0xad, 0x30, 0xdd, 0x52, 0x1a, 0x27,
	0x1d, 0x2c, 0x7e, 0x34, 0xa0, 0xf0, 0x7b, 0x75, 0x48, 0x01, 0xf2, 0x17, 0x6c, 0x84, 0x0d, 0xfc,
	0xcf, 0x4b, 0x96, 0xe4, 0x0d, 0x98, 0x43, 0xda, 0x8f, 0xd3, 0x79, 0x72, 0x72, 0xf3, 0x6e, 0x4c,
	0x0b, 0xf5, 0xd4, 0x11, 0x95, 0xdc, 0x8e, 0x31, 0xff, 0x95, 0xcd, 0x9c, 0x53, 0x6b, 0x40, 0xd5,
	0xe7, 0x4a, 0xa9, 0xca, 0x9a, 0x51, 0x74, 0x4d, 0x4d, 0xa9, 0xa3, 0xe4, 0x6f, 0x76, 0x64, 0x9c,
	0x2f, 0xe1, 0x6f, 0x6d, 0xeb, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x12, 0xbf, 0x1a, 0xf4,
	0x07, 0x00, 0x00,
}