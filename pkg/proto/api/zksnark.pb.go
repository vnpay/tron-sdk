// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: api/zksnark.proto

package api

import (
	context "context"
	core "github.com/vnpay/tron-sdk/pkg/proto/core"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ZksnarkResponse_Code int32

const (
	ZksnarkResponse_SUCCESS ZksnarkResponse_Code = 0
	ZksnarkResponse_FAILED  ZksnarkResponse_Code = 1
)

// Enum value maps for ZksnarkResponse_Code.
var (
	ZksnarkResponse_Code_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILED",
	}
	ZksnarkResponse_Code_value = map[string]int32{
		"SUCCESS": 0,
		"FAILED":  1,
	}
)

func (x ZksnarkResponse_Code) Enum() *ZksnarkResponse_Code {
	p := new(ZksnarkResponse_Code)
	*p = x
	return p
}

func (x ZksnarkResponse_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZksnarkResponse_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_api_zksnark_proto_enumTypes[0].Descriptor()
}

func (ZksnarkResponse_Code) Type() protoreflect.EnumType {
	return &file_api_zksnark_proto_enumTypes[0]
}

func (x ZksnarkResponse_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZksnarkResponse_Code.Descriptor instead.
func (ZksnarkResponse_Code) EnumDescriptor() ([]byte, []int) {
	return file_api_zksnark_proto_rawDescGZIP(), []int{1, 0}
}

type ZksnarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction  *core.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Sighash      []byte            `protobuf:"bytes,2,opt,name=sighash,proto3" json:"sighash,omitempty"`
	ValueBalance int64             `protobuf:"varint,3,opt,name=valueBalance,proto3" json:"valueBalance,omitempty"`
	TxId         string            `protobuf:"bytes,4,opt,name=txId,proto3" json:"txId,omitempty"`
}

func (x *ZksnarkRequest) Reset() {
	*x = ZksnarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_zksnark_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZksnarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZksnarkRequest) ProtoMessage() {}

func (x *ZksnarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_zksnark_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZksnarkRequest.ProtoReflect.Descriptor instead.
func (*ZksnarkRequest) Descriptor() ([]byte, []int) {
	return file_api_zksnark_proto_rawDescGZIP(), []int{0}
}

func (x *ZksnarkRequest) GetTransaction() *core.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *ZksnarkRequest) GetSighash() []byte {
	if x != nil {
		return x.Sighash
	}
	return nil
}

func (x *ZksnarkRequest) GetValueBalance() int64 {
	if x != nil {
		return x.ValueBalance
	}
	return 0
}

func (x *ZksnarkRequest) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

type ZksnarkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ZksnarkResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=protocol.ZksnarkResponse_Code" json:"code,omitempty"`
}

func (x *ZksnarkResponse) Reset() {
	*x = ZksnarkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_zksnark_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZksnarkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZksnarkResponse) ProtoMessage() {}

func (x *ZksnarkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_zksnark_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZksnarkResponse.ProtoReflect.Descriptor instead.
func (*ZksnarkResponse) Descriptor() ([]byte, []int) {
	return file_api_zksnark_proto_rawDescGZIP(), []int{1}
}

func (x *ZksnarkResponse) GetCode() ZksnarkResponse_Code {
	if x != nil {
		return x.Code
	}
	return ZksnarkResponse_SUCCESS
}

var File_api_zksnark_proto protoreflect.FileDescriptor

var file_api_zksnark_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x7a, 0x6b, 0x73, 0x6e, 0x61, 0x72, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x0f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x54, 0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b,
	0x01, 0x0a, 0x0e, 0x5a, 0x6b, 0x73, 0x6e, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x37, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x69,
	0x67, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x69, 0x67,
	0x68, 0x61, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x0f,
	0x5a, 0x6b, 0x73, 0x6e, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x5a, 0x6b, 0x73, 0x6e, 0x61, 0x72, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x1f, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x01, 0x32, 0x59, 0x0a, 0x0b, 0x54, 0x72, 0x6f, 0x6e, 0x5a, 0x6b, 0x73, 0x6e,
	0x61, 0x72, 0x6b, 0x12, 0x4a, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x5a, 0x6b, 0x73, 0x6e,
	0x61, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x5a, 0x6b, 0x73, 0x6e, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x5a, 0x6b,
	0x73, 0x6e, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x4e, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x42,
	0x0e, 0x5a, 0x6b, 0x73, 0x6e, 0x61, 0x72, 0x6b, 0x47, 0x72, 0x70, 0x63, 0x41, 0x50, 0x49, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x62, 0x73, 0x6f,
	0x62, 0x72, 0x65, 0x69, 0x72, 0x61, 0x2f, 0x67, 0x6f, 0x74, 0x72, 0x6f, 0x6e, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_zksnark_proto_rawDescOnce sync.Once
	file_api_zksnark_proto_rawDescData = file_api_zksnark_proto_rawDesc
)

func file_api_zksnark_proto_rawDescGZIP() []byte {
	file_api_zksnark_proto_rawDescOnce.Do(func() {
		file_api_zksnark_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_zksnark_proto_rawDescData)
	})
	return file_api_zksnark_proto_rawDescData
}

var file_api_zksnark_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_zksnark_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_zksnark_proto_goTypes = []interface{}{
	(ZksnarkResponse_Code)(0), // 0: protocol.ZksnarkResponse.Code
	(*ZksnarkRequest)(nil),    // 1: protocol.ZksnarkRequest
	(*ZksnarkResponse)(nil),   // 2: protocol.ZksnarkResponse
	(*core.Transaction)(nil),  // 3: protocol.Transaction
}
var file_api_zksnark_proto_depIdxs = []int32{
	3, // 0: protocol.ZksnarkRequest.transaction:type_name -> protocol.Transaction
	0, // 1: protocol.ZksnarkResponse.code:type_name -> protocol.ZksnarkResponse.Code
	1, // 2: protocol.TronZksnark.CheckZksnarkProof:input_type -> protocol.ZksnarkRequest
	2, // 3: protocol.TronZksnark.CheckZksnarkProof:output_type -> protocol.ZksnarkResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_zksnark_proto_init() }
func file_api_zksnark_proto_init() {
	if File_api_zksnark_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_zksnark_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZksnarkRequest); i {
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
		file_api_zksnark_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZksnarkResponse); i {
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
			RawDescriptor: file_api_zksnark_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_zksnark_proto_goTypes,
		DependencyIndexes: file_api_zksnark_proto_depIdxs,
		EnumInfos:         file_api_zksnark_proto_enumTypes,
		MessageInfos:      file_api_zksnark_proto_msgTypes,
	}.Build()
	File_api_zksnark_proto = out.File
	file_api_zksnark_proto_rawDesc = nil
	file_api_zksnark_proto_goTypes = nil
	file_api_zksnark_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TronZksnarkClient is the client API for TronZksnark service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TronZksnarkClient interface {
	CheckZksnarkProof(ctx context.Context, in *ZksnarkRequest, opts ...grpc.CallOption) (*ZksnarkResponse, error)
}

type tronZksnarkClient struct {
	cc grpc.ClientConnInterface
}

func NewTronZksnarkClient(cc grpc.ClientConnInterface) TronZksnarkClient {
	return &tronZksnarkClient{cc}
}

func (c *tronZksnarkClient) CheckZksnarkProof(ctx context.Context, in *ZksnarkRequest, opts ...grpc.CallOption) (*ZksnarkResponse, error) {
	out := new(ZksnarkResponse)
	err := c.cc.Invoke(ctx, "/protocol.TronZksnark/CheckZksnarkProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TronZksnarkServer is the server API for TronZksnark service.
type TronZksnarkServer interface {
	CheckZksnarkProof(context.Context, *ZksnarkRequest) (*ZksnarkResponse, error)
}

// UnimplementedTronZksnarkServer can be embedded to have forward compatible implementations.
type UnimplementedTronZksnarkServer struct {
}

func (*UnimplementedTronZksnarkServer) CheckZksnarkProof(context.Context, *ZksnarkRequest) (*ZksnarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckZksnarkProof not implemented")
}

func RegisterTronZksnarkServer(s *grpc.Server, srv TronZksnarkServer) {
	s.RegisterService(&_TronZksnark_serviceDesc, srv)
}

func _TronZksnark_CheckZksnarkProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZksnarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TronZksnarkServer).CheckZksnarkProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.TronZksnark/CheckZksnarkProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TronZksnarkServer).CheckZksnarkProof(ctx, req.(*ZksnarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TronZksnark_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.TronZksnark",
	HandlerType: (*TronZksnarkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckZksnarkProof",
			Handler:    _TronZksnark_CheckZksnarkProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/zksnark.proto",
}
