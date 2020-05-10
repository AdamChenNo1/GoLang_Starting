// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package RestService

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type StringMessage struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringMessage) Reset()         { *m = StringMessage{} }
func (m *StringMessage) String() string { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()    {}
func (*StringMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *StringMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMessage.Unmarshal(m, b)
}
func (m *StringMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMessage.Marshal(b, m, deterministic)
}
func (m *StringMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMessage.Merge(m, src)
}
func (m *StringMessage) XXX_Size() int {
	return xxx_messageInfo_StringMessage.Size(m)
}
func (m *StringMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StringMessage proto.InternalMessageInfo

func (m *StringMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StringMessage)(nil), "RestService.StringMessage")
}

func init() {
	proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626)
}

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x0e, 0x4a, 0x2d, 0x2e, 0x09,
	0x86, 0x08, 0x49, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27,
	0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0x94, 0x2a, 0xa9, 0x72, 0xf1,
	0x06, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0xfb, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a, 0x89, 0x70,
	0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x46,
	0x5b, 0x18, 0xb9, 0x90, 0x0d, 0x15, 0x0a, 0xe6, 0x62, 0x76, 0x4f, 0x2d, 0x11, 0x92, 0xd2, 0x43,
	0x12, 0xd4, 0x43, 0x31, 0x48, 0x0a, 0x8f, 0x9c, 0x92, 0x48, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0xf8,
	0x84, 0x78, 0xf4, 0xd3, 0x53, 0x4b, 0xf4, 0xab, 0xc1, 0x76, 0xd4, 0x0a, 0x05, 0x70, 0xb1, 0x04,
	0xe4, 0x17, 0x93, 0x6f, 0xaa, 0x00, 0xd8, 0x54, 0x2e, 0x25, 0x56, 0xfd, 0x82, 0xfc, 0xe2, 0x12,
	0x2b, 0x46, 0xad, 0x24, 0x36, 0xb0, 0x27, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x85, 0x34,
	0x3d, 0xef, 0x20, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RestServiceClient is the client API for RestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RestServiceClient interface {
	Get(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
	Post(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type restServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRestServiceClient(cc grpc.ClientConnInterface) RestServiceClient {
	return &restServiceClient{cc}
}

func (c *restServiceClient) Get(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/RestService.RestService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restServiceClient) Post(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/RestService.RestService/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestServiceServer is the server API for RestService service.
type RestServiceServer interface {
	Get(context.Context, *StringMessage) (*StringMessage, error)
	Post(context.Context, *StringMessage) (*StringMessage, error)
}

// UnimplementedRestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRestServiceServer struct {
}

func (*UnimplementedRestServiceServer) Get(ctx context.Context, req *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRestServiceServer) Post(ctx context.Context, req *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}

func RegisterRestServiceServer(s *grpc.Server, srv RestServiceServer) {
	s.RegisterService(&_RestService_serviceDesc, srv)
}

func _RestService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RestService.RestService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestServiceServer).Get(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestService_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestServiceServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RestService.RestService/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestServiceServer).Post(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _RestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RestService.RestService",
	HandlerType: (*RestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RestService_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _RestService_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}