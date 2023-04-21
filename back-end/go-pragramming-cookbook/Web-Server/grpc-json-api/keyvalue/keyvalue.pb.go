// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: keyvalue.proto

package keyvalue

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

type SetKeyValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetKeyValueRequest) Reset() {
	*x = SetKeyValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keyvalue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetKeyValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetKeyValueRequest) ProtoMessage() {}

func (x *SetKeyValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keyvalue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetKeyValueRequest.ProtoReflect.Descriptor instead.
func (*SetKeyValueRequest) Descriptor() ([]byte, []int) {
	return file_keyvalue_proto_rawDescGZIP(), []int{0}
}

func (x *SetKeyValueRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetKeyValueRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetKeyValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetKeyValueRequest) Reset() {
	*x = GetKeyValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keyvalue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyValueRequest) ProtoMessage() {}

func (x *GetKeyValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keyvalue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyValueRequest.ProtoReflect.Descriptor instead.
func (*GetKeyValueRequest) Descriptor() ([]byte, []int) {
	return file_keyvalue_proto_rawDescGZIP(), []int{1}
}

func (x *GetKeyValueRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type KeyValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValueResponse) Reset() {
	*x = KeyValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keyvalue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValueResponse) ProtoMessage() {}

func (x *KeyValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keyvalue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValueResponse.ProtoReflect.Descriptor instead.
func (*KeyValueResponse) Descriptor() ([]byte, []int) {
	return file_keyvalue_proto_rawDescGZIP(), []int{2}
}

func (x *KeyValueResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValueResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_keyvalue_proto protoreflect.FileDescriptor

var file_keyvalue_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3c, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x26,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x3a, 0x0a, 0x10, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x32, 0x6c, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2f,
	0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x53, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x4b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x4b, 0x65,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keyvalue_proto_rawDescOnce sync.Once
	file_keyvalue_proto_rawDescData = file_keyvalue_proto_rawDesc
)

func file_keyvalue_proto_rawDescGZIP() []byte {
	file_keyvalue_proto_rawDescOnce.Do(func() {
		file_keyvalue_proto_rawDescData = protoimpl.X.CompressGZIP(file_keyvalue_proto_rawDescData)
	})
	return file_keyvalue_proto_rawDescData
}

var file_keyvalue_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_keyvalue_proto_goTypes = []interface{}{
	(*SetKeyValueRequest)(nil), // 0: SetKeyValueRequest
	(*GetKeyValueRequest)(nil), // 1: GetKeyValueRequest
	(*KeyValueResponse)(nil),   // 2: KeyValueResponse
}
var file_keyvalue_proto_depIdxs = []int32{
	0, // 0: KeyValue.Set:input_type -> SetKeyValueRequest
	1, // 1: KeyValue.Get:input_type -> GetKeyValueRequest
	2, // 2: KeyValue.Set:output_type -> KeyValueResponse
	2, // 3: KeyValue.Get:output_type -> KeyValueResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_keyvalue_proto_init() }
func file_keyvalue_proto_init() {
	if File_keyvalue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keyvalue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetKeyValueRequest); i {
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
		file_keyvalue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyValueRequest); i {
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
		file_keyvalue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValueResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_keyvalue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keyvalue_proto_goTypes,
		DependencyIndexes: file_keyvalue_proto_depIdxs,
		MessageInfos:      file_keyvalue_proto_msgTypes,
	}.Build()
	File_keyvalue_proto = out.File
	file_keyvalue_proto_rawDesc = nil
	file_keyvalue_proto_goTypes = nil
	file_keyvalue_proto_depIdxs = nil
}
