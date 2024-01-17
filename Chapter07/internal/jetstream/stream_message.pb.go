// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: stream_message.proto

package jetstream

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StreamMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamMessage) Reset() {
	*x = StreamMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMessage) ProtoMessage() {}

func (x *StreamMessage) ProtoReflect() protoreflect.Message {
	mi := &file_stream_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMessage.ProtoReflect.Descriptor instead.
func (*StreamMessage) Descriptor() ([]byte, []int) {
	return file_stream_message_proto_rawDescGZIP(), []int{0}
}

func (x *StreamMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StreamMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StreamMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_stream_message_proto protoreflect.FileDescriptor

var file_stream_message_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6a, 0x65, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x47, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x89, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x2e, 0x6a, 0x65, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x12, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x20, 0x65, 0x64, 0x61, 0x2d, 0x69, 0x6e, 0x2d, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6a, 0x65, 0x74, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0xa2, 0x02, 0x03, 0x4a, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x4a, 0x65,
	0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0xca, 0x02, 0x09, 0x4a, 0x65, 0x74, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0xe2, 0x02, 0x15, 0x4a, 0x65, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x4a, 0x65,
	0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stream_message_proto_rawDescOnce sync.Once
	file_stream_message_proto_rawDescData = file_stream_message_proto_rawDesc
)

func file_stream_message_proto_rawDescGZIP() []byte {
	file_stream_message_proto_rawDescOnce.Do(func() {
		file_stream_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_message_proto_rawDescData)
	})
	return file_stream_message_proto_rawDescData
}

var file_stream_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_stream_message_proto_goTypes = []interface{}{
	(*StreamMessage)(nil), // 0: jetstream.StreamMessage
}
var file_stream_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stream_message_proto_init() }
func file_stream_message_proto_init() {
	if File_stream_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamMessage); i {
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
			RawDescriptor: file_stream_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stream_message_proto_goTypes,
		DependencyIndexes: file_stream_message_proto_depIdxs,
		MessageInfos:      file_stream_message_proto_msgTypes,
	}.Build()
	File_stream_message_proto = out.File
	file_stream_message_proto_rawDesc = nil
	file_stream_message_proto_goTypes = nil
	file_stream_message_proto_depIdxs = nil
}
