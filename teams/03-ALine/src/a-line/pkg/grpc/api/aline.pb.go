// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: pkg/grpc/api/aline.proto

package api

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

type AlineMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 1: register 2: offline 3: heartbeat  4: execute 5: cancel 6: executeResultNotify 7: getLog
	Type int64 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// registry
	Address string      `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	ExecReq *ExecuteReq `protobuf:"bytes,3,opt,name=execReq,proto3" json:"execReq,omitempty"`
	// execute result
	Result string `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	// log
	Log *Log `protobuf:"bytes,5,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *AlineMessage) Reset() {
	*x = AlineMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_api_aline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlineMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlineMessage) ProtoMessage() {}

func (x *AlineMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_api_aline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlineMessage.ProtoReflect.Descriptor instead.
func (*AlineMessage) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_api_aline_proto_rawDescGZIP(), []int{0}
}

func (x *AlineMessage) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AlineMessage) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AlineMessage) GetExecReq() *ExecuteReq {
	if x != nil {
		return x.ExecReq
	}
	return nil
}

func (x *AlineMessage) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *AlineMessage) GetLog() *Log {
	if x != nil {
		return x.Log
	}
	return nil
}

type ExecuteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pipeline file
	PipelineFile string `protobuf:"bytes,1,opt,name=pipelineFile,proto3" json:"pipelineFile,omitempty"`
	// job exec id
	JobDetailId int64 `protobuf:"varint,2,opt,name=jobDetailId,proto3" json:"jobDetailId,omitempty"`
}

func (x *ExecuteReq) Reset() {
	*x = ExecuteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_api_aline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteReq) ProtoMessage() {}

func (x *ExecuteReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_api_aline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteReq.ProtoReflect.Descriptor instead.
func (*ExecuteReq) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_api_aline_proto_rawDescGZIP(), []int{1}
}

func (x *ExecuteReq) GetPipelineFile() string {
	if x != nil {
		return x.PipelineFile
	}
	return ""
}

func (x *ExecuteReq) GetJobDetailId() int64 {
	if x != nil {
		return x.JobDetailId
	}
	return 0
}

type ExecuteResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobDetail string `protobuf:"bytes,1,opt,name=jobDetail,proto3" json:"jobDetail,omitempty"`
	StartTime string `protobuf:"bytes,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	Duration  int64  `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *ExecuteResult) Reset() {
	*x = ExecuteResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_api_aline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteResult) ProtoMessage() {}

func (x *ExecuteResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_api_aline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteResult.ProtoReflect.Descriptor instead.
func (*ExecuteResult) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_api_aline_proto_rawDescGZIP(), []int{2}
}

func (x *ExecuteResult) GetJobDetail() string {
	if x != nil {
		return x.JobDetail
	}
	return ""
}

func (x *ExecuteResult) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ExecuteResult) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stage   string `protobuf:"bytes,1,opt,name=stage,proto3" json:"stage,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	More    bool   `protobuf:"varint,3,opt,name=more,proto3" json:"more,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_api_aline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_api_aline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_api_aline_proto_rawDescGZIP(), []int{3}
}

func (x *Log) GetStage() string {
	if x != nil {
		return x.Stage
	}
	return ""
}

func (x *Log) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Log) GetMore() bool {
	if x != nil {
		return x.More
	}
	return false
}

var File_pkg_grpc_api_aline_proto protoreflect.FileDescriptor

var file_pkg_grpc_api_aline_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22,
	0x9b, 0x01, 0x0a, 0x0c, 0x41, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29,
	0x0a, 0x07, 0x65, 0x78, 0x65, 0x63, 0x52, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x52, 0x07, 0x65, 0x78, 0x65, 0x63, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x1a, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x22, 0x52, 0x0a,
	0x0a, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6a, 0x6f, 0x62, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49,
	0x64, 0x22, 0x67, 0x0a, 0x0d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x49, 0x0a, 0x03, 0x4c, 0x6f,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x6d, 0x6f, 0x72, 0x65, 0x32, 0x43, 0x0a, 0x08, 0x41, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x50,
	0x43, 0x12, 0x37, 0x0a, 0x09, 0x41, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x11,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x5e, 0x0a, 0x1f, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x68, 0x61, 0x6d, 0x73, 0x74, 0x65, 0x72,
	0x2d, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x61, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x0a, 0x41,
	0x6c, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x6d, 0x73, 0x74, 0x65, 0x72, 0x2d,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x61, 0x2d, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_grpc_api_aline_proto_rawDescOnce sync.Once
	file_pkg_grpc_api_aline_proto_rawDescData = file_pkg_grpc_api_aline_proto_rawDesc
)

func file_pkg_grpc_api_aline_proto_rawDescGZIP() []byte {
	file_pkg_grpc_api_aline_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_api_aline_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_api_aline_proto_rawDescData)
	})
	return file_pkg_grpc_api_aline_proto_rawDescData
}

var file_pkg_grpc_api_aline_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_grpc_api_aline_proto_goTypes = []interface{}{
	(*AlineMessage)(nil),  // 0: api.AlineMessage
	(*ExecuteReq)(nil),    // 1: api.ExecuteReq
	(*ExecuteResult)(nil), // 2: api.ExecuteResult
	(*Log)(nil),           // 3: api.Log
}
var file_pkg_grpc_api_aline_proto_depIdxs = []int32{
	1, // 0: api.AlineMessage.execReq:type_name -> api.ExecuteReq
	3, // 1: api.AlineMessage.log:type_name -> api.Log
	0, // 2: api.AlineRPC.AlineChat:input_type -> api.AlineMessage
	0, // 3: api.AlineRPC.AlineChat:output_type -> api.AlineMessage
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_grpc_api_aline_proto_init() }
func file_pkg_grpc_api_aline_proto_init() {
	if File_pkg_grpc_api_aline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_api_aline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlineMessage); i {
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
		file_pkg_grpc_api_aline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteReq); i {
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
		file_pkg_grpc_api_aline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteResult); i {
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
		file_pkg_grpc_api_aline_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_pkg_grpc_api_aline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_api_aline_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_api_aline_proto_depIdxs,
		MessageInfos:      file_pkg_grpc_api_aline_proto_msgTypes,
	}.Build()
	File_pkg_grpc_api_aline_proto = out.File
	file_pkg_grpc_api_aline_proto_rawDesc = nil
	file_pkg_grpc_api_aline_proto_goTypes = nil
	file_pkg_grpc_api_aline_proto_depIdxs = nil
}