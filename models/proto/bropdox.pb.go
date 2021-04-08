// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: bropdox.proto

package proto

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

type TypeResponse int32

const (
	TypeResponse_CREATED TypeResponse = 0
	TypeResponse_UPDATED TypeResponse = 1
	TypeResponse_DELETED TypeResponse = 2
)

// Enum value maps for TypeResponse.
var (
	TypeResponse_name = map[int32]string{
		0: "CREATED",
		1: "UPDATED",
		2: "DELETED",
	}
	TypeResponse_value = map[string]int32{
		"CREATED": 0,
		"UPDATED": 1,
		"DELETED": 2,
	}
)

func (x TypeResponse) Enum() *TypeResponse {
	p := new(TypeResponse)
	*p = x
	return p
}

func (x TypeResponse) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TypeResponse) Descriptor() protoreflect.EnumDescriptor {
	return file_bropdox_proto_enumTypes[0].Descriptor()
}

func (TypeResponse) Type() protoreflect.EnumType {
	return &file_bropdox_proto_enumTypes[0]
}

func (x TypeResponse) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TypeResponse.Descriptor instead.
func (TypeResponse) EnumDescriptor() ([]byte, []int) {
	return file_bropdox_proto_rawDescGZIP(), []int{0}
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Path      string `protobuf:"bytes,2,opt,name=Path,proto3" json:"Path,omitempty"`
	Size      int64  `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	CreatedAt int64  `protobuf:"varint,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Content   []byte `protobuf:"bytes,6,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bropdox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_bropdox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_bropdox_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *File) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *File) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *File) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type FileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File *File        `protobuf:"bytes,1,opt,name=File,proto3" json:"File,omitempty"`
	Type TypeResponse `protobuf:"varint,2,opt,name=Type,proto3,enum=proto.TypeResponse" json:"Type,omitempty"`
}

func (x *FileResponse) Reset() {
	*x = FileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bropdox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse) ProtoMessage() {}

func (x *FileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bropdox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse.ProtoReflect.Descriptor instead.
func (*FileResponse) Descriptor() ([]byte, []int) {
	return file_bropdox_proto_rawDescGZIP(), []int{1}
}

func (x *FileResponse) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *FileResponse) GetType() TypeResponse {
	if x != nil {
		return x.Type
	}
	return TypeResponse_CREATED
}

var File_bropdox_proto protoreflect.FileDescriptor

var file_bropdox_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x72, 0x6f, 0x70, 0x64, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x58, 0x0a,
	0x0c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x27,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x35, 0x0a, 0x0c, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0xce,
	0x01, 0x0a, 0x07, 0x42, 0x72, 0x6f, 0x70, 0x64, 0x6f, 0x78, 0x12, 0x2e, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0d, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42,
	0x0e, 0x5a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bropdox_proto_rawDescOnce sync.Once
	file_bropdox_proto_rawDescData = file_bropdox_proto_rawDesc
)

func file_bropdox_proto_rawDescGZIP() []byte {
	file_bropdox_proto_rawDescOnce.Do(func() {
		file_bropdox_proto_rawDescData = protoimpl.X.CompressGZIP(file_bropdox_proto_rawDescData)
	})
	return file_bropdox_proto_rawDescData
}

var file_bropdox_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bropdox_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bropdox_proto_goTypes = []interface{}{
	(TypeResponse)(0),    // 0: proto.TypeResponse
	(*File)(nil),         // 1: proto.File
	(*FileResponse)(nil), // 2: proto.FileResponse
}
var file_bropdox_proto_depIdxs = []int32{
	1, // 0: proto.FileResponse.File:type_name -> proto.File
	0, // 1: proto.FileResponse.Type:type_name -> proto.TypeResponse
	1, // 2: proto.Bropdox.CreateFile:input_type -> proto.File
	1, // 3: proto.Bropdox.UpdateFile:input_type -> proto.File
	1, // 4: proto.Bropdox.RemoveFile:input_type -> proto.File
	1, // 5: proto.Bropdox.Notifications:input_type -> proto.File
	2, // 6: proto.Bropdox.CreateFile:output_type -> proto.FileResponse
	2, // 7: proto.Bropdox.UpdateFile:output_type -> proto.FileResponse
	2, // 8: proto.Bropdox.RemoveFile:output_type -> proto.FileResponse
	2, // 9: proto.Bropdox.Notifications:output_type -> proto.FileResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bropdox_proto_init() }
func file_bropdox_proto_init() {
	if File_bropdox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bropdox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_bropdox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileResponse); i {
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
			RawDescriptor: file_bropdox_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bropdox_proto_goTypes,
		DependencyIndexes: file_bropdox_proto_depIdxs,
		EnumInfos:         file_bropdox_proto_enumTypes,
		MessageInfos:      file_bropdox_proto_msgTypes,
	}.Build()
	File_bropdox_proto = out.File
	file_bropdox_proto_rawDesc = nil
	file_bropdox_proto_goTypes = nil
	file_bropdox_proto_depIdxs = nil
}
