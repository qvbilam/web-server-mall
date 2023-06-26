// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.19.4
// source: page.proto

package pageV1

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

type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int64 `protobuf:"varint,2,opt,name=perPage,proto3" json:"perPage,omitempty"`
}

func (x *PageRequest) Reset() {
	*x = PageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequest) ProtoMessage() {}

func (x *PageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequest.ProtoReflect.Descriptor instead.
func (*PageRequest) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{0}
}

func (x *PageRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageRequest) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

var File_page_proto protoreflect.FileDescriptor

var file_page_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x22, 0x3b, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65,
	0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72,
	0x50, 0x61, 0x67, 0x65, 0x42, 0x21, 0x5a, 0x1f, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x71, 0x76, 0x62, 0x69, 0x6c, 0x61, 0x6d, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x70, 0x61, 0x67, 0x65, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_page_proto_rawDescOnce sync.Once
	file_page_proto_rawDescData = file_page_proto_rawDesc
)

func file_page_proto_rawDescGZIP() []byte {
	file_page_proto_rawDescOnce.Do(func() {
		file_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_page_proto_rawDescData)
	})
	return file_page_proto_rawDescData
}

var file_page_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_page_proto_goTypes = []interface{}{
	(*PageRequest)(nil), // 0: pagePb.v1.PageRequest
}
var file_page_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_page_proto_init() }
func file_page_proto_init() {
	if File_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRequest); i {
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
			RawDescriptor: file_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_page_proto_goTypes,
		DependencyIndexes: file_page_proto_depIdxs,
		MessageInfos:      file_page_proto_msgTypes,
	}.Build()
	File_page_proto = out.File
	file_page_proto_rawDesc = nil
	file_page_proto_goTypes = nil
	file_page_proto_depIdxs = nil
}
