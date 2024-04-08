//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/service/discovery/v3/discovery.proto

package discoveryv3

import (
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	anypb "github.com/planetscale/vtprotobuf/types/known/anypb"
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

func (m *ResourceLocator) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *ResourceLocator) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *ResourceLocator) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.DynamicParameters) > 0 {
		for k := range m.DynamicParameters {
			v := m.DynamicParameters[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = protohelpers.EncodeVarint(dAtA, i, uint64(baseI-i))
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

func (m *ResourceName) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *ResourceName) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *ResourceName) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.DynamicParameterConstraints != nil {
		size, err := m.DynamicParameterConstraints.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
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

func (m *DiscoveryRequest) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DiscoveryRequest) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DiscoveryRequest) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.ResourceLocators) > 0 {
		for iNdEx := len(m.ResourceLocators) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.ResourceLocators[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.ErrorDetail != nil {
		if vtmsg, ok := interface{}(m.ErrorDetail).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.ErrorDetail)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.ResponseNonce) > 0 {
		i -= len(m.ResponseNonce)
		copy(dAtA[i:], m.ResponseNonce)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ResponseNonce)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TypeUrl) > 0 {
		i -= len(m.TypeUrl)
		copy(dAtA[i:], m.TypeUrl)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.TypeUrl)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ResourceNames) > 0 {
		for iNdEx := len(m.ResourceNames) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ResourceNames[iNdEx])
			copy(dAtA[i:], m.ResourceNames[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ResourceNames[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Node != nil {
		if vtmsg, ok := interface{}(m.Node).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Node)
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
	if len(m.VersionInfo) > 0 {
		i -= len(m.VersionInfo)
		copy(dAtA[i:], m.VersionInfo)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.VersionInfo)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DiscoveryResponse) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DiscoveryResponse) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DiscoveryResponse) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.ControlPlane != nil {
		if vtmsg, ok := interface{}(m.ControlPlane).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.ControlPlane)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.Nonce) > 0 {
		i -= len(m.Nonce)
		copy(dAtA[i:], m.Nonce)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Nonce)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TypeUrl) > 0 {
		i -= len(m.TypeUrl)
		copy(dAtA[i:], m.TypeUrl)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.TypeUrl)))
		i--
		dAtA[i] = 0x22
	}
	if m.Canary {
		i--
		if m.Canary {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Resources) > 0 {
		for iNdEx := len(m.Resources) - 1; iNdEx >= 0; iNdEx-- {
			size, err := (*anypb.Any)(m.Resources[iNdEx]).MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.VersionInfo) > 0 {
		i -= len(m.VersionInfo)
		copy(dAtA[i:], m.VersionInfo)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.VersionInfo)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DeltaDiscoveryRequest) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DeltaDiscoveryRequest) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DeltaDiscoveryRequest) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.ResourceLocatorsUnsubscribe) > 0 {
		for iNdEx := len(m.ResourceLocatorsUnsubscribe) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.ResourceLocatorsUnsubscribe[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.ResourceLocatorsSubscribe) > 0 {
		for iNdEx := len(m.ResourceLocatorsSubscribe) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.ResourceLocatorsSubscribe[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x42
		}
	}
	if m.ErrorDetail != nil {
		if vtmsg, ok := interface{}(m.ErrorDetail).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.ErrorDetail)
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
	if len(m.ResponseNonce) > 0 {
		i -= len(m.ResponseNonce)
		copy(dAtA[i:], m.ResponseNonce)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ResponseNonce)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.InitialResourceVersions) > 0 {
		for k := range m.InitialResourceVersions {
			v := m.InitialResourceVersions[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = protohelpers.EncodeVarint(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ResourceNamesUnsubscribe) > 0 {
		for iNdEx := len(m.ResourceNamesUnsubscribe) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ResourceNamesUnsubscribe[iNdEx])
			copy(dAtA[i:], m.ResourceNamesUnsubscribe[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ResourceNamesUnsubscribe[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ResourceNamesSubscribe) > 0 {
		for iNdEx := len(m.ResourceNamesSubscribe) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ResourceNamesSubscribe[iNdEx])
			copy(dAtA[i:], m.ResourceNamesSubscribe[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ResourceNamesSubscribe[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.TypeUrl) > 0 {
		i -= len(m.TypeUrl)
		copy(dAtA[i:], m.TypeUrl)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.TypeUrl)))
		i--
		dAtA[i] = 0x12
	}
	if m.Node != nil {
		if vtmsg, ok := interface{}(m.Node).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Node)
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

func (m *DeltaDiscoveryResponse) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DeltaDiscoveryResponse) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DeltaDiscoveryResponse) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.RemovedResourceNames) > 0 {
		for iNdEx := len(m.RemovedResourceNames) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.RemovedResourceNames[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x42
		}
	}
	if m.ControlPlane != nil {
		if vtmsg, ok := interface{}(m.ControlPlane).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.ControlPlane)
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
	if len(m.RemovedResources) > 0 {
		for iNdEx := len(m.RemovedResources) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RemovedResources[iNdEx])
			copy(dAtA[i:], m.RemovedResources[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.RemovedResources[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Nonce) > 0 {
		i -= len(m.Nonce)
		copy(dAtA[i:], m.Nonce)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Nonce)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TypeUrl) > 0 {
		i -= len(m.TypeUrl)
		copy(dAtA[i:], m.TypeUrl)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.TypeUrl)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Resources) > 0 {
		for iNdEx := len(m.Resources) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Resources[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.SystemVersionInfo) > 0 {
		i -= len(m.SystemVersionInfo)
		copy(dAtA[i:], m.SystemVersionInfo)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.SystemVersionInfo)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DynamicParameterConstraints_SingleConstraint_Exists) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DynamicParameterConstraints_SingleConstraint_Exists) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DynamicParameterConstraints_SingleConstraint_Exists) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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

func (m *DynamicParameterConstraints_SingleConstraint) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DynamicParameterConstraints_SingleConstraint) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DynamicParameterConstraints_SingleConstraint) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if msg, ok := m.ConstraintType.(*DynamicParameterConstraints_SingleConstraint_Exists_); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ConstraintType.(*DynamicParameterConstraints_SingleConstraint_Value); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
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

func (m *DynamicParameterConstraints_SingleConstraint_Value) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DynamicParameterConstraints_SingleConstraint_Value) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Value)
	copy(dAtA[i:], m.Value)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Value)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *DynamicParameterConstraints_SingleConstraint_Exists_) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DynamicParameterConstraints_SingleConstraint_Exists_) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Exists != nil {
		size, err := m.Exists.MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *DynamicParameterConstraints_ConstraintList) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DynamicParameterConstraints_ConstraintList) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DynamicParameterConstraints_ConstraintList) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.Constraints) > 0 {
		for iNdEx := len(m.Constraints) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Constraints[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
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

func (m *DynamicParameterConstraints) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DynamicParameterConstraints) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DynamicParameterConstraints) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if msg, ok := m.Type.(*DynamicParameterConstraints_NotConstraints); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Type.(*DynamicParameterConstraints_AndConstraints); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Type.(*DynamicParameterConstraints_OrConstraints); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Type.(*DynamicParameterConstraints_Constraint); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *DynamicParameterConstraints_Constraint) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DynamicParameterConstraints_Constraint) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Constraint != nil {
		size, err := m.Constraint.MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *DynamicParameterConstraints_OrConstraints) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DynamicParameterConstraints_OrConstraints) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.OrConstraints != nil {
		size, err := m.OrConstraints.MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *DynamicParameterConstraints_AndConstraints) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DynamicParameterConstraints_AndConstraints) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AndConstraints != nil {
		size, err := m.AndConstraints.MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *DynamicParameterConstraints_NotConstraints) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DynamicParameterConstraints_NotConstraints) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NotConstraints != nil {
		size, err := m.NotConstraints.MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *Resource_CacheControl) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *Resource_CacheControl) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Resource_CacheControl) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.DoNotCache {
		i--
		if m.DoNotCache {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Resource) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *Resource) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Resource) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
		if vtmsg, ok := interface{}(m.Metadata).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Metadata)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.ResourceName != nil {
		size, err := m.ResourceName.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x42
	}
	if m.CacheControl != nil {
		size, err := m.CacheControl.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x3a
	}
	if m.Ttl != nil {
		size, err := (*durationpb.Duration)(m.Ttl).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Aliases) > 0 {
		for iNdEx := len(m.Aliases) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Aliases[iNdEx])
			copy(dAtA[i:], m.Aliases[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Aliases[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Resource != nil {
		size, err := (*anypb.Any)(m.Resource).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResourceLocator) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.DynamicParameters) > 0 {
		for k, v := range m.DynamicParameters {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + protohelpers.SizeOfVarint(uint64(len(k))) + 1 + len(v) + protohelpers.SizeOfVarint(uint64(len(v)))
			n += mapEntrySize + 1 + protohelpers.SizeOfVarint(uint64(mapEntrySize))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *ResourceName) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.DynamicParameterConstraints != nil {
		l = m.DynamicParameterConstraints.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *DiscoveryRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.VersionInfo)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Node != nil {
		if size, ok := interface{}(m.Node).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Node)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.ResourceNames) > 0 {
		for _, s := range m.ResourceNames {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	l = len(m.TypeUrl)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.ResponseNonce)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.ErrorDetail != nil {
		if size, ok := interface{}(m.ErrorDetail).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.ErrorDetail)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.ResourceLocators) > 0 {
		for _, e := range m.ResourceLocators {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *DiscoveryResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.VersionInfo)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.Resources) > 0 {
		for _, e := range m.Resources {
			l = (*anypb.Any)(e).SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.Canary {
		n += 2
	}
	l = len(m.TypeUrl)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.ControlPlane != nil {
		if size, ok := interface{}(m.ControlPlane).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.ControlPlane)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *DeltaDiscoveryRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Node != nil {
		if size, ok := interface{}(m.Node).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Node)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.TypeUrl)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.ResourceNamesSubscribe) > 0 {
		for _, s := range m.ResourceNamesSubscribe {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.ResourceNamesUnsubscribe) > 0 {
		for _, s := range m.ResourceNamesUnsubscribe {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.InitialResourceVersions) > 0 {
		for k, v := range m.InitialResourceVersions {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + protohelpers.SizeOfVarint(uint64(len(k))) + 1 + len(v) + protohelpers.SizeOfVarint(uint64(len(v)))
			n += mapEntrySize + 1 + protohelpers.SizeOfVarint(uint64(mapEntrySize))
		}
	}
	l = len(m.ResponseNonce)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.ErrorDetail != nil {
		if size, ok := interface{}(m.ErrorDetail).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.ErrorDetail)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.ResourceLocatorsSubscribe) > 0 {
		for _, e := range m.ResourceLocatorsSubscribe {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.ResourceLocatorsUnsubscribe) > 0 {
		for _, e := range m.ResourceLocatorsUnsubscribe {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *DeltaDiscoveryResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SystemVersionInfo)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.Resources) > 0 {
		for _, e := range m.Resources {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	l = len(m.TypeUrl)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.RemovedResources) > 0 {
		for _, s := range m.RemovedResources {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.ControlPlane != nil {
		if size, ok := interface{}(m.ControlPlane).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.ControlPlane)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.RemovedResourceNames) > 0 {
		for _, e := range m.RemovedResourceNames {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *DynamicParameterConstraints_SingleConstraint_Exists) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *DynamicParameterConstraints_SingleConstraint) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if vtmsg, ok := m.ConstraintType.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *DynamicParameterConstraints_SingleConstraint_Value) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
func (m *DynamicParameterConstraints_SingleConstraint_Exists_) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Exists != nil {
		l = m.Exists.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *DynamicParameterConstraints_ConstraintList) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Constraints) > 0 {
		for _, e := range m.Constraints {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *DynamicParameterConstraints) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Type.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *DynamicParameterConstraints_Constraint) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Constraint != nil {
		l = m.Constraint.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *DynamicParameterConstraints_OrConstraints) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrConstraints != nil {
		l = m.OrConstraints.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *DynamicParameterConstraints_AndConstraints) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AndConstraints != nil {
		l = m.AndConstraints.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *DynamicParameterConstraints_NotConstraints) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NotConstraints != nil {
		l = m.NotConstraints.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *Resource_CacheControl) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DoNotCache {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *Resource) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Resource != nil {
		l = (*anypb.Any)(m.Resource).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.Aliases) > 0 {
		for _, s := range m.Aliases {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.Ttl != nil {
		l = (*durationpb.Duration)(m.Ttl).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.CacheControl != nil {
		l = m.CacheControl.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.ResourceName != nil {
		l = m.ResourceName.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Metadata != nil {
		if size, ok := interface{}(m.Metadata).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Metadata)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}
