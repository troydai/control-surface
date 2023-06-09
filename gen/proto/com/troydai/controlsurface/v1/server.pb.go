// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: proto/com/troydai/controlsurface/v1/server.proto

package controlsurfacev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ParameterValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Value:
	//
	//	*ParameterValue_FloatValue
	//	*ParameterValue_IntValue
	//	*ParameterValue_BoolValue
	//	*ParameterValue_StringValue
	Value isParameterValue_Value `protobuf_oneof:"value"`
}

func (x *ParameterValue) Reset() {
	*x = ParameterValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_com_troydai_controlsurface_v1_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParameterValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParameterValue) ProtoMessage() {}

func (x *ParameterValue) ProtoReflect() protoreflect.Message {
	mi := &file_proto_com_troydai_controlsurface_v1_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParameterValue.ProtoReflect.Descriptor instead.
func (*ParameterValue) Descriptor() ([]byte, []int) {
	return file_proto_com_troydai_controlsurface_v1_server_proto_rawDescGZIP(), []int{0}
}

func (x *ParameterValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *ParameterValue) GetValue() isParameterValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *ParameterValue) GetFloatValue() float32 {
	if x, ok := x.GetValue().(*ParameterValue_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (x *ParameterValue) GetIntValue() int32 {
	if x, ok := x.GetValue().(*ParameterValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *ParameterValue) GetBoolValue() bool {
	if x, ok := x.GetValue().(*ParameterValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *ParameterValue) GetStringValue() string {
	if x, ok := x.GetValue().(*ParameterValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

type isParameterValue_Value interface {
	isParameterValue_Value()
}

type ParameterValue_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,2,opt,name=float_value,json=floatValue,proto3,oneof"`
}

type ParameterValue_IntValue struct {
	IntValue int32 `protobuf:"varint,3,opt,name=int_value,json=intValue,proto3,oneof"`
}

type ParameterValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type ParameterValue_StringValue struct {
	StringValue string `protobuf:"bytes,5,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*ParameterValue_FloatValue) isParameterValue_Value() {}

func (*ParameterValue_IntValue) isParameterValue_Value() {}

func (*ParameterValue_BoolValue) isParameterValue_Value() {}

func (*ParameterValue_StringValue) isParameterValue_Value() {}

type GetParametersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId       string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientNs       string   `protobuf:"bytes,2,opt,name=client_ns,json=clientNs,proto3" json:"client_ns,omitempty"`
	ParameterNames []string `protobuf:"bytes,3,rep,name=parameter_names,json=parameterNames,proto3" json:"parameter_names,omitempty"`
}

func (x *GetParametersRequest) Reset() {
	*x = GetParametersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_com_troydai_controlsurface_v1_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetParametersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParametersRequest) ProtoMessage() {}

func (x *GetParametersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_com_troydai_controlsurface_v1_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParametersRequest.ProtoReflect.Descriptor instead.
func (*GetParametersRequest) Descriptor() ([]byte, []int) {
	return file_proto_com_troydai_controlsurface_v1_server_proto_rawDescGZIP(), []int{1}
}

func (x *GetParametersRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *GetParametersRequest) GetClientNs() string {
	if x != nil {
		return x.ClientNs
	}
	return ""
}

func (x *GetParametersRequest) GetParameterNames() []string {
	if x != nil {
		return x.ParameterNames
	}
	return nil
}

type GetParametersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParameterValues []*ParameterValue `protobuf:"bytes,1,rep,name=parameter_values,json=parameterValues,proto3" json:"parameter_values,omitempty"`
}

func (x *GetParametersResponse) Reset() {
	*x = GetParametersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_com_troydai_controlsurface_v1_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetParametersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParametersResponse) ProtoMessage() {}

func (x *GetParametersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_com_troydai_controlsurface_v1_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParametersResponse.ProtoReflect.Descriptor instead.
func (*GetParametersResponse) Descriptor() ([]byte, []int) {
	return file_proto_com_troydai_controlsurface_v1_server_proto_rawDescGZIP(), []int{2}
}

func (x *GetParametersResponse) GetParameterValues() []*ParameterValue {
	if x != nil {
		return x.ParameterValues
	}
	return nil
}

var File_proto_com_troydai_controlsurface_v1_server_proto protoreflect.FileDescriptor

var file_proto_com_troydai_controlsurface_v1_server_proto_rawDesc = []byte{
	0x0a, 0x30, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x6f, 0x79,
	0x64, 0x61, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x75, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x72, 0x6f, 0x79, 0x64, 0x61, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x22, 0xb5, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52,
	0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69,
	0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f,
	0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x79, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x22, 0x71, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a,
	0x10, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x72,
	0x6f, 0x79, 0x64, 0x61, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x75, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x32, 0x8e, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x33, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x72, 0x6f, 0x79, 0x64, 0x61, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x72, 0x6f, 0x79, 0x64, 0x61, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x9e, 0x02, 0x0a, 0x21, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x72, 0x6f, 0x79, 0x64, 0x61, 0x69, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0b,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x55, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2d, 0x74, 0x6f, 0x75, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x6f, 0x79, 0x64, 0x61,
	0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x75, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x54, 0x43, 0xaa, 0x02, 0x1d, 0x43, 0x6f, 0x6d,
	0x2e, 0x54, 0x72, 0x6f, 0x79, 0x64, 0x61, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1d, 0x43, 0x6f, 0x6d,
	0x5c, 0x54, 0x72, 0x6f, 0x79, 0x64, 0x61, 0x69, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x29, 0x43, 0x6f, 0x6d,
	0x5c, 0x54, 0x72, 0x6f, 0x79, 0x64, 0x61, 0x69, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x43, 0x6f, 0x6d, 0x3a, 0x3a, 0x54, 0x72,
	0x6f, 0x79, 0x64, 0x61, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x75,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_com_troydai_controlsurface_v1_server_proto_rawDescOnce sync.Once
	file_proto_com_troydai_controlsurface_v1_server_proto_rawDescData = file_proto_com_troydai_controlsurface_v1_server_proto_rawDesc
)

func file_proto_com_troydai_controlsurface_v1_server_proto_rawDescGZIP() []byte {
	file_proto_com_troydai_controlsurface_v1_server_proto_rawDescOnce.Do(func() {
		file_proto_com_troydai_controlsurface_v1_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_com_troydai_controlsurface_v1_server_proto_rawDescData)
	})
	return file_proto_com_troydai_controlsurface_v1_server_proto_rawDescData
}

var file_proto_com_troydai_controlsurface_v1_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_com_troydai_controlsurface_v1_server_proto_goTypes = []interface{}{
	(*ParameterValue)(nil),        // 0: com.troydai.controlsurface.v1.ParameterValue
	(*GetParametersRequest)(nil),  // 1: com.troydai.controlsurface.v1.GetParametersRequest
	(*GetParametersResponse)(nil), // 2: com.troydai.controlsurface.v1.GetParametersResponse
}
var file_proto_com_troydai_controlsurface_v1_server_proto_depIdxs = []int32{
	0, // 0: com.troydai.controlsurface.v1.GetParametersResponse.parameter_values:type_name -> com.troydai.controlsurface.v1.ParameterValue
	1, // 1: com.troydai.controlsurface.v1.ControlService.GetParameters:input_type -> com.troydai.controlsurface.v1.GetParametersRequest
	2, // 2: com.troydai.controlsurface.v1.ControlService.GetParameters:output_type -> com.troydai.controlsurface.v1.GetParametersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_com_troydai_controlsurface_v1_server_proto_init() }
func file_proto_com_troydai_controlsurface_v1_server_proto_init() {
	if File_proto_com_troydai_controlsurface_v1_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_com_troydai_controlsurface_v1_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParameterValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_com_troydai_controlsurface_v1_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetParametersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_com_troydai_controlsurface_v1_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetParametersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_com_troydai_controlsurface_v1_server_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ParameterValue_FloatValue)(nil),
		(*ParameterValue_IntValue)(nil),
		(*ParameterValue_BoolValue)(nil),
		(*ParameterValue_StringValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_com_troydai_controlsurface_v1_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_com_troydai_controlsurface_v1_server_proto_goTypes,
		DependencyIndexes: file_proto_com_troydai_controlsurface_v1_server_proto_depIdxs,
		MessageInfos:      file_proto_com_troydai_controlsurface_v1_server_proto_msgTypes,
	}.Build()
	File_proto_com_troydai_controlsurface_v1_server_proto = out.File
	file_proto_com_troydai_controlsurface_v1_server_proto_rawDesc = nil
	file_proto_com_troydai_controlsurface_v1_server_proto_goTypes = nil
	file_proto_com_troydai_controlsurface_v1_server_proto_depIdxs = nil
}
