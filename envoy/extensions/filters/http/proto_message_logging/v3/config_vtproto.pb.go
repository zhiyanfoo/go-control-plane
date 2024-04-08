//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/extensions/filters/http/proto_message_logging/v3/config.proto

package proto_message_loggingv3

import (
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *ProtoMessageLoggingConfig) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *ProtoMessageLoggingConfig) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *ProtoMessageLoggingConfig) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.LoggingByMethod) > 0 {
		for k := range m.LoggingByMethod {
			v := m.LoggingByMethod[k]
			baseI := i
			size, err := v.MarshalToSizedBufferVTStrict(dAtA[:i])
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
			dAtA[i] = 0x22
		}
	}
	if m.Mode != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Mode))
		i--
		dAtA[i] = 0x18
	}
	if msg, ok := m.DescriptorSet.(*ProtoMessageLoggingConfig_ProtoDescriptorTypedMetadata); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.DescriptorSet.(*ProtoMessageLoggingConfig_DataSource); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *ProtoMessageLoggingConfig_DataSource) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *ProtoMessageLoggingConfig_DataSource) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DataSource != nil {
		if vtmsg, ok := interface{}(m.DataSource).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.DataSource)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0xa
	} else {
		i = protohelpers.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *ProtoMessageLoggingConfig_ProtoDescriptorTypedMetadata) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *ProtoMessageLoggingConfig_ProtoDescriptorTypedMetadata) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.ProtoDescriptorTypedMetadata)
	copy(dAtA[i:], m.ProtoDescriptorTypedMetadata)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ProtoDescriptorTypedMetadata)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *MethodLogging) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *MethodLogging) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *MethodLogging) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.ResponseLoggingByField) > 0 {
		for k := range m.ResponseLoggingByField {
			v := m.ResponseLoggingByField[k]
			baseI := i
			i = protohelpers.EncodeVarint(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = protohelpers.EncodeVarint(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RequestLoggingByField) > 0 {
		for k := range m.RequestLoggingByField {
			v := m.RequestLoggingByField[k]
			baseI := i
			i = protohelpers.EncodeVarint(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
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
	return len(dAtA) - i, nil
}

func (m *ProtoMessageLoggingConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.DescriptorSet.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if m.Mode != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Mode))
	}
	if len(m.LoggingByMethod) > 0 {
		for k, v := range m.LoggingByMethod {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.SizeVT()
			}
			l += 1 + protohelpers.SizeOfVarint(uint64(l))
			mapEntrySize := 1 + len(k) + protohelpers.SizeOfVarint(uint64(len(k))) + l
			n += mapEntrySize + 1 + protohelpers.SizeOfVarint(uint64(mapEntrySize))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *ProtoMessageLoggingConfig_DataSource) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DataSource != nil {
		if size, ok := interface{}(m.DataSource).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.DataSource)
		}
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	} else {
		n += 3
	}
	return n
}
func (m *ProtoMessageLoggingConfig_ProtoDescriptorTypedMetadata) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProtoDescriptorTypedMetadata)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
func (m *MethodLogging) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RequestLoggingByField) > 0 {
		for k, v := range m.RequestLoggingByField {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + protohelpers.SizeOfVarint(uint64(len(k))) + 1 + protohelpers.SizeOfVarint(uint64(v))
			n += mapEntrySize + 1 + protohelpers.SizeOfVarint(uint64(mapEntrySize))
		}
	}
	if len(m.ResponseLoggingByField) > 0 {
		for k, v := range m.ResponseLoggingByField {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + protohelpers.SizeOfVarint(uint64(len(k))) + 1 + protohelpers.SizeOfVarint(uint64(v))
			n += mapEntrySize + 1 + protohelpers.SizeOfVarint(uint64(mapEntrySize))
		}
	}
	n += len(m.unknownFields)
	return n
}
