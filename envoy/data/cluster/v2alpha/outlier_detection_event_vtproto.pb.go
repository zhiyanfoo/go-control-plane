//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/data/cluster/v2alpha/outlier_detection_event.proto

package v2alpha

import (
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	timestamppb "github.com/planetscale/vtprotobuf/types/known/timestamppb"
	wrapperspb "github.com/planetscale/vtprotobuf/types/known/wrapperspb"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *OutlierDetectionEvent) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *OutlierDetectionEvent) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OutlierDetectionEvent) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if msg, ok := m.Event.(*OutlierDetectionEvent_EjectFailurePercentageEvent); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Event.(*OutlierDetectionEvent_EjectConsecutiveEvent); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Event.(*OutlierDetectionEvent_EjectSuccessRateEvent); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if m.Enforced {
		i--
		if m.Enforced {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.NumEjections != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.NumEjections))
		i--
		dAtA[i] = 0x38
	}
	if m.Action != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Action))
		i--
		dAtA[i] = 0x30
	}
	if len(m.UpstreamUrl) > 0 {
		i -= len(m.UpstreamUrl)
		copy(dAtA[i:], m.UpstreamUrl)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.UpstreamUrl)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ClusterName) > 0 {
		i -= len(m.ClusterName)
		copy(dAtA[i:], m.ClusterName)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ClusterName)))
		i--
		dAtA[i] = 0x22
	}
	if m.SecsSinceLastAction != nil {
		size, err := (*wrapperspb.UInt64Value)(m.SecsSinceLastAction).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.Timestamp != nil {
		size, err := (*timestamppb.Timestamp)(m.Timestamp).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OutlierDetectionEvent_EjectSuccessRateEvent) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OutlierDetectionEvent_EjectSuccessRateEvent) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EjectSuccessRateEvent != nil {
		size, err := m.EjectSuccessRateEvent.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x4a
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func (m *OutlierDetectionEvent_EjectConsecutiveEvent) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OutlierDetectionEvent_EjectConsecutiveEvent) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EjectConsecutiveEvent != nil {
		size, err := m.EjectConsecutiveEvent.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x52
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x52
	}
	return len(dAtA) - i, nil
}
func (m *OutlierDetectionEvent_EjectFailurePercentageEvent) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OutlierDetectionEvent_EjectFailurePercentageEvent) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EjectFailurePercentageEvent != nil {
		size, err := m.EjectFailurePercentageEvent.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x5a
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x5a
	}
	return len(dAtA) - i, nil
}
func (m *OutlierEjectSuccessRate) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *OutlierEjectSuccessRate) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OutlierEjectSuccessRate) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.ClusterSuccessRateEjectionThreshold != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.ClusterSuccessRateEjectionThreshold))
		i--
		dAtA[i] = 0x18
	}
	if m.ClusterAverageSuccessRate != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.ClusterAverageSuccessRate))
		i--
		dAtA[i] = 0x10
	}
	if m.HostSuccessRate != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.HostSuccessRate))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OutlierEjectConsecutive) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *OutlierEjectConsecutive) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OutlierEjectConsecutive) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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

func (m *OutlierEjectFailurePercentage) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *OutlierEjectFailurePercentage) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OutlierEjectFailurePercentage) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.HostSuccessRate != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.HostSuccessRate))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OutlierDetectionEvent) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Type))
	}
	if m.Timestamp != nil {
		l = (*timestamppb.Timestamp)(m.Timestamp).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.SecsSinceLastAction != nil {
		l = (*wrapperspb.UInt64Value)(m.SecsSinceLastAction).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.ClusterName)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.UpstreamUrl)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Action != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Action))
	}
	if m.NumEjections != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.NumEjections))
	}
	if m.Enforced {
		n += 2
	}
	if vtmsg, ok := m.Event.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *OutlierDetectionEvent_EjectSuccessRateEvent) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EjectSuccessRateEvent != nil {
		l = m.EjectSuccessRateEvent.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *OutlierDetectionEvent_EjectConsecutiveEvent) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EjectConsecutiveEvent != nil {
		l = m.EjectConsecutiveEvent.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *OutlierDetectionEvent_EjectFailurePercentageEvent) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EjectFailurePercentageEvent != nil {
		l = m.EjectFailurePercentageEvent.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *OutlierEjectSuccessRate) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HostSuccessRate != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.HostSuccessRate))
	}
	if m.ClusterAverageSuccessRate != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.ClusterAverageSuccessRate))
	}
	if m.ClusterSuccessRateEjectionThreshold != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.ClusterSuccessRateEjectionThreshold))
	}
	n += len(m.unknownFields)
	return n
}

func (m *OutlierEjectConsecutive) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *OutlierEjectFailurePercentage) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HostSuccessRate != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.HostSuccessRate))
	}
	n += len(m.unknownFields)
	return n
}
