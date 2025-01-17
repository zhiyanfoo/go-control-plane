//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/extensions/filters/http/ratelimit/v3/rate_limit.proto

package ratelimitv3

import (
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	durationpb "github.com/planetscale/vtprotobuf/types/known/durationpb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *RateLimit) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimit) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimit) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.StatPrefix) > 0 {
		i -= len(m.StatPrefix)
		copy(dAtA[i:], m.StatPrefix)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.StatPrefix)))
		i--
		dAtA[i] = 0x6a
	}
	if m.StatusOnError != nil {
		if vtmsg, ok := interface{}(m.StatusOnError).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.StatusOnError)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x62
	}
	if len(m.ResponseHeadersToAdd) > 0 {
		for iNdEx := len(m.ResponseHeadersToAdd) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.ResponseHeadersToAdd[iNdEx]).(interface {
				MarshalToSizedBufferVTStrict([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.ResponseHeadersToAdd[iNdEx])
				if err != nil {
					return 0, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if m.RateLimitedStatus != nil {
		if vtmsg, ok := interface{}(m.RateLimitedStatus).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.RateLimitedStatus)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.DisableXEnvoyRatelimitedHeader {
		i--
		if m.DisableXEnvoyRatelimitedHeader {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.EnableXRatelimitHeaders != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.EnableXRatelimitHeaders))
		i--
		dAtA[i] = 0x40
	}
	if m.RateLimitService != nil {
		if vtmsg, ok := interface{}(m.RateLimitService).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.RateLimitService)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.RateLimitedAsResourceExhausted {
		i--
		if m.RateLimitedAsResourceExhausted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.FailureModeDeny {
		i--
		if m.FailureModeDeny {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Timeout != nil {
		size, err := (*durationpb.Duration)(m.Timeout).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RequestType) > 0 {
		i -= len(m.RequestType)
		copy(dAtA[i:], m.RequestType)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.RequestType)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Stage != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Stage))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitConfig_Action_SourceCluster) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitConfig_Action_SourceCluster) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action_SourceCluster) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitConfig_Action_DestinationCluster) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitConfig_Action_DestinationCluster) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action_DestinationCluster) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitConfig_Action_RequestHeaders) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitConfig_Action_RequestHeaders) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action_RequestHeaders) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.SkipIfAbsent {
		i--
		if m.SkipIfAbsent {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.DescriptorKey) > 0 {
		i -= len(m.DescriptorKey)
		copy(dAtA[i:], m.DescriptorKey)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.DescriptorKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.HeaderName) > 0 {
		i -= len(m.HeaderName)
		copy(dAtA[i:], m.HeaderName)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.HeaderName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitConfig_Action_RemoteAddress) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitConfig_Action_RemoteAddress) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action_RemoteAddress) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitConfig_Action_GenericKey) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitConfig_Action_GenericKey) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action_GenericKey) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.DescriptorKey) > 0 {
		i -= len(m.DescriptorKey)
		copy(dAtA[i:], m.DescriptorKey)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.DescriptorKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DescriptorValue) > 0 {
		i -= len(m.DescriptorValue)
		copy(dAtA[i:], m.DescriptorValue)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.DescriptorValue)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitConfig_Action_HeaderValueMatch) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitConfig_Action_HeaderValueMatch) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action_HeaderValueMatch) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Headers) > 0 {
		for iNdEx := len(m.Headers) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.Headers[iNdEx]).(interface {
				MarshalToSizedBufferVTStrict([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.Headers[iNdEx])
				if err != nil {
					return 0, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ExpectMatch {
		i--
		if m.ExpectMatch {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.DescriptorValue) > 0 {
		i -= len(m.DescriptorValue)
		copy(dAtA[i:], m.DescriptorValue)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.DescriptorValue)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitConfig_Action_MetaData) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitConfig_Action_MetaData) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action_MetaData) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.SkipIfAbsent {
		i--
		if m.SkipIfAbsent {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Source != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Source))
		i--
		dAtA[i] = 0x20
	}
	if len(m.DefaultValue) > 0 {
		i -= len(m.DefaultValue)
		copy(dAtA[i:], m.DefaultValue)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.DefaultValue)))
		i--
		dAtA[i] = 0x1a
	}
	if m.MetadataKey != nil {
		if vtmsg, ok := interface{}(m.MetadataKey).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.MetadataKey)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.DescriptorKey) > 0 {
		i -= len(m.DescriptorKey)
		copy(dAtA[i:], m.DescriptorKey)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.DescriptorKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitConfig_Action) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitConfig_Action) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if msg, ok := m.ActionSpecifier.(*RateLimitConfig_Action_Extension); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ActionSpecifier.(*RateLimitConfig_Action_Metadata); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ActionSpecifier.(*RateLimitConfig_Action_HeaderValueMatch_); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ActionSpecifier.(*RateLimitConfig_Action_GenericKey_); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ActionSpecifier.(*RateLimitConfig_Action_RemoteAddress_); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ActionSpecifier.(*RateLimitConfig_Action_RequestHeaders_); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ActionSpecifier.(*RateLimitConfig_Action_DestinationCluster_); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ActionSpecifier.(*RateLimitConfig_Action_SourceCluster_); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitConfig_Action_SourceCluster_) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action_SourceCluster_) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SourceCluster != nil {
		size, err := m.SourceCluster.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *RateLimitConfig_Action_DestinationCluster_) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action_DestinationCluster_) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DestinationCluster != nil {
		size, err := m.DestinationCluster.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *RateLimitConfig_Action_RequestHeaders_) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action_RequestHeaders_) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RequestHeaders != nil {
		size, err := m.RequestHeaders.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *RateLimitConfig_Action_RemoteAddress_) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action_RemoteAddress_) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RemoteAddress != nil {
		size, err := m.RemoteAddress.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *RateLimitConfig_Action_GenericKey_) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action_GenericKey_) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.GenericKey != nil {
		size, err := m.GenericKey.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x2a
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *RateLimitConfig_Action_HeaderValueMatch_) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action_HeaderValueMatch_) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HeaderValueMatch != nil {
		size, err := m.HeaderValueMatch.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x32
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *RateLimitConfig_Action_Metadata) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action_Metadata) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Metadata != nil {
		size, err := m.Metadata.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x42
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func (m *RateLimitConfig_Action_Extension) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Action_Extension) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Extension != nil {
		if vtmsg, ok := interface{}(m.Extension).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Extension)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x4a
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func (m *RateLimitConfig_Override_DynamicMetadata) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitConfig_Override_DynamicMetadata) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Override_DynamicMetadata) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.MetadataKey != nil {
		if vtmsg, ok := interface{}(m.MetadataKey).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.MetadataKey)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitConfig_Override) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitConfig_Override) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Override) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if msg, ok := m.OverrideSpecifier.(*RateLimitConfig_Override_DynamicMetadata_); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitConfig_Override_DynamicMetadata_) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig_Override_DynamicMetadata_) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DynamicMetadata != nil {
		size, err := m.DynamicMetadata.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *RateLimitConfig) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitConfig) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitConfig) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.Limit != nil {
		size, err := m.Limit.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Actions) > 0 {
		for iNdEx := len(m.Actions) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Actions[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DisableKey) > 0 {
		i -= len(m.DisableKey)
		copy(dAtA[i:], m.DisableKey)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.DisableKey)))
		i--
		dAtA[i] = 0x12
	}
	if m.Stage != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Stage))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitPerRoute) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitPerRoute) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitPerRoute) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RateLimits) > 0 {
		for iNdEx := len(m.RateLimits) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.RateLimits[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.OverrideOption != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.OverrideOption))
		i--
		dAtA[i] = 0x10
	}
	if m.VhRateLimits != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.VhRateLimits))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RateLimit) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Stage != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Stage))
	}
	l = len(m.RequestType)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Timeout != nil {
		l = (*durationpb.Duration)(m.Timeout).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.FailureModeDeny {
		n += 2
	}
	if m.RateLimitedAsResourceExhausted {
		n += 2
	}
	if m.RateLimitService != nil {
		if size, ok := interface{}(m.RateLimitService).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.RateLimitService)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.EnableXRatelimitHeaders != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.EnableXRatelimitHeaders))
	}
	if m.DisableXEnvoyRatelimitedHeader {
		n += 2
	}
	if m.RateLimitedStatus != nil {
		if size, ok := interface{}(m.RateLimitedStatus).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.RateLimitedStatus)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.ResponseHeadersToAdd) > 0 {
		for _, e := range m.ResponseHeadersToAdd {
			if size, ok := interface{}(e).(interface {
				SizeVT() int
			}); ok {
				l = size.SizeVT()
			} else {
				l = proto.Size(e)
			}
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.StatusOnError != nil {
		if size, ok := interface{}(m.StatusOnError).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.StatusOnError)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitConfig_Action_SourceCluster) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitConfig_Action_DestinationCluster) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitConfig_Action_RequestHeaders) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HeaderName)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.DescriptorKey)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.SkipIfAbsent {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitConfig_Action_RemoteAddress) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitConfig_Action_GenericKey) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DescriptorValue)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.DescriptorKey)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitConfig_Action_HeaderValueMatch) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DescriptorValue)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.ExpectMatch {
		n += 2
	}
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			if size, ok := interface{}(e).(interface {
				SizeVT() int
			}); ok {
				l = size.SizeVT()
			} else {
				l = proto.Size(e)
			}
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitConfig_Action_MetaData) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DescriptorKey)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.MetadataKey != nil {
		if size, ok := interface{}(m.MetadataKey).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.MetadataKey)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.DefaultValue)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Source != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Source))
	}
	if m.SkipIfAbsent {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitConfig_Action) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.ActionSpecifier.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitConfig_Action_SourceCluster_) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SourceCluster != nil {
		l = m.SourceCluster.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *RateLimitConfig_Action_DestinationCluster_) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DestinationCluster != nil {
		l = m.DestinationCluster.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *RateLimitConfig_Action_RequestHeaders_) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestHeaders != nil {
		l = m.RequestHeaders.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *RateLimitConfig_Action_RemoteAddress_) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RemoteAddress != nil {
		l = m.RemoteAddress.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *RateLimitConfig_Action_GenericKey_) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GenericKey != nil {
		l = m.GenericKey.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *RateLimitConfig_Action_HeaderValueMatch_) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HeaderValueMatch != nil {
		l = m.HeaderValueMatch.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *RateLimitConfig_Action_Metadata) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *RateLimitConfig_Action_Extension) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Extension != nil {
		if size, ok := interface{}(m.Extension).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Extension)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *RateLimitConfig_Override_DynamicMetadata) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MetadataKey != nil {
		if size, ok := interface{}(m.MetadataKey).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.MetadataKey)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitConfig_Override) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.OverrideSpecifier.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitConfig_Override_DynamicMetadata_) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DynamicMetadata != nil {
		l = m.DynamicMetadata.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *RateLimitConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Stage != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Stage))
	}
	l = len(m.DisableKey)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.Actions) > 0 {
		for _, e := range m.Actions {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.Limit != nil {
		l = m.Limit.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitPerRoute) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VhRateLimits != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.VhRateLimits))
	}
	if m.OverrideOption != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.OverrideOption))
	}
	if len(m.RateLimits) > 0 {
		for _, e := range m.RateLimits {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}
