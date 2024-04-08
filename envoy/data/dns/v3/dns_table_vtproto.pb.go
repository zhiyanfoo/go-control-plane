//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/data/dns/v3/dns_table.proto

package dnsv3

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

func (m *DnsTable_AddressList) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DnsTable_AddressList) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DnsTable_AddressList) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.Address) > 0 {
		for iNdEx := len(m.Address) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Address[iNdEx])
			copy(dAtA[i:], m.Address[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Address[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DnsTable_DnsServiceProtocol) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DnsTable_DnsServiceProtocol) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DnsTable_DnsServiceProtocol) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if msg, ok := m.ProtocolConfig.(*DnsTable_DnsServiceProtocol_Name); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ProtocolConfig.(*DnsTable_DnsServiceProtocol_Number); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *DnsTable_DnsServiceProtocol_Number) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DnsTable_DnsServiceProtocol_Number) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Number))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}
func (m *DnsTable_DnsServiceProtocol_Name) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DnsTable_DnsServiceProtocol_Name) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Name)
	copy(dAtA[i:], m.Name)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Name)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *DnsTable_DnsServiceTarget) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DnsTable_DnsServiceTarget) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DnsTable_DnsServiceTarget) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.Port != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x28
	}
	if m.Weight != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Weight))
		i--
		dAtA[i] = 0x20
	}
	if m.Priority != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Priority))
		i--
		dAtA[i] = 0x18
	}
	if msg, ok := m.EndpointType.(*DnsTable_DnsServiceTarget_ClusterName); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.EndpointType.(*DnsTable_DnsServiceTarget_HostName); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *DnsTable_DnsServiceTarget_HostName) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DnsTable_DnsServiceTarget_HostName) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.HostName)
	copy(dAtA[i:], m.HostName)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.HostName)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *DnsTable_DnsServiceTarget_ClusterName) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DnsTable_DnsServiceTarget_ClusterName) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.ClusterName)
	copy(dAtA[i:], m.ClusterName)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ClusterName)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *DnsTable_DnsService) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DnsTable_DnsService) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DnsTable_DnsService) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.Targets) > 0 {
		for iNdEx := len(m.Targets) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Targets[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Ttl != nil {
		size, err := (*durationpb.Duration)(m.Ttl).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.Protocol != nil {
		size, err := m.Protocol.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ServiceName) > 0 {
		i -= len(m.ServiceName)
		copy(dAtA[i:], m.ServiceName)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ServiceName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DnsTable_DnsServiceList) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DnsTable_DnsServiceList) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DnsTable_DnsServiceList) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.Services) > 0 {
		for iNdEx := len(m.Services) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Services[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
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

func (m *DnsTable_DnsEndpoint) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DnsTable_DnsEndpoint) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DnsTable_DnsEndpoint) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if msg, ok := m.EndpointConfig.(*DnsTable_DnsEndpoint_ServiceList); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.EndpointConfig.(*DnsTable_DnsEndpoint_ClusterName); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.EndpointConfig.(*DnsTable_DnsEndpoint_AddressList); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *DnsTable_DnsEndpoint_AddressList) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DnsTable_DnsEndpoint_AddressList) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AddressList != nil {
		size, err := m.AddressList.MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *DnsTable_DnsEndpoint_ClusterName) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DnsTable_DnsEndpoint_ClusterName) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.ClusterName)
	copy(dAtA[i:], m.ClusterName)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ClusterName)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *DnsTable_DnsEndpoint_ServiceList) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DnsTable_DnsEndpoint_ServiceList) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ServiceList != nil {
		size, err := m.ServiceList.MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *DnsTable_DnsVirtualDomain) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DnsTable_DnsVirtualDomain) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DnsTable_DnsVirtualDomain) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.AnswerTtl != nil {
		size, err := (*durationpb.Duration)(m.AnswerTtl).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.Endpoint != nil {
		size, err := m.Endpoint.MarshalToSizedBufferVTStrict(dAtA[:i])
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

func (m *DnsTable) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *DnsTable) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *DnsTable) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.KnownSuffixes) > 0 {
		for iNdEx := len(m.KnownSuffixes) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.KnownSuffixes[iNdEx]).(interface {
				MarshalToSizedBufferVTStrict([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.KnownSuffixes[iNdEx])
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
	if len(m.VirtualDomains) > 0 {
		for iNdEx := len(m.VirtualDomains) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.VirtualDomains[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.ExternalRetryCount != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.ExternalRetryCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DnsTable_AddressList) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Address) > 0 {
		for _, s := range m.Address {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *DnsTable_DnsServiceProtocol) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.ProtocolConfig.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *DnsTable_DnsServiceProtocol_Number) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + protohelpers.SizeOfVarint(uint64(m.Number))
	return n
}
func (m *DnsTable_DnsServiceProtocol_Name) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
func (m *DnsTable_DnsServiceTarget) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.EndpointType.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if m.Priority != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Priority))
	}
	if m.Weight != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Weight))
	}
	if m.Port != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Port))
	}
	n += len(m.unknownFields)
	return n
}

func (m *DnsTable_DnsServiceTarget_HostName) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HostName)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
func (m *DnsTable_DnsServiceTarget_ClusterName) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClusterName)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
func (m *DnsTable_DnsService) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Protocol != nil {
		l = m.Protocol.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Ttl != nil {
		l = (*durationpb.Duration)(m.Ttl).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.Targets) > 0 {
		for _, e := range m.Targets {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *DnsTable_DnsServiceList) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Services) > 0 {
		for _, e := range m.Services {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *DnsTable_DnsEndpoint) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.EndpointConfig.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *DnsTable_DnsEndpoint_AddressList) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AddressList != nil {
		l = m.AddressList.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *DnsTable_DnsEndpoint_ClusterName) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClusterName)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
func (m *DnsTable_DnsEndpoint_ServiceList) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ServiceList != nil {
		l = m.ServiceList.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *DnsTable_DnsVirtualDomain) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Endpoint != nil {
		l = m.Endpoint.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.AnswerTtl != nil {
		l = (*durationpb.Duration)(m.AnswerTtl).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *DnsTable) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExternalRetryCount != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.ExternalRetryCount))
	}
	if len(m.VirtualDomains) > 0 {
		for _, e := range m.VirtualDomains {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.KnownSuffixes) > 0 {
		for _, e := range m.KnownSuffixes {
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
