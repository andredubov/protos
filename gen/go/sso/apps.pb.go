// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: sso/apps.proto

package ssov1

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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName   string `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`       // app's name
	AppSecret string `protobuf:"bytes,2,opt,name=app_secret,json=appSecret,proto3" json:"app_secret,omitempty"` // app's jwt token signing key
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_apps_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_apps_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_sso_apps_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *CreateRequest) GetAppSecret() string {
	if x != nil {
		return x.AppSecret
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"` // User ID of the registered user.
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_apps_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_apps_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_sso_apps_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

var File_sso_apps_proto protoreflect.FileDescriptor

var file_sso_apps_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x73, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x61, 0x70, 0x70, 0x73, 0x22, 0x49, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x22, 0x27, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x32, 0x3b, 0x0a, 0x04, 0x41, 0x70,
	0x70, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x61,
	0x70, 0x70, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17, 0x61, 0x6e, 0x64, 0x72, 0x65,
	0x64, 0x75, 0x62, 0x6f, 0x76, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x3b, 0x73, 0x73, 0x6f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sso_apps_proto_rawDescOnce sync.Once
	file_sso_apps_proto_rawDescData = file_sso_apps_proto_rawDesc
)

func file_sso_apps_proto_rawDescGZIP() []byte {
	file_sso_apps_proto_rawDescOnce.Do(func() {
		file_sso_apps_proto_rawDescData = protoimpl.X.CompressGZIP(file_sso_apps_proto_rawDescData)
	})
	return file_sso_apps_proto_rawDescData
}

var file_sso_apps_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sso_apps_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),  // 0: apps.CreateRequest
	(*CreateResponse)(nil), // 1: apps.CreateResponse
}
var file_sso_apps_proto_depIdxs = []int32{
	0, // 0: apps.Apps.Create:input_type -> apps.CreateRequest
	1, // 1: apps.Apps.Create:output_type -> apps.CreateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sso_apps_proto_init() }
func file_sso_apps_proto_init() {
	if File_sso_apps_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sso_apps_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_sso_apps_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
			RawDescriptor: file_sso_apps_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_apps_proto_goTypes,
		DependencyIndexes: file_sso_apps_proto_depIdxs,
		MessageInfos:      file_sso_apps_proto_msgTypes,
	}.Build()
	File_sso_apps_proto = out.File
	file_sso_apps_proto_rawDesc = nil
	file_sso_apps_proto_goTypes = nil
	file_sso_apps_proto_depIdxs = nil
}
