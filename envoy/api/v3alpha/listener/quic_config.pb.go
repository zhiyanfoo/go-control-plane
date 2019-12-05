// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/listener/quic_config.proto

package envoy_api_v3alpha_listener

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/api/annotations"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type QuicProtocolOptions struct {
	MaxConcurrentStreams   *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_concurrent_streams,json=maxConcurrentStreams,proto3" json:"max_concurrent_streams,omitempty"`
	IdleTimeout            *duration.Duration    `protobuf:"bytes,2,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	CryptoHandshakeTimeout *duration.Duration    `protobuf:"bytes,3,opt,name=crypto_handshake_timeout,json=cryptoHandshakeTimeout,proto3" json:"crypto_handshake_timeout,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}              `json:"-"`
	XXX_unrecognized       []byte                `json:"-"`
	XXX_sizecache          int32                 `json:"-"`
}

func (m *QuicProtocolOptions) Reset()         { *m = QuicProtocolOptions{} }
func (m *QuicProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*QuicProtocolOptions) ProtoMessage()    {}
func (*QuicProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_30110b89b5635b16, []int{0}
}

func (m *QuicProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuicProtocolOptions.Unmarshal(m, b)
}
func (m *QuicProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuicProtocolOptions.Marshal(b, m, deterministic)
}
func (m *QuicProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuicProtocolOptions.Merge(m, src)
}
func (m *QuicProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_QuicProtocolOptions.Size(m)
}
func (m *QuicProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_QuicProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_QuicProtocolOptions proto.InternalMessageInfo

func (m *QuicProtocolOptions) GetMaxConcurrentStreams() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConcurrentStreams
	}
	return nil
}

func (m *QuicProtocolOptions) GetIdleTimeout() *duration.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *QuicProtocolOptions) GetCryptoHandshakeTimeout() *duration.Duration {
	if m != nil {
		return m.CryptoHandshakeTimeout
	}
	return nil
}

func init() {
	proto.RegisterType((*QuicProtocolOptions)(nil), "envoy.api.v3alpha.listener.QuicProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/listener/quic_config.proto", fileDescriptor_30110b89b5635b16)
}

var fileDescriptor_30110b89b5635b16 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcd, 0x4a, 0x33, 0x31,
	0x18, 0x85, 0x69, 0x3f, 0xf8, 0x16, 0xa9, 0x20, 0x8c, 0x52, 0x6a, 0x91, 0x22, 0x82, 0x50, 0x51,
	0x12, 0x69, 0x77, 0xa2, 0x9b, 0xd6, 0x85, 0xae, 0xac, 0xad, 0xba, 0x1d, 0xde, 0xce, 0xa4, 0x6d,
	0x30, 0x93, 0x37, 0xe6, 0xa7, 0xb6, 0x77, 0xe0, 0x35, 0x78, 0x11, 0xde, 0x85, 0xf7, 0x25, 0x93,
	0x99, 0x56, 0xa1, 0x8a, 0xcb, 0x21, 0xcf, 0x79, 0x72, 0x26, 0x87, 0x9c, 0x72, 0x35, 0xc7, 0x25,
	0x03, 0x2d, 0xd8, 0xbc, 0x0b, 0x52, 0xcf, 0x80, 0x49, 0x61, 0x1d, 0x57, 0xdc, 0xb0, 0x67, 0x2f,
	0x92, 0x38, 0x41, 0x35, 0x11, 0x53, 0xaa, 0x0d, 0x3a, 0x8c, 0x9a, 0x81, 0xa6, 0xa0, 0x05, 0x2d,
	0x69, 0xba, 0xa2, 0x9b, 0xad, 0x29, 0xe2, 0x54, 0x72, 0x16, 0xc8, 0xb1, 0x9f, 0xb0, 0xd4, 0x1b,
	0x70, 0x02, 0x55, 0x91, 0xdd, 0x3c, 0x7f, 0x31, 0xa0, 0x35, 0x37, 0xb6, 0x3c, 0x3f, 0xf2, 0xa9,
	0x86, 0x50, 0x04, 0x94, 0x42, 0x17, 0xa2, 0x96, 0xcd, 0xb9, 0xb1, 0x02, 0x95, 0x50, 0x65, 0x85,
	0xc3, 0xf7, 0x2a, 0xd9, 0xb9, 0xf3, 0x22, 0x19, 0xe4, 0x5f, 0x09, 0xca, 0x5b, 0x1d, 0xc0, 0x68,
	0x48, 0xea, 0x19, 0x2c, 0xf2, 0xba, 0x89, 0x37, 0x86, 0x2b, 0x17, 0x5b, 0x67, 0x38, 0x64, 0xb6,
	0x51, 0x39, 0xa8, 0xb4, 0x6b, 0x9d, 0x7d, 0x5a, 0xdc, 0x4f, 0x57, 0xf7, 0xd3, 0x87, 0x1b, 0xe5,
	0xba, 0x9d, 0x47, 0x90, 0x9e, 0x0f, 0x77, 0x33, 0x58, 0xf4, 0xd7, 0xd1, 0x51, 0x91, 0x8c, 0x2e,
	0xc8, 0x96, 0x48, 0x25, 0x8f, 0x9d, 0xc8, 0x38, 0x7a, 0xd7, 0xa8, 0x06, 0xd3, 0xde, 0x86, 0xe9,
	0xaa, 0xfc, 0xd3, 0x61, 0x2d, 0xc7, 0xef, 0x0b, 0x3a, 0x1a, 0x91, 0x46, 0x62, 0x96, 0xda, 0x61,
	0x3c, 0x03, 0x95, 0xda, 0x19, 0x3c, 0x7d, 0x99, 0xfe, 0xfd, 0x65, 0xaa, 0x17, 0xd1, 0xeb, 0x55,
	0xb2, 0x94, 0x9e, 0x9f, 0xbd, 0x7d, 0xbc, 0xb6, 0x4e, 0xc8, 0xf1, 0xb7, 0x21, 0x3a, 0xeb, 0x0d,
	0xe8, 0x0f, 0x0f, 0xd3, 0xbb, 0x24, 0x6d, 0x81, 0x34, 0xf0, 0xda, 0xe0, 0x62, 0x49, 0x7f, 0xdf,
	0xb0, 0xb7, 0x9d, 0x0b, 0xfa, 0x61, 0xf1, 0xa0, 0x19, 0x54, 0xc6, 0xff, 0x43, 0xb7, 0xee, 0x67,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0xb8, 0xbb, 0x4e, 0x29, 0x02, 0x00, 0x00,
}