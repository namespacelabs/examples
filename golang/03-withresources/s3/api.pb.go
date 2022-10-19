// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: golang/03-withresources/s3/api.proto

package s3

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

type BucketIntent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region     string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	BucketName string `protobuf:"bytes,2,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
}

func (x *BucketIntent) Reset() {
	*x = BucketIntent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_golang_03_withresources_s3_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BucketIntent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketIntent) ProtoMessage() {}

func (x *BucketIntent) ProtoReflect() protoreflect.Message {
	mi := &file_golang_03_withresources_s3_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketIntent.ProtoReflect.Descriptor instead.
func (*BucketIntent) Descriptor() ([]byte, []int) {
	return file_golang_03_withresources_s3_api_proto_rawDescGZIP(), []int{0}
}

func (x *BucketIntent) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *BucketIntent) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

type BucketInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region          string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	BucketName      string `protobuf:"bytes,2,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	Url             string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	AccessKey       string `protobuf:"bytes,4,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretAccessKey string `protobuf:"bytes,5,opt,name=secret_access_key,json=secretAccessKey,proto3" json:"secret_access_key,omitempty"`
}

func (x *BucketInstance) Reset() {
	*x = BucketInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_golang_03_withresources_s3_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BucketInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketInstance) ProtoMessage() {}

func (x *BucketInstance) ProtoReflect() protoreflect.Message {
	mi := &file_golang_03_withresources_s3_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketInstance.ProtoReflect.Descriptor instead.
func (*BucketInstance) Descriptor() ([]byte, []int) {
	return file_golang_03_withresources_s3_api_proto_rawDescGZIP(), []int{1}
}

func (x *BucketInstance) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *BucketInstance) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *BucketInstance) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *BucketInstance) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *BucketInstance) GetSecretAccessKey() string {
	if x != nil {
		return x.SecretAccessKey
	}
	return ""
}

var File_golang_03_withresources_s3_api_proto protoreflect.FileDescriptor

var file_golang_03_withresources_s3_api_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x30, 0x33, 0x2d, 0x77, 0x69, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x33, 0x2f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x33, 0x22, 0x47, 0x0a, 0x0c, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0xa6, 0x01, 0x0a, 0x0e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x11, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x42, 0x37, 0x5a, 0x35, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x30, 0x33, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x73, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_golang_03_withresources_s3_api_proto_rawDescOnce sync.Once
	file_golang_03_withresources_s3_api_proto_rawDescData = file_golang_03_withresources_s3_api_proto_rawDesc
)

func file_golang_03_withresources_s3_api_proto_rawDescGZIP() []byte {
	file_golang_03_withresources_s3_api_proto_rawDescOnce.Do(func() {
		file_golang_03_withresources_s3_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_golang_03_withresources_s3_api_proto_rawDescData)
	})
	return file_golang_03_withresources_s3_api_proto_rawDescData
}

var file_golang_03_withresources_s3_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_golang_03_withresources_s3_api_proto_goTypes = []interface{}{
	(*BucketIntent)(nil),   // 0: examples.golang.withresources.s3.BucketIntent
	(*BucketInstance)(nil), // 1: examples.golang.withresources.s3.BucketInstance
}
var file_golang_03_withresources_s3_api_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_golang_03_withresources_s3_api_proto_init() }
func file_golang_03_withresources_s3_api_proto_init() {
	if File_golang_03_withresources_s3_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_golang_03_withresources_s3_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketIntent); i {
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
		file_golang_03_withresources_s3_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketInstance); i {
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
			RawDescriptor: file_golang_03_withresources_s3_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_golang_03_withresources_s3_api_proto_goTypes,
		DependencyIndexes: file_golang_03_withresources_s3_api_proto_depIdxs,
		MessageInfos:      file_golang_03_withresources_s3_api_proto_msgTypes,
	}.Build()
	File_golang_03_withresources_s3_api_proto = out.File
	file_golang_03_withresources_s3_api_proto_rawDesc = nil
	file_golang_03_withresources_s3_api_proto_goTypes = nil
	file_golang_03_withresources_s3_api_proto_depIdxs = nil
}
