//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/config/overload/v2alpha/overload.proto

package v2alpha

import (
	binary "encoding/binary"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	anypb "github.com/planetscale/vtprotobuf/types/known/anypb"
	durationpb "github.com/planetscale/vtprotobuf/types/known/durationpb"
	structpb "github.com/planetscale/vtprotobuf/types/known/structpb"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	math "math"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *ResourceMonitor) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *ResourceMonitor) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *ResourceMonitor) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if msg, ok := m.ConfigType.(*ResourceMonitor_TypedConfig); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ConfigType.(*ResourceMonitor_Config); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResourceMonitor_Config) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *ResourceMonitor_Config) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Config != nil {
		size, err := (*structpb.Struct)(m.Config).MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *ResourceMonitor_TypedConfig) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *ResourceMonitor_TypedConfig) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.TypedConfig != nil {
		size, err := (*anypb.Any)(m.TypedConfig).MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *ThresholdTrigger) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *ThresholdTrigger) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *ThresholdTrigger) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.Value != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func (m *Trigger) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *Trigger) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Trigger) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if msg, ok := m.TriggerOneof.(*Trigger_Threshold); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Trigger_Threshold) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Trigger_Threshold) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Threshold != nil {
		size, err := m.Threshold.MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *OverloadAction) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *OverloadAction) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OverloadAction) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.Triggers) > 0 {
		for iNdEx := len(m.Triggers) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Triggers[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OverloadManager) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *OverloadManager) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OverloadManager) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.ResourceMonitors) > 0 {
		for iNdEx := len(m.ResourceMonitors) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.ResourceMonitors[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.RefreshInterval != nil {
		size, err := (*durationpb.Duration)(m.RefreshInterval).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResourceMonitor) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if vtmsg, ok := m.ConfigType.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *ResourceMonitor_Config) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Config != nil {
		l = (*structpb.Struct)(m.Config).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *ResourceMonitor_TypedConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TypedConfig != nil {
		l = (*anypb.Any)(m.TypedConfig).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *ThresholdTrigger) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 9
	}
	n += len(m.unknownFields)
	return n
}

func (m *Trigger) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if vtmsg, ok := m.TriggerOneof.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *Trigger_Threshold) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Threshold != nil {
		l = m.Threshold.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *OverloadAction) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.Triggers) > 0 {
		for _, e := range m.Triggers {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *OverloadManager) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RefreshInterval != nil {
		l = (*durationpb.Duration)(m.RefreshInterval).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.ResourceMonitors) > 0 {
		for _, e := range m.ResourceMonitors {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.Actions) > 0 {
		for _, e := range m.Actions {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}
