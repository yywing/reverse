// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reverse.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ProtocolEnum int32

const (
	ProtocolEnum_GRPC ProtocolEnum = 0
)

var ProtocolEnum_name = map[int32]string{
	0: "GRPC",
}

var ProtocolEnum_value = map[string]int32{
	"GRPC": 0,
}

func (x ProtocolEnum) String() string {
	return proto.EnumName(ProtocolEnum_name, int32(x))
}

func (ProtocolEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_142d12653d52cef8, []int{0}
}

type ConnectRequest struct {
	IP                   string       `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
	Port                 int32        `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
	ServiceName          string       `protobuf:"bytes,3,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	Protocol             ProtocolEnum `protobuf:"varint,4,opt,name=Protocol,proto3,enum=pb.ProtocolEnum" json:"Protocol,omitempty"`
	Timeout              int32        `protobuf:"varint,5,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_142d12653d52cef8, []int{0}
}

func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *ConnectRequest) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ConnectRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ConnectRequest) GetProtocol() ProtocolEnum {
	if m != nil {
		return m.Protocol
	}
	return ProtocolEnum_GRPC
}

func (m *ConnectRequest) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type ClientConnectRequest struct {
	// server ID unique
	ID                   string          `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Request              *ConnectRequest `protobuf:"bytes,2,opt,name=Request,proto3" json:"Request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ClientConnectRequest) Reset()         { *m = ClientConnectRequest{} }
func (m *ClientConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ClientConnectRequest) ProtoMessage()    {}
func (*ClientConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_142d12653d52cef8, []int{1}
}

func (m *ClientConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientConnectRequest.Unmarshal(m, b)
}
func (m *ClientConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientConnectRequest.Marshal(b, m, deterministic)
}
func (m *ClientConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientConnectRequest.Merge(m, src)
}
func (m *ClientConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ClientConnectRequest.Size(m)
}
func (m *ClientConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientConnectRequest proto.InternalMessageInfo

func (m *ClientConnectRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ClientConnectRequest) GetRequest() *ConnectRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

type ConnectResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectResponse) Reset()         { *m = ConnectResponse{} }
func (m *ConnectResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectResponse) ProtoMessage()    {}
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_142d12653d52cef8, []int{2}
}

func (m *ConnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectResponse.Unmarshal(m, b)
}
func (m *ConnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectResponse.Marshal(b, m, deterministic)
}
func (m *ConnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectResponse.Merge(m, src)
}
func (m *ConnectResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectResponse.Size(m)
}
func (m *ConnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectResponse proto.InternalMessageInfo

type RegisterMessage struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterMessage) Reset()         { *m = RegisterMessage{} }
func (m *RegisterMessage) String() string { return proto.CompactTextString(m) }
func (*RegisterMessage) ProtoMessage()    {}
func (*RegisterMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_142d12653d52cef8, []int{3}
}

func (m *RegisterMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterMessage.Unmarshal(m, b)
}
func (m *RegisterMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterMessage.Marshal(b, m, deterministic)
}
func (m *RegisterMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterMessage.Merge(m, src)
}
func (m *RegisterMessage) XXX_Size() int {
	return xxx_messageInfo_RegisterMessage.Size(m)
}
func (m *RegisterMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterMessage proto.InternalMessageInfo

func (m *RegisterMessage) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterEnum("pb.ProtocolEnum", ProtocolEnum_name, ProtocolEnum_value)
	proto.RegisterType((*ConnectRequest)(nil), "pb.ConnectRequest")
	proto.RegisterType((*ClientConnectRequest)(nil), "pb.ClientConnectRequest")
	proto.RegisterType((*ConnectResponse)(nil), "pb.ConnectResponse")
	proto.RegisterType((*RegisterMessage)(nil), "pb.RegisterMessage")
}

func init() { proto.RegisterFile("reverse.proto", fileDescriptor_142d12653d52cef8) }

var fileDescriptor_142d12653d52cef8 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x6e, 0xf2, 0x30,
	0x10, 0xc5, 0x71, 0x3e, 0xf8, 0xa0, 0x43, 0xf9, 0xd3, 0xa1, 0x0b, 0x8b, 0x55, 0x9a, 0x15, 0xaa,
	0x10, 0xaa, 0xd2, 0x6d, 0x17, 0x95, 0x42, 0x55, 0xb1, 0x68, 0x15, 0xb9, 0x5c, 0x80, 0x44, 0x23,
	0x14, 0x89, 0xc4, 0xa9, 0xed, 0x70, 0x87, 0x5e, 0xa2, 0x67, 0xad, 0x62, 0x25, 0x11, 0x41, 0xd9,
	0xd9, 0x33, 0xcf, 0xf3, 0x7e, 0xf3, 0x0c, 0x13, 0x45, 0x67, 0x52, 0x9a, 0x36, 0xb9, 0x92, 0x46,
	0xa2, 0x93, 0x47, 0xde, 0x2f, 0x83, 0x69, 0x20, 0xb3, 0x8c, 0x62, 0x23, 0xe8, 0xbb, 0x20, 0x6d,
	0x70, 0x0a, 0xce, 0x2e, 0xe4, 0xcc, 0x65, 0xab, 0x1b, 0xe1, 0xec, 0x42, 0x44, 0xe8, 0x87, 0x52,
	0x19, 0xee, 0xb8, 0x6c, 0x35, 0x10, 0xf6, 0x8c, 0x2e, 0x8c, 0xbf, 0x48, 0x9d, 0x93, 0x98, 0x3e,
	0x0f, 0x29, 0xf1, 0x7f, 0x56, 0x7c, 0x59, 0xc2, 0x35, 0x8c, 0xc2, 0xd2, 0x25, 0x96, 0x27, 0xde,
	0x77, 0xd9, 0x6a, 0xea, 0xcf, 0x37, 0x79, 0xb4, 0xa9, 0x6b, 0x6f, 0x59, 0x91, 0x8a, 0x46, 0x81,
	0x1c, 0x86, 0xfb, 0x24, 0x25, 0x59, 0x18, 0x3e, 0xb0, 0x36, 0xf5, 0xd5, 0xdb, 0xc3, 0x7d, 0x70,
	0x4a, 0x28, 0x33, 0x1d, 0x94, 0xdb, 0x86, 0x72, 0x8b, 0x6b, 0x18, 0x56, 0x2d, 0x0b, 0x3a, 0xf6,
	0xb1, 0xb4, 0x6b, 0x3f, 0x12, 0xb5, 0xc4, 0xbb, 0x83, 0x59, 0xd3, 0xd2, 0xb9, 0xcc, 0x34, 0x79,
	0x0f, 0x30, 0x13, 0x74, 0x4c, 0xb4, 0x21, 0xf5, 0x41, 0x5a, 0x1f, 0x8e, 0x74, 0xed, 0xf1, 0xc8,
	0xe1, 0xf6, 0x92, 0x1f, 0x47, 0xd0, 0x7f, 0x17, 0x61, 0x30, 0xef, 0xf9, 0x3f, 0xac, 0xb4, 0xb7,
	0xe1, 0xe2, 0x0b, 0x4c, 0xca, 0x20, 0x48, 0x55, 0x0e, 0xb8, 0x28, 0x49, 0xae, 0x66, 0x2f, 0x3b,
	0xf0, 0xbc, 0xde, 0x13, 0xc3, 0x57, 0x98, 0xb4, 0xf6, 0x45, 0x6e, 0x85, 0x1d, 0x11, 0x2c, 0x17,
	0xad, 0x11, 0xd5, 0x1a, 0xbd, 0xe8, 0xbf, 0xfd, 0xdd, 0xe7, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd0, 0x84, 0x53, 0x33, 0xee, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReverseClient is the client API for Reverse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReverseClient interface {
	ServerConnect(ctx context.Context, in *RegisterMessage, opts ...grpc.CallOption) (Reverse_ServerConnectClient, error)
	ClientConnect(ctx context.Context, in *ClientConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
}

type reverseClient struct {
	cc *grpc.ClientConn
}

func NewReverseClient(cc *grpc.ClientConn) ReverseClient {
	return &reverseClient{cc}
}

func (c *reverseClient) ServerConnect(ctx context.Context, in *RegisterMessage, opts ...grpc.CallOption) (Reverse_ServerConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Reverse_serviceDesc.Streams[0], "/pb.Reverse/ServerConnect", opts...)
	if err != nil {
		return nil, err
	}
	x := &reverseServerConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Reverse_ServerConnectClient interface {
	Recv() (*ConnectRequest, error)
	grpc.ClientStream
}

type reverseServerConnectClient struct {
	grpc.ClientStream
}

func (x *reverseServerConnectClient) Recv() (*ConnectRequest, error) {
	m := new(ConnectRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *reverseClient) ClientConnect(ctx context.Context, in *ClientConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, "/pb.Reverse/ClientConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReverseServer is the server API for Reverse service.
type ReverseServer interface {
	ServerConnect(*RegisterMessage, Reverse_ServerConnectServer) error
	ClientConnect(context.Context, *ClientConnectRequest) (*ConnectResponse, error)
}

// UnimplementedReverseServer can be embedded to have forward compatible implementations.
type UnimplementedReverseServer struct {
}

func (*UnimplementedReverseServer) ServerConnect(req *RegisterMessage, srv Reverse_ServerConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerConnect not implemented")
}
func (*UnimplementedReverseServer) ClientConnect(ctx context.Context, req *ClientConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientConnect not implemented")
}

func RegisterReverseServer(s *grpc.Server, srv ReverseServer) {
	s.RegisterService(&_Reverse_serviceDesc, srv)
}

func _Reverse_ServerConnect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RegisterMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReverseServer).ServerConnect(m, &reverseServerConnectServer{stream})
}

type Reverse_ServerConnectServer interface {
	Send(*ConnectRequest) error
	grpc.ServerStream
}

type reverseServerConnectServer struct {
	grpc.ServerStream
}

func (x *reverseServerConnectServer) Send(m *ConnectRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _Reverse_ClientConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseServer).ClientConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Reverse/ClientConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseServer).ClientConnect(ctx, req.(*ClientConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Reverse_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Reverse",
	HandlerType: (*ReverseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClientConnect",
			Handler:    _Reverse_ClientConnect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerConnect",
			Handler:       _Reverse_ServerConnect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "reverse.proto",
}