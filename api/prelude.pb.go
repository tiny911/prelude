// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.24.3
// source: api/prelude.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// 服务错误码
// 下列错误码会同步至commont/error/error.go文件中
type ErrCode int32

const (
	ErrCode_ErrOk          ErrCode = 0   //成功
	ErrCode_ErrUnknown     ErrCode = 101 //未知错误
	ErrCode_ErrArgsInvalid ErrCode = 102 //参数异常
	ErrCode_ErrArgsEmpty   ErrCode = 103 //参数为空
	ErrCode_ErrSystem      ErrCode = 104 //系统错误
	ErrCode_ErrDB          ErrCode = 105 //数据库错误
	ErrCode_ErrNoServe     ErrCode = 106 //未提供服务
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0:   "ErrOk",
		101: "ErrUnknown",
		102: "ErrArgsInvalid",
		103: "ErrArgsEmpty",
		104: "ErrSystem",
		105: "ErrDB",
		106: "ErrNoServe",
	}
	ErrCode_value = map[string]int32{
		"ErrOk":          0,
		"ErrUnknown":     101,
		"ErrArgsInvalid": 102,
		"ErrArgsEmpty":   103,
		"ErrSystem":      104,
		"ErrDB":          105,
		"ErrNoServe":     106,
	}
)

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_api_prelude_proto_enumTypes[0].Descriptor()
}

func (ErrCode) Type() protoreflect.EnumType {
	return &file_api_prelude_proto_enumTypes[0]
}

func (x ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrCode.Descriptor instead.
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_api_prelude_proto_rawDescGZIP(), []int{0}
}

// Ping请求
type STPreludePingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` //名称
}

func (x *STPreludePingReq) Reset() {
	*x = STPreludePingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prelude_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *STPreludePingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*STPreludePingReq) ProtoMessage() {}

func (x *STPreludePingReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_prelude_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use STPreludePingReq.ProtoReflect.Descriptor instead.
func (*STPreludePingReq) Descriptor() ([]byte, []int) {
	return file_api_prelude_proto_rawDescGZIP(), []int{0}
}

func (x *STPreludePingReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Ping响应
type STPreludePingRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId string `protobuf:"bytes,1,opt,name=traceId,proto3" json:"traceId,omitempty"` //trace标识
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` //响应信息
}

func (x *STPreludePingRsp) Reset() {
	*x = STPreludePingRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prelude_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *STPreludePingRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*STPreludePingRsp) ProtoMessage() {}

func (x *STPreludePingRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_prelude_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use STPreludePingRsp.ProtoReflect.Descriptor instead.
func (*STPreludePingRsp) Descriptor() ([]byte, []int) {
	return file_api_prelude_proto_rawDescGZIP(), []int{1}
}

func (x *STPreludePingRsp) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *STPreludePingRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_prelude_proto protoreflect.FileDescriptor

var file_api_prelude_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x65, 0x6c, 0x75, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x65, 0x6c, 0x75, 0x64, 0x65, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x53, 0x54,
	0x50, 0x72, 0x65, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x46, 0x0a, 0x10, 0x53, 0x54, 0x50, 0x72, 0x65, 0x6c, 0x75, 0x64, 0x65, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x74, 0x0a, 0x07, 0x45, 0x72,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x4f, 0x6b, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x72, 0x72, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x65,
	0x12, 0x12, 0x0a, 0x0e, 0x45, 0x72, 0x72, 0x41, 0x72, 0x67, 0x73, 0x49, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x10, 0x66, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x41, 0x72, 0x67, 0x73, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x10, 0x67, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x10, 0x68, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x44, 0x42, 0x10, 0x69,
	0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x72, 0x72, 0x4e, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x10, 0x6a,
	0x32, 0x5d, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x6c, 0x75, 0x64, 0x65, 0x12, 0x52, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x65, 0x6c, 0x75, 0x64, 0x65, 0x2e, 0x53, 0x54,
	0x50, 0x72, 0x65, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x19,
	0x2e, 0x70, 0x72, 0x65, 0x6c, 0x75, 0x64, 0x65, 0x2e, 0x53, 0x54, 0x50, 0x72, 0x65, 0x6c, 0x75,
	0x64, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x12, 0x0c, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42,
	0x1d, 0x5a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_prelude_proto_rawDescOnce sync.Once
	file_api_prelude_proto_rawDescData = file_api_prelude_proto_rawDesc
)

func file_api_prelude_proto_rawDescGZIP() []byte {
	file_api_prelude_proto_rawDescOnce.Do(func() {
		file_api_prelude_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_prelude_proto_rawDescData)
	})
	return file_api_prelude_proto_rawDescData
}

var file_api_prelude_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_prelude_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_prelude_proto_goTypes = []interface{}{
	(ErrCode)(0),             // 0: prelude.ErrCode
	(*STPreludePingReq)(nil), // 1: prelude.STPreludePingReq
	(*STPreludePingRsp)(nil), // 2: prelude.STPreludePingRsp
}
var file_api_prelude_proto_depIdxs = []int32{
	1, // 0: prelude.Prelude.Ping:input_type -> prelude.STPreludePingReq
	2, // 1: prelude.Prelude.Ping:output_type -> prelude.STPreludePingRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_prelude_proto_init() }
func file_api_prelude_proto_init() {
	if File_api_prelude_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_prelude_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*STPreludePingReq); i {
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
		file_api_prelude_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*STPreludePingRsp); i {
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
			RawDescriptor: file_api_prelude_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_prelude_proto_goTypes,
		DependencyIndexes: file_api_prelude_proto_depIdxs,
		EnumInfos:         file_api_prelude_proto_enumTypes,
		MessageInfos:      file_api_prelude_proto_msgTypes,
	}.Build()
	File_api_prelude_proto = out.File
	file_api_prelude_proto_rawDesc = nil
	file_api_prelude_proto_goTypes = nil
	file_api_prelude_proto_depIdxs = nil
}
