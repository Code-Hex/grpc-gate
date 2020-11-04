// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: internal/proto/gate.proto

package proto

import (
	context "context"
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

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_gate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_gate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_internal_proto_gate_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_internal_proto_gate_proto protoreflect.FileDescriptor

var file_internal_proto_gate_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32,
	0x38, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2e, 0x0a, 0x0c, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x28, 0x01, 0x30, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x64, 0x65, 0x2d, 0x48, 0x65, 0x78,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_gate_proto_rawDescOnce sync.Once
	file_internal_proto_gate_proto_rawDescData = file_internal_proto_gate_proto_rawDesc
)

func file_internal_proto_gate_proto_rawDescGZIP() []byte {
	file_internal_proto_gate_proto_rawDescOnce.Do(func() {
		file_internal_proto_gate_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_gate_proto_rawDescData)
	})
	return file_internal_proto_gate_proto_rawDescData
}

var file_internal_proto_gate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_proto_gate_proto_goTypes = []interface{}{
	(*Chunk)(nil), // 0: proto.Chunk
}
var file_internal_proto_gate_proto_depIdxs = []int32{
	0, // 0: proto.Stream.ServerStream:input_type -> proto.Chunk
	0, // 1: proto.Stream.ServerStream:output_type -> proto.Chunk
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_gate_proto_init() }
func file_internal_proto_gate_proto_init() {
	if File_internal_proto_gate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_gate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
			RawDescriptor: file_internal_proto_gate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_gate_proto_goTypes,
		DependencyIndexes: file_internal_proto_gate_proto_depIdxs,
		MessageInfos:      file_internal_proto_gate_proto_msgTypes,
	}.Build()
	File_internal_proto_gate_proto = out.File
	file_internal_proto_gate_proto_rawDesc = nil
	file_internal_proto_gate_proto_goTypes = nil
	file_internal_proto_gate_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamClient is the client API for Stream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamClient interface {
	ServerStream(ctx context.Context, opts ...grpc.CallOption) (Stream_ServerStreamClient, error)
}

type streamClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamClient(cc grpc.ClientConnInterface) StreamClient {
	return &streamClient{cc}
}

func (c *streamClient) ServerStream(ctx context.Context, opts ...grpc.CallOption) (Stream_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Stream_serviceDesc.Streams[0], "/proto.Stream/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServerStreamClient{stream}
	return x, nil
}

type Stream_ServerStreamClient interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type streamServerStreamClient struct {
	grpc.ClientStream
}

func (x *streamServerStreamClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServerStreamClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServer is the server API for Stream service.
type StreamServer interface {
	ServerStream(Stream_ServerStreamServer) error
}

// UnimplementedStreamServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServer struct {
}

func (*UnimplementedStreamServer) ServerStream(Stream_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}

func RegisterStreamServer(s *grpc.Server, srv StreamServer) {
	s.RegisterService(&_Stream_serviceDesc, srv)
}

func _Stream_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServer).ServerStream(&streamServerStreamServer{stream})
}

type Stream_ServerStreamServer interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type streamServerStreamServer struct {
	grpc.ServerStream
}

func (x *streamServerStreamServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServerStreamServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Stream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Stream",
	HandlerType: (*StreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _Stream_ServerStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/proto/gate.proto",
}