// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go_rpc.proto

package go_rpc

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

type HelloRequest struct {
	Req                  string   `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac0911f1623afdf3, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetReq() string {
	if m != nil {
		return m.Req
	}
	return ""
}

type HelloRes struct {
	Resp                 string   `protobuf:"bytes,1,opt,name=resp,proto3" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRes) Reset()         { *m = HelloRes{} }
func (m *HelloRes) String() string { return proto.CompactTextString(m) }
func (*HelloRes) ProtoMessage()    {}
func (*HelloRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac0911f1623afdf3, []int{1}
}

func (m *HelloRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRes.Unmarshal(m, b)
}
func (m *HelloRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRes.Marshal(b, m, deterministic)
}
func (m *HelloRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRes.Merge(m, src)
}
func (m *HelloRes) XXX_Size() int {
	return xxx_messageInfo_HelloRes.Size(m)
}
func (m *HelloRes) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRes.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRes proto.InternalMessageInfo

func (m *HelloRes) GetResp() string {
	if m != nil {
		return m.Resp
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "go_rpc.HelloRequest")
	proto.RegisterType((*HelloRes)(nil), "go_rpc.HelloRes")
}

func init() { proto.RegisterFile("go_rpc.proto", fileDescriptor_ac0911f1623afdf3) }

var fileDescriptor_ac0911f1623afdf3 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcf, 0x8f, 0x2f,
	0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x14, 0xb8, 0x78,
	0x3c, 0x52, 0x73, 0x72, 0xf2, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8, 0x98,
	0x8b, 0x52, 0x0b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x25, 0x39, 0x2e, 0x0e,
	0xa8, 0x8a, 0x62, 0x21, 0x21, 0x2e, 0x96, 0xa2, 0xd4, 0xe2, 0x02, 0xa8, 0x34, 0x98, 0x6d, 0x64,
	0xcf, 0xc5, 0xee, 0x5e, 0x94, 0x9a, 0x5a, 0x92, 0x5a, 0x24, 0x64, 0xc2, 0xc5, 0x05, 0x56, 0x1a,
	0x9e, 0x5f, 0x94, 0x93, 0x22, 0x24, 0xa2, 0x07, 0xb5, 0x11, 0xd9, 0x02, 0x29, 0x01, 0x34, 0xd1,
	0xe2, 0x24, 0x36, 0xb0, 0x8b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x67, 0x5e, 0x91, 0x62,
	0xa1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	HelloWorld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloRes, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) HelloWorld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloRes, error) {
	out := new(HelloRes)
	err := c.cc.Invoke(ctx, "/go_rpc.Greeter/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	HelloWorld(context.Context, *HelloRequest) (*HelloRes, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) HelloWorld(ctx context.Context, req *HelloRequest) (*HelloRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_rpc.Greeter/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).HelloWorld(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go_rpc.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _Greeter_HelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go_rpc.proto",
}
