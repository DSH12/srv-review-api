// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: ozonmp/srv_review_api/v1/srv_review_api.proto

package srv_review_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Foo     uint64                 `protobuf:"varint,2,opt,name=foo,proto3" json:"foo,omitempty"`
	Created *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_srv_review_api_v1_srv_review_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_srv_review_api_v1_srv_review_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_ozonmp_srv_review_api_v1_srv_review_api_proto_rawDescGZIP(), []int{0}
}

func (x *Review) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Review) GetFoo() uint64 {
	if x != nil {
		return x.Foo
	}
	return 0
}

func (x *Review) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

type DescribeReviewV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReviewId uint64 `protobuf:"varint,1,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
}

func (x *DescribeReviewV1Request) Reset() {
	*x = DescribeReviewV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_srv_review_api_v1_srv_review_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeReviewV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeReviewV1Request) ProtoMessage() {}

func (x *DescribeReviewV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_srv_review_api_v1_srv_review_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeReviewV1Request.ProtoReflect.Descriptor instead.
func (*DescribeReviewV1Request) Descriptor() ([]byte, []int) {
	return file_ozonmp_srv_review_api_v1_srv_review_api_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeReviewV1Request) GetReviewId() uint64 {
	if x != nil {
		return x.ReviewId
	}
	return 0
}

type DescribeReviewV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *Review `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DescribeReviewV1Response) Reset() {
	*x = DescribeReviewV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozonmp_srv_review_api_v1_srv_review_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeReviewV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeReviewV1Response) ProtoMessage() {}

func (x *DescribeReviewV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ozonmp_srv_review_api_v1_srv_review_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeReviewV1Response.ProtoReflect.Descriptor instead.
func (*DescribeReviewV1Response) Descriptor() ([]byte, []int) {
	return file_ozonmp_srv_review_api_v1_srv_review_api_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeReviewV1Response) GetValue() *Review {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_ozonmp_srv_review_api_v1_srv_review_api_proto protoreflect.FileDescriptor

var file_ozonmp_srv_review_api_v1_srv_review_api_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2f, 0x73, 0x72, 0x76, 0x5f, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x72, 0x76, 0x5f, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x73, 0x72, 0x76, 0x5f, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x60, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x66,
	0x6f, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x34, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x22, 0x3f, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x09, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x73, 0x72, 0x76, 0x5f, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xb2, 0x01, 0x0a, 0x13, 0x53, 0x72, 0x76,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x9a, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x56, 0x31, 0x12, 0x31, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x73,
	0x72, 0x76, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d,
	0x70, 0x2e, 0x73, 0x72, 0x76, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x2f, 0x7b, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x44, 0x5a,
	0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e,
	0x6d, 0x70, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x2d, 0x61, 0x70, 0x69, 0x3b, 0x73, 0x72, 0x76, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ozonmp_srv_review_api_v1_srv_review_api_proto_rawDescOnce sync.Once
	file_ozonmp_srv_review_api_v1_srv_review_api_proto_rawDescData = file_ozonmp_srv_review_api_v1_srv_review_api_proto_rawDesc
)

func file_ozonmp_srv_review_api_v1_srv_review_api_proto_rawDescGZIP() []byte {
	file_ozonmp_srv_review_api_v1_srv_review_api_proto_rawDescOnce.Do(func() {
		file_ozonmp_srv_review_api_v1_srv_review_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_ozonmp_srv_review_api_v1_srv_review_api_proto_rawDescData)
	})
	return file_ozonmp_srv_review_api_v1_srv_review_api_proto_rawDescData
}

var file_ozonmp_srv_review_api_v1_srv_review_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ozonmp_srv_review_api_v1_srv_review_api_proto_goTypes = []interface{}{
	(*Review)(nil),                   // 0: ozonmp.srv_review_api.v1.Review
	(*DescribeReviewV1Request)(nil),  // 1: ozonmp.srv_review_api.v1.DescribeReviewV1Request
	(*DescribeReviewV1Response)(nil), // 2: ozonmp.srv_review_api.v1.DescribeReviewV1Response
	(*timestamppb.Timestamp)(nil),    // 3: google.protobuf.Timestamp
}
var file_ozonmp_srv_review_api_v1_srv_review_api_proto_depIdxs = []int32{
	3, // 0: ozonmp.srv_review_api.v1.Review.created:type_name -> google.protobuf.Timestamp
	0, // 1: ozonmp.srv_review_api.v1.DescribeReviewV1Response.value:type_name -> ozonmp.srv_review_api.v1.Review
	1, // 2: ozonmp.srv_review_api.v1.SrvReviewApiService.DescribeReviewV1:input_type -> ozonmp.srv_review_api.v1.DescribeReviewV1Request
	2, // 3: ozonmp.srv_review_api.v1.SrvReviewApiService.DescribeReviewV1:output_type -> ozonmp.srv_review_api.v1.DescribeReviewV1Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ozonmp_srv_review_api_v1_srv_review_api_proto_init() }
func file_ozonmp_srv_review_api_v1_srv_review_api_proto_init() {
	if File_ozonmp_srv_review_api_v1_srv_review_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ozonmp_srv_review_api_v1_srv_review_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Review); i {
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
		file_ozonmp_srv_review_api_v1_srv_review_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeReviewV1Request); i {
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
		file_ozonmp_srv_review_api_v1_srv_review_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeReviewV1Response); i {
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
			RawDescriptor: file_ozonmp_srv_review_api_v1_srv_review_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ozonmp_srv_review_api_v1_srv_review_api_proto_goTypes,
		DependencyIndexes: file_ozonmp_srv_review_api_v1_srv_review_api_proto_depIdxs,
		MessageInfos:      file_ozonmp_srv_review_api_v1_srv_review_api_proto_msgTypes,
	}.Build()
	File_ozonmp_srv_review_api_v1_srv_review_api_proto = out.File
	file_ozonmp_srv_review_api_v1_srv_review_api_proto_rawDesc = nil
	file_ozonmp_srv_review_api_v1_srv_review_api_proto_goTypes = nil
	file_ozonmp_srv_review_api_v1_srv_review_api_proto_depIdxs = nil
}
