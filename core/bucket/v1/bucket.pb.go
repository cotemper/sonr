// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: core/bucket/v1/bucket.proto

// Package Service defines Service configuration for a Application on the Sonr network.

package v1

import (
	v1 "github.com/sonr-io/sonr/core/object/v1"
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

// BucketType is the type of a bucket.
type BucketType int32

const (
	// BucketTypeUnspecified is the default value.
	BucketType_BUCKET_TYPE_UNSPECIFIED BucketType = 0
	// BucketTypeApp is an App specific bucket. For Assets regarding the service.
	BucketType_BUCKET_TYPE_APP BucketType = 1
	// BucketTypeUser is a User specific bucket. For any remote user data that is required
	// to be stored in the Network.
	BucketType_BUCKET_TYPE_USER BucketType = 2
)

// Enum value maps for BucketType.
var (
	BucketType_name = map[int32]string{
		0: "BUCKET_TYPE_UNSPECIFIED",
		1: "BUCKET_TYPE_APP",
		2: "BUCKET_TYPE_USER",
	}
	BucketType_value = map[string]int32{
		"BUCKET_TYPE_UNSPECIFIED": 0,
		"BUCKET_TYPE_APP":         1,
		"BUCKET_TYPE_USER":        2,
	}
)

func (x BucketType) Enum() *BucketType {
	p := new(BucketType)
	*p = x
	return p
}

func (x BucketType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BucketType) Descriptor() protoreflect.EnumDescriptor {
	return file_core_bucket_v1_bucket_proto_enumTypes[0].Descriptor()
}

func (BucketType) Type() protoreflect.EnumType {
	return &file_core_bucket_v1_bucket_proto_enumTypes[0]
}

func (x BucketType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BucketType.Descriptor instead.
func (BucketType) EnumDescriptor() ([]byte, []int) {
	return file_core_bucket_v1_bucket_proto_rawDescGZIP(), []int{0}
}

// Bucket is a collection of objects.
type Bucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Label is human-readable name of the bucket.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// Description is a human-readable description of the bucket.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Type is the kind of bucket for either App specific or User specific data.
	Type BucketType `protobuf:"varint,3,opt,name=type,proto3,enum=core.bucket.v1.BucketType" json:"type,omitempty"`
	// Did is the identifier of the bucket.
	Did string `protobuf:"bytes,4,opt,name=did,proto3" json:"did,omitempty"`
	// Objects are stored in a tree structure.
	Objects []*v1.ObjectDoc `protobuf:"bytes,5,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *Bucket) Reset() {
	*x = Bucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_bucket_v1_bucket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucket) ProtoMessage() {}

func (x *Bucket) ProtoReflect() protoreflect.Message {
	mi := &file_core_bucket_v1_bucket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucket.ProtoReflect.Descriptor instead.
func (*Bucket) Descriptor() ([]byte, []int) {
	return file_core_bucket_v1_bucket_proto_rawDescGZIP(), []int{0}
}

func (x *Bucket) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Bucket) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Bucket) GetType() BucketType {
	if x != nil {
		return x.Type
	}
	return BucketType_BUCKET_TYPE_UNSPECIFIED
}

func (x *Bucket) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *Bucket) GetObjects() []*v1.ObjectDoc {
	if x != nil {
		return x.Objects
	}
	return nil
}

var File_core_bucket_v1_bucket_proto protoreflect.FileDescriptor

var file_core_bucket_v1_bucket_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x06, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12,
	0x33, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x63, 0x52, 0x07, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2a, 0x54, 0x0a, 0x0a, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41,
	0x50, 0x50, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x02, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f,
	0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_bucket_v1_bucket_proto_rawDescOnce sync.Once
	file_core_bucket_v1_bucket_proto_rawDescData = file_core_bucket_v1_bucket_proto_rawDesc
)

func file_core_bucket_v1_bucket_proto_rawDescGZIP() []byte {
	file_core_bucket_v1_bucket_proto_rawDescOnce.Do(func() {
		file_core_bucket_v1_bucket_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_bucket_v1_bucket_proto_rawDescData)
	})
	return file_core_bucket_v1_bucket_proto_rawDescData
}

var file_core_bucket_v1_bucket_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_bucket_v1_bucket_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_core_bucket_v1_bucket_proto_goTypes = []interface{}{
	(BucketType)(0),      // 0: core.bucket.v1.BucketType
	(*Bucket)(nil),       // 1: core.bucket.v1.Bucket
	(*v1.ObjectDoc)(nil), // 2: core.object.v1.ObjectDoc
}
var file_core_bucket_v1_bucket_proto_depIdxs = []int32{
	0, // 0: core.bucket.v1.Bucket.type:type_name -> core.bucket.v1.BucketType
	2, // 1: core.bucket.v1.Bucket.objects:type_name -> core.object.v1.ObjectDoc
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_core_bucket_v1_bucket_proto_init() }
func file_core_bucket_v1_bucket_proto_init() {
	if File_core_bucket_v1_bucket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_bucket_v1_bucket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bucket); i {
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
			RawDescriptor: file_core_bucket_v1_bucket_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_bucket_v1_bucket_proto_goTypes,
		DependencyIndexes: file_core_bucket_v1_bucket_proto_depIdxs,
		EnumInfos:         file_core_bucket_v1_bucket_proto_enumTypes,
		MessageInfos:      file_core_bucket_v1_bucket_proto_msgTypes,
	}.Build()
	File_core_bucket_v1_bucket_proto = out.File
	file_core_bucket_v1_bucket_proto_rawDesc = nil
	file_core_bucket_v1_bucket_proto_goTypes = nil
	file_core_bucket_v1_bucket_proto_depIdxs = nil
}
