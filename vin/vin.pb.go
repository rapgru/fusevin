// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.3
// source: vin/vin.proto

package vin

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

type CreatePuppetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VinFilename string `protobuf:"bytes,2,opt,name=vin_filename,json=vinFilename,proto3" json:"vin_filename,omitempty"`
}

func (x *CreatePuppetReply) Reset() {
	*x = CreatePuppetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vin_vin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePuppetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePuppetReply) ProtoMessage() {}

func (x *CreatePuppetReply) ProtoReflect() protoreflect.Message {
	mi := &file_vin_vin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePuppetReply.ProtoReflect.Descriptor instead.
func (*CreatePuppetReply) Descriptor() ([]byte, []int) {
	return file_vin_vin_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePuppetReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreatePuppetReply) GetVinFilename() string {
	if x != nil {
		return x.VinFilename
	}
	return ""
}

type StartStdinNotifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StartStdinNotifyRequest) Reset() {
	*x = StartStdinNotifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vin_vin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartStdinNotifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartStdinNotifyRequest) ProtoMessage() {}

func (x *StartStdinNotifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vin_vin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartStdinNotifyRequest.ProtoReflect.Descriptor instead.
func (*StartStdinNotifyRequest) Descriptor() ([]byte, []int) {
	return file_vin_vin_proto_rawDescGZIP(), []int{1}
}

func (x *StartStdinNotifyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StdinNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StdinNotify) Reset() {
	*x = StdinNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vin_vin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StdinNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StdinNotify) ProtoMessage() {}

func (x *StdinNotify) ProtoReflect() protoreflect.Message {
	mi := &file_vin_vin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StdinNotify.ProtoReflect.Descriptor instead.
func (*StdinNotify) Descriptor() ([]byte, []int) {
	return file_vin_vin_proto_rawDescGZIP(), []int{2}
}

type StdinContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *StdinContent) Reset() {
	*x = StdinContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vin_vin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StdinContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StdinContent) ProtoMessage() {}

func (x *StdinContent) ProtoReflect() protoreflect.Message {
	mi := &file_vin_vin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StdinContent.ProtoReflect.Descriptor instead.
func (*StdinContent) Descriptor() ([]byte, []int) {
	return file_vin_vin_proto_rawDescGZIP(), []int{3}
}

func (x *StdinContent) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vin_vin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_vin_vin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_vin_vin_proto_rawDescGZIP(), []int{4}
}

var File_vin_vin_proto protoreflect.FileDescriptor

var file_vin_vin_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x69, 0x6e, 0x2f, 0x76, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x76, 0x69, 0x6e, 0x22, 0x46, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75,
	0x70, 0x70, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x69, 0x6e,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x76, 0x69, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x17,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x64, 0x69, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x74, 0x64, 0x69, 0x6e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x22, 0x28, 0x0a, 0x0c, 0x53, 0x74, 0x64, 0x69, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xb8, 0x01, 0x0a, 0x07, 0x46, 0x75,
	0x73, 0x65, 0x56, 0x69, 0x6e, 0x12, 0x32, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x75, 0x70, 0x70, 0x65, 0x74, 0x12, 0x0a, 0x2e, 0x76, 0x69, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x76, 0x69, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75,
	0x70, 0x70, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x10, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x53, 0x74, 0x64, 0x69, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1c, 0x2e,
	0x76, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x64, 0x69, 0x6e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x76, 0x69,
	0x6e, 0x2e, 0x53, 0x74, 0x64, 0x69, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x30, 0x01, 0x12,
	0x33, 0x0a, 0x12, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x64, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x76, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x64, 0x69,
	0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x0a, 0x2e, 0x76, 0x69, 0x6e, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x1a, 0x5a, 0x18, 0x74, 0x75, 0x77, 0x69, 0x65, 0x6e, 0x2e, 0x61,
	0x63, 0x2e, 0x61, 0x74, 0x2f, 0x66, 0x75, 0x73, 0x65, 0x76, 0x69, 0x6e, 0x2f, 0x76, 0x69, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vin_vin_proto_rawDescOnce sync.Once
	file_vin_vin_proto_rawDescData = file_vin_vin_proto_rawDesc
)

func file_vin_vin_proto_rawDescGZIP() []byte {
	file_vin_vin_proto_rawDescOnce.Do(func() {
		file_vin_vin_proto_rawDescData = protoimpl.X.CompressGZIP(file_vin_vin_proto_rawDescData)
	})
	return file_vin_vin_proto_rawDescData
}

var file_vin_vin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vin_vin_proto_goTypes = []interface{}{
	(*CreatePuppetReply)(nil),       // 0: vin.CreatePuppetReply
	(*StartStdinNotifyRequest)(nil), // 1: vin.StartStdinNotifyRequest
	(*StdinNotify)(nil),             // 2: vin.StdinNotify
	(*StdinContent)(nil),            // 3: vin.StdinContent
	(*Empty)(nil),                   // 4: vin.Empty
}
var file_vin_vin_proto_depIdxs = []int32{
	4, // 0: vin.FuseVin.CreatePuppet:input_type -> vin.Empty
	1, // 1: vin.FuseVin.StartStdinNotify:input_type -> vin.StartStdinNotifyRequest
	3, // 2: vin.FuseVin.SupplyStdinContent:input_type -> vin.StdinContent
	0, // 3: vin.FuseVin.CreatePuppet:output_type -> vin.CreatePuppetReply
	2, // 4: vin.FuseVin.StartStdinNotify:output_type -> vin.StdinNotify
	4, // 5: vin.FuseVin.SupplyStdinContent:output_type -> vin.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vin_vin_proto_init() }
func file_vin_vin_proto_init() {
	if File_vin_vin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vin_vin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePuppetReply); i {
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
		file_vin_vin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartStdinNotifyRequest); i {
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
		file_vin_vin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StdinNotify); i {
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
		file_vin_vin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StdinContent); i {
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
		file_vin_vin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_vin_vin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vin_vin_proto_goTypes,
		DependencyIndexes: file_vin_vin_proto_depIdxs,
		MessageInfos:      file_vin_vin_proto_msgTypes,
	}.Build()
	File_vin_vin_proto = out.File
	file_vin_vin_proto_rawDesc = nil
	file_vin_vin_proto_goTypes = nil
	file_vin_vin_proto_depIdxs = nil
}