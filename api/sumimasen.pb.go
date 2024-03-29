// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/sumimasen.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type SumimasenRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumimasenRequest) Reset()         { *m = SumimasenRequest{} }
func (m *SumimasenRequest) String() string { return proto.CompactTextString(m) }
func (*SumimasenRequest) ProtoMessage()    {}
func (*SumimasenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d32e0f8866184b25, []int{0}
}

func (m *SumimasenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumimasenRequest.Unmarshal(m, b)
}
func (m *SumimasenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumimasenRequest.Marshal(b, m, deterministic)
}
func (m *SumimasenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumimasenRequest.Merge(m, src)
}
func (m *SumimasenRequest) XXX_Size() int {
	return xxx_messageInfo_SumimasenRequest.Size(m)
}
func (m *SumimasenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumimasenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumimasenRequest proto.InternalMessageInfo

type SumimasenResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumimasenResponse) Reset()         { *m = SumimasenResponse{} }
func (m *SumimasenResponse) String() string { return proto.CompactTextString(m) }
func (*SumimasenResponse) ProtoMessage()    {}
func (*SumimasenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d32e0f8866184b25, []int{1}
}

func (m *SumimasenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumimasenResponse.Unmarshal(m, b)
}
func (m *SumimasenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumimasenResponse.Marshal(b, m, deterministic)
}
func (m *SumimasenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumimasenResponse.Merge(m, src)
}
func (m *SumimasenResponse) XXX_Size() int {
	return xxx_messageInfo_SumimasenResponse.Size(m)
}
func (m *SumimasenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumimasenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumimasenResponse proto.InternalMessageInfo

func (m *SumimasenResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SumimasenRequest)(nil), "api.SumimasenRequest")
	proto.RegisterType((*SumimasenResponse)(nil), "api.SumimasenResponse")
}

func init() { proto.RegisterFile("api/sumimasen.proto", fileDescriptor_d32e0f8866184b25) }

var fileDescriptor_d32e0f8866184b25 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2c, 0xc8, 0xd4,
	0x2f, 0x2e, 0xcd, 0xcd, 0xcc, 0x4d, 0x2c, 0x4e, 0xcd, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x12, 0xe2, 0x12, 0x08, 0x86, 0x89, 0x07, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x28, 0xe9, 0x72, 0x09, 0x22, 0x89, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x49,
	0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0xc1, 0xb8, 0x46, 0x9e, 0x5c, 0x9c, 0x70, 0xe5, 0x42, 0x36, 0xc8, 0x1c, 0x51, 0xbd, 0xc4, 0x82,
	0x4c, 0x3d, 0x74, 0xf3, 0xa5, 0xc4, 0xd0, 0x85, 0x21, 0x56, 0x28, 0x31, 0x38, 0x45, 0x73, 0xc9,
	0x26, 0xe5, 0xe7, 0xa4, 0xc7, 0xa7, 0xa4, 0x96, 0xa5, 0xe6, 0xe4, 0x17, 0xa4, 0x16, 0x15, 0xeb,
	0x21, 0x5c, 0x9d, 0x58, 0x90, 0xe9, 0xc4, 0x07, 0xd7, 0x15, 0x00, 0xf2, 0x43, 0x00, 0x63, 0x94,
	0x46, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x48, 0xaf, 0x2e, 0x42,
	0x2f, 0xc2, 0xc7, 0xfa, 0x89, 0x05, 0x99, 0x49, 0x6c, 0x60, 0x6f, 0x1b, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x80, 0x0a, 0xbc, 0xac, 0x0d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SumimasenClient is the client API for Sumimasen service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumimasenClient interface {
	Sumimasen(ctx context.Context, in *SumimasenRequest, opts ...grpc.CallOption) (*SumimasenResponse, error)
}

type sumimasenClient struct {
	cc *grpc.ClientConn
}

func NewSumimasenClient(cc *grpc.ClientConn) SumimasenClient {
	return &sumimasenClient{cc}
}

func (c *sumimasenClient) Sumimasen(ctx context.Context, in *SumimasenRequest, opts ...grpc.CallOption) (*SumimasenResponse, error) {
	out := new(SumimasenResponse)
	err := c.cc.Invoke(ctx, "/api.Sumimasen/Sumimasen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SumimasenServer is the server API for Sumimasen service.
type SumimasenServer interface {
	Sumimasen(context.Context, *SumimasenRequest) (*SumimasenResponse, error)
}

func RegisterSumimasenServer(s *grpc.Server, srv SumimasenServer) {
	s.RegisterService(&_Sumimasen_serviceDesc, srv)
}

func _Sumimasen_Sumimasen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumimasenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumimasenServer).Sumimasen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Sumimasen/Sumimasen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumimasenServer).Sumimasen(ctx, req.(*SumimasenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sumimasen_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Sumimasen",
	HandlerType: (*SumimasenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sumimasen",
			Handler:    _Sumimasen_Sumimasen_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/sumimasen.proto",
}
