//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/extensions/filters/http/fault/v3/fault.proto

package faultv3

import (
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	structpb "github.com/planetscale/vtprotobuf/types/known/structpb"
	wrapperspb "github.com/planetscale/vtprotobuf/types/known/wrapperspb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *FaultAbort_HeaderAbort) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *FaultAbort_HeaderAbort) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *FaultAbort_HeaderAbort) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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

func (m *FaultAbort) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *FaultAbort) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *FaultAbort) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if msg, ok := m.ErrorType.(*FaultAbort_GrpcStatus); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ErrorType.(*FaultAbort_HeaderAbort_); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if m.Percentage != nil {
		if vtmsg, ok := interface{}(m.Percentage).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Percentage)
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
	if msg, ok := m.ErrorType.(*FaultAbort_HttpStatus); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *FaultAbort_HttpStatus) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *FaultAbort_HttpStatus) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(m.HttpStatus))
	i--
	dAtA[i] = 0x10
	return len(dAtA) - i, nil
}
func (m *FaultAbort_HeaderAbort_) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *FaultAbort_HeaderAbort_) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HeaderAbort != nil {
		size, err := m.HeaderAbort.MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *FaultAbort_GrpcStatus) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *FaultAbort_GrpcStatus) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(m.GrpcStatus))
	i--
	dAtA[i] = 0x28
	return len(dAtA) - i, nil
}
func (m *HTTPFault) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *HTTPFault) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *HTTPFault) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.FilterMetadata != nil {
		size, err := (*structpb.Struct)(m.FilterMetadata).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if m.DisableDownstreamClusterStats {
		i--
		if m.DisableDownstreamClusterStats {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x78
	}
	if len(m.AbortGrpcStatusRuntime) > 0 {
		i -= len(m.AbortGrpcStatusRuntime)
		copy(dAtA[i:], m.AbortGrpcStatusRuntime)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.AbortGrpcStatusRuntime)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.ResponseRateLimitPercentRuntime) > 0 {
		i -= len(m.ResponseRateLimitPercentRuntime)
		copy(dAtA[i:], m.ResponseRateLimitPercentRuntime)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ResponseRateLimitPercentRuntime)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.MaxActiveFaultsRuntime) > 0 {
		i -= len(m.MaxActiveFaultsRuntime)
		copy(dAtA[i:], m.MaxActiveFaultsRuntime)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.MaxActiveFaultsRuntime)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.AbortHttpStatusRuntime) > 0 {
		i -= len(m.AbortHttpStatusRuntime)
		copy(dAtA[i:], m.AbortHttpStatusRuntime)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.AbortHttpStatusRuntime)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.DelayDurationRuntime) > 0 {
		i -= len(m.DelayDurationRuntime)
		copy(dAtA[i:], m.DelayDurationRuntime)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.DelayDurationRuntime)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.AbortPercentRuntime) > 0 {
		i -= len(m.AbortPercentRuntime)
		copy(dAtA[i:], m.AbortPercentRuntime)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.AbortPercentRuntime)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.DelayPercentRuntime) > 0 {
		i -= len(m.DelayPercentRuntime)
		copy(dAtA[i:], m.DelayPercentRuntime)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.DelayPercentRuntime)))
		i--
		dAtA[i] = 0x42
	}
	if m.ResponseRateLimit != nil {
		if vtmsg, ok := interface{}(m.ResponseRateLimit).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.ResponseRateLimit)
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
	if m.MaxActiveFaults != nil {
		size, err := (*wrapperspb.UInt32Value)(m.MaxActiveFaults).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DownstreamNodes) > 0 {
		for iNdEx := len(m.DownstreamNodes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DownstreamNodes[iNdEx])
			copy(dAtA[i:], m.DownstreamNodes[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.DownstreamNodes[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
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
			dAtA[i] = 0x22
		}
	}
	if len(m.UpstreamCluster) > 0 {
		i -= len(m.UpstreamCluster)
		copy(dAtA[i:], m.UpstreamCluster)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.UpstreamCluster)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Abort != nil {
		size, err := m.Abort.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.Delay != nil {
		if vtmsg, ok := interface{}(m.Delay).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Delay)
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

func (m *FaultAbort_HeaderAbort) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *FaultAbort) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.ErrorType.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if m.Percentage != nil {
		if size, ok := interface{}(m.Percentage).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Percentage)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *FaultAbort_HttpStatus) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + protohelpers.SizeOfVarint(uint64(m.HttpStatus))
	return n
}
func (m *FaultAbort_HeaderAbort_) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HeaderAbort != nil {
		l = m.HeaderAbort.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *FaultAbort_GrpcStatus) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + protohelpers.SizeOfVarint(uint64(m.GrpcStatus))
	return n
}
func (m *HTTPFault) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Delay != nil {
		if size, ok := interface{}(m.Delay).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Delay)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Abort != nil {
		l = m.Abort.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.UpstreamCluster)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
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
	if len(m.DownstreamNodes) > 0 {
		for _, s := range m.DownstreamNodes {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.MaxActiveFaults != nil {
		l = (*wrapperspb.UInt32Value)(m.MaxActiveFaults).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.ResponseRateLimit != nil {
		if size, ok := interface{}(m.ResponseRateLimit).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.ResponseRateLimit)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.DelayPercentRuntime)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.AbortPercentRuntime)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.DelayDurationRuntime)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.AbortHttpStatusRuntime)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.MaxActiveFaultsRuntime)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.ResponseRateLimitPercentRuntime)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.AbortGrpcStatusRuntime)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.DisableDownstreamClusterStats {
		n += 2
	}
	if m.FilterMetadata != nil {
		l = (*structpb.Struct)(m.FilterMetadata).SizeVT()
		n += 2 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}
