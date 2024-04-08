//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/api/v2/core/base.proto

package core

import (
	binary "encoding/binary"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	anypb "github.com/planetscale/vtprotobuf/types/known/anypb"
	structpb "github.com/planetscale/vtprotobuf/types/known/structpb"
	wrapperspb "github.com/planetscale/vtprotobuf/types/known/wrapperspb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	math "math"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *Locality) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *Locality) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Locality) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.SubZone) > 0 {
		i -= len(m.SubZone)
		copy(dAtA[i:], m.SubZone)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.SubZone)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Zone) > 0 {
		i -= len(m.Zone)
		copy(dAtA[i:], m.Zone)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Zone)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Region) > 0 {
		i -= len(m.Region)
		copy(dAtA[i:], m.Region)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Region)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BuildVersion) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *BuildVersion) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *BuildVersion) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.Metadata != nil {
		size, err := (*structpb.Struct)(m.Metadata).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.Version != nil {
		if vtmsg, ok := interface{}(m.Version).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Version)
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

func (m *Extension) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *Extension) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Extension) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.Disabled {
		i--
		if m.Disabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Version != nil {
		size, err := m.Version.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TypeDescriptor) > 0 {
		i -= len(m.TypeDescriptor)
		copy(dAtA[i:], m.TypeDescriptor)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.TypeDescriptor)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Category) > 0 {
		i -= len(m.Category)
		copy(dAtA[i:], m.Category)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Category)))
		i--
		dAtA[i] = 0x12
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

func (m *Node) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *Node) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Node) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.ListeningAddresses) > 0 {
		for iNdEx := len(m.ListeningAddresses) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.ListeningAddresses[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.ClientFeatures) > 0 {
		for iNdEx := len(m.ClientFeatures) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ClientFeatures[iNdEx])
			copy(dAtA[i:], m.ClientFeatures[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ClientFeatures[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.Extensions) > 0 {
		for iNdEx := len(m.Extensions) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Extensions[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x4a
		}
	}
	if msg, ok := m.UserAgentVersionType.(*Node_UserAgentBuildVersion); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.UserAgentVersionType.(*Node_UserAgentVersion); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if len(m.UserAgentName) > 0 {
		i -= len(m.UserAgentName)
		copy(dAtA[i:], m.UserAgentName)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.UserAgentName)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.BuildVersion) > 0 {
		i -= len(m.BuildVersion)
		copy(dAtA[i:], m.BuildVersion)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.BuildVersion)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Locality != nil {
		size, err := m.Locality.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if m.Metadata != nil {
		size, err := (*structpb.Struct)(m.Metadata).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Cluster) > 0 {
		i -= len(m.Cluster)
		copy(dAtA[i:], m.Cluster)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Cluster)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Node_UserAgentVersion) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Node_UserAgentVersion) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.UserAgentVersion)
	copy(dAtA[i:], m.UserAgentVersion)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.UserAgentVersion)))
	i--
	dAtA[i] = 0x3a
	return len(dAtA) - i, nil
}
func (m *Node_UserAgentBuildVersion) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Node_UserAgentBuildVersion) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.UserAgentBuildVersion != nil {
		size, err := m.UserAgentBuildVersion.MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *Metadata) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *Metadata) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.FilterMetadata) > 0 {
		for k := range m.FilterMetadata {
			v := m.FilterMetadata[k]
			baseI := i
			size, err := (*structpb.Struct)(v).MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = protohelpers.EncodeVarint(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RuntimeUInt32) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RuntimeUInt32) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RuntimeUInt32) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.RuntimeKey) > 0 {
		i -= len(m.RuntimeKey)
		copy(dAtA[i:], m.RuntimeKey)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.RuntimeKey)))
		i--
		dAtA[i] = 0x1a
	}
	if m.DefaultValue != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.DefaultValue))
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func (m *RuntimeDouble) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RuntimeDouble) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RuntimeDouble) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.RuntimeKey) > 0 {
		i -= len(m.RuntimeKey)
		copy(dAtA[i:], m.RuntimeKey)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.RuntimeKey)))
		i--
		dAtA[i] = 0x12
	}
	if m.DefaultValue != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.DefaultValue))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func (m *RuntimeFeatureFlag) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RuntimeFeatureFlag) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RuntimeFeatureFlag) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.RuntimeKey) > 0 {
		i -= len(m.RuntimeKey)
		copy(dAtA[i:], m.RuntimeKey)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.RuntimeKey)))
		i--
		dAtA[i] = 0x12
	}
	if m.DefaultValue != nil {
		size, err := (*wrapperspb.BoolValue)(m.DefaultValue).MarshalToSizedBufferVTStrict(dAtA[:i])
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

func (m *HeaderValue) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *HeaderValue) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *HeaderValue) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HeaderValueOption) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *HeaderValueOption) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *HeaderValueOption) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.Append != nil {
		size, err := (*wrapperspb.BoolValue)(m.Append).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		size, err := m.Header.MarshalToSizedBufferVTStrict(dAtA[:i])
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

func (m *HeaderMap) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *HeaderMap) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *HeaderMap) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
			size, err := m.Headers[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DataSource) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DataSource) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DataSource) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if msg, ok := m.Specifier.(*DataSource_InlineString); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Specifier.(*DataSource_InlineBytes); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Specifier.(*DataSource_Filename); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *DataSource_Filename) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DataSource_Filename) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Filename)
	copy(dAtA[i:], m.Filename)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Filename)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *DataSource_InlineBytes) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DataSource_InlineBytes) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.InlineBytes)
	copy(dAtA[i:], m.InlineBytes)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.InlineBytes)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *DataSource_InlineString) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DataSource_InlineString) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.InlineString)
	copy(dAtA[i:], m.InlineString)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.InlineString)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func (m *RetryPolicy) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RetryPolicy) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RetryPolicy) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.NumRetries != nil {
		size, err := (*wrapperspb.UInt32Value)(m.NumRetries).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.RetryBackOff != nil {
		size, err := m.RetryBackOff.MarshalToSizedBufferVTStrict(dAtA[:i])
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

func (m *RemoteDataSource) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RemoteDataSource) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RemoteDataSource) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.RetryPolicy != nil {
		size, err := m.RetryPolicy.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sha256) > 0 {
		i -= len(m.Sha256)
		copy(dAtA[i:], m.Sha256)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Sha256)))
		i--
		dAtA[i] = 0x12
	}
	if m.HttpUri != nil {
		size, err := m.HttpUri.MarshalToSizedBufferVTStrict(dAtA[:i])
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

func (m *AsyncDataSource) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *AsyncDataSource) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *AsyncDataSource) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if msg, ok := m.Specifier.(*AsyncDataSource_Remote); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Specifier.(*AsyncDataSource_Local); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *AsyncDataSource_Local) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *AsyncDataSource_Local) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Local != nil {
		size, err := m.Local.MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *AsyncDataSource_Remote) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *AsyncDataSource_Remote) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Remote != nil {
		size, err := m.Remote.MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *TransportSocket) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *TransportSocket) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *TransportSocket) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if msg, ok := m.ConfigType.(*TransportSocket_TypedConfig); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ConfigType.(*TransportSocket_Config); ok {
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

func (m *TransportSocket_Config) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *TransportSocket_Config) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
func (m *TransportSocket_TypedConfig) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *TransportSocket_TypedConfig) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
func (m *RuntimeFractionalPercent) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RuntimeFractionalPercent) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RuntimeFractionalPercent) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.RuntimeKey) > 0 {
		i -= len(m.RuntimeKey)
		copy(dAtA[i:], m.RuntimeKey)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.RuntimeKey)))
		i--
		dAtA[i] = 0x12
	}
	if m.DefaultValue != nil {
		if vtmsg, ok := interface{}(m.DefaultValue).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.DefaultValue)
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

func (m *ControlPlane) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *ControlPlane) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *ControlPlane) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.Identifier) > 0 {
		i -= len(m.Identifier)
		copy(dAtA[i:], m.Identifier)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Identifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Locality) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Region)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.Zone)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.SubZone)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *BuildVersion) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != nil {
		if size, ok := interface{}(m.Version).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Version)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Metadata != nil {
		l = (*structpb.Struct)(m.Metadata).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *Extension) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.Category)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.TypeDescriptor)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Version != nil {
		l = m.Version.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Disabled {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *Node) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.Cluster)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Metadata != nil {
		l = (*structpb.Struct)(m.Metadata).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Locality != nil {
		l = m.Locality.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.BuildVersion)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.UserAgentName)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if vtmsg, ok := m.UserAgentVersionType.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if len(m.Extensions) > 0 {
		for _, e := range m.Extensions {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.ClientFeatures) > 0 {
		for _, s := range m.ClientFeatures {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.ListeningAddresses) > 0 {
		for _, e := range m.ListeningAddresses {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *Node_UserAgentVersion) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserAgentVersion)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
func (m *Node_UserAgentBuildVersion) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserAgentBuildVersion != nil {
		l = m.UserAgentBuildVersion.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *Metadata) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FilterMetadata) > 0 {
		for k, v := range m.FilterMetadata {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = (*structpb.Struct)(v).SizeVT()
			}
			l += 1 + protohelpers.SizeOfVarint(uint64(l))
			mapEntrySize := 1 + len(k) + protohelpers.SizeOfVarint(uint64(len(k))) + l
			n += mapEntrySize + 1 + protohelpers.SizeOfVarint(uint64(mapEntrySize))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *RuntimeUInt32) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DefaultValue != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.DefaultValue))
	}
	l = len(m.RuntimeKey)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RuntimeDouble) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DefaultValue != 0 {
		n += 9
	}
	l = len(m.RuntimeKey)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RuntimeFeatureFlag) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DefaultValue != nil {
		l = (*wrapperspb.BoolValue)(m.DefaultValue).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.RuntimeKey)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *HeaderValue) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *HeaderValueOption) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Append != nil {
		l = (*wrapperspb.BoolValue)(m.Append).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *HeaderMap) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *DataSource) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Specifier.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *DataSource_Filename) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Filename)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
func (m *DataSource_InlineBytes) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InlineBytes)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
func (m *DataSource_InlineString) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InlineString)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
func (m *RetryPolicy) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RetryBackOff != nil {
		l = m.RetryBackOff.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.NumRetries != nil {
		l = (*wrapperspb.UInt32Value)(m.NumRetries).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RemoteDataSource) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HttpUri != nil {
		l = m.HttpUri.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.Sha256)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.RetryPolicy != nil {
		l = m.RetryPolicy.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *AsyncDataSource) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Specifier.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *AsyncDataSource_Local) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Local != nil {
		l = m.Local.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *AsyncDataSource_Remote) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Remote != nil {
		l = m.Remote.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *TransportSocket) SizeVT() (n int) {
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

func (m *TransportSocket_Config) SizeVT() (n int) {
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
func (m *TransportSocket_TypedConfig) SizeVT() (n int) {
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
func (m *RuntimeFractionalPercent) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DefaultValue != nil {
		if size, ok := interface{}(m.DefaultValue).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.DefaultValue)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.RuntimeKey)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ControlPlane) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}
