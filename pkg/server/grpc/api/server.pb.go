// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package grpc

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

type Session struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Session) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type User struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Auth                 *Session `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetAuth() *Session {
	if m != nil {
		return m.Auth
	}
	return nil
}

type Response struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Session              *Session `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
	Ekadashi             int64    `protobuf:"varint,3,opt,name=Ekadashi,json=ekadashi,proto3" json:"Ekadashi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func (m *Response) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *Response) GetEkadashi() int64 {
	if m != nil {
		return m.Ekadashi
	}
	return 0
}

func init() {
	proto.RegisterType((*Session)(nil), "grpc.Session")
	proto.RegisterType((*User)(nil), "grpc.User")
	proto.RegisterType((*Response)(nil), "grpc.Response")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbd, 0x4e, 0xc3, 0x40,
	0x10, 0x84, 0x31, 0x31, 0xc4, 0x6c, 0x02, 0xc5, 0x8a, 0xc2, 0x4a, 0x15, 0xdc, 0x10, 0x09, 0xe4,
	0xc2, 0x79, 0x06, 0x24, 0x0a, 0xaa, 0x8b, 0x28, 0x28, 0x0f, 0xbc, 0xb2, 0x2d, 0xc3, 0x9d, 0x75,
	0x7b, 0xc1, 0xcf, 0xc2, 0xdb, 0xa2, 0xfb, 0xb1, 0x29, 0x90, 0xa0, 0x9b, 0xf1, 0xee, 0x37, 0xb3,
	0xd6, 0xc1, 0x9a, 0xc9, 0x7c, 0x92, 0x29, 0x07, 0xa3, 0xad, 0xc6, 0xb4, 0x31, 0xc3, 0x5b, 0xb1,
	0x87, 0xe5, 0x81, 0x98, 0x3b, 0xad, 0x10, 0x21, 0x55, 0xf2, 0x83, 0xf2, 0x64, 0x9b, 0xec, 0x2e,
	0x84, 0xd7, 0x78, 0x0d, 0x67, 0x56, 0xf7, 0xa4, 0xf2, 0x53, 0xff, 0x31, 0x98, 0xe2, 0x05, 0xd2,
	0x67, 0x26, 0xe3, 0x88, 0x23, 0x93, 0x99, 0x08, 0xa7, 0x71, 0x03, 0xd9, 0x20, 0x99, 0x47, 0x6d,
	0xea, 0x08, 0xcd, 0x1e, 0x6f, 0x20, 0x95, 0x47, 0xdb, 0xe6, 0x8b, 0x6d, 0xb2, 0x5b, 0x55, 0x97,
	0xa5, 0xbb, 0xa0, 0x8c, 0xf5, 0xc2, 0x8f, 0x8a, 0x1e, 0x32, 0x41, 0x3c, 0x68, 0xc5, 0xe4, 0xa2,
	0x4c, 0xd4, 0xb1, 0x62, 0xf6, 0x78, 0x0b, 0x4b, 0x0e, 0xa0, 0x6f, 0xf9, 0x95, 0x36, 0x4d, 0x5d,
	0xc8, 0x43, 0x2f, 0x6b, 0xc9, 0x6d, 0xe7, 0x7b, 0x17, 0x22, 0xa3, 0xe8, 0xab, 0xaf, 0xe4, 0x67,
	0x88, 0x15, 0xe0, 0xa3, 0x54, 0xf5, 0x3b, 0x09, 0x6a, 0x3a, 0xb6, 0x46, 0x5a, 0x87, 0x43, 0x88,
	0x75, 0xbf, 0xbb, 0xb9, 0x0a, 0x7a, 0xba, 0xaf, 0x38, 0xc1, 0x3b, 0x58, 0x05, 0xe6, 0x49, 0x37,
	0xdd, 0x7f, 0xcb, 0xf7, 0xb0, 0x3e, 0xb4, 0x7a, 0x9c, 0x0b, 0xff, 0xdc, 0x7e, 0x3d, 0xf7, 0xaf,
	0xb4, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x37, 0x87, 0xb3, 0xb5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EkadashiClient is the client API for Ekadashi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EkadashiClient interface {
	HandleRegistration(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	HandleLogin(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	ShowEkadashi(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
}

type ekadashiClient struct {
	cc *grpc.ClientConn
}

func NewEkadashiClient(cc *grpc.ClientConn) EkadashiClient {
	return &ekadashiClient{cc}
}

func (c *ekadashiClient) HandleRegistration(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.Ekadashi/HandleRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekadashiClient) HandleLogin(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.Ekadashi/HandleLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekadashiClient) ShowEkadashi(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.Ekadashi/ShowEkadashi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EkadashiServer is the server API for Ekadashi service.
type EkadashiServer interface {
	HandleRegistration(context.Context, *User) (*Response, error)
	HandleLogin(context.Context, *User) (*Response, error)
	ShowEkadashi(context.Context, *User) (*Response, error)
}

func RegisterEkadashiServer(s *grpc.Server, srv EkadashiServer) {
	s.RegisterService(&_Ekadashi_serviceDesc, srv)
}

func _Ekadashi_HandleRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkadashiServer).HandleRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Ekadashi/HandleRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkadashiServer).HandleRegistration(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekadashi_HandleLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkadashiServer).HandleLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Ekadashi/HandleLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkadashiServer).HandleLogin(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekadashi_ShowEkadashi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkadashiServer).ShowEkadashi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Ekadashi/ShowEkadashi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkadashiServer).ShowEkadashi(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ekadashi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Ekadashi",
	HandlerType: (*EkadashiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleRegistration",
			Handler:    _Ekadashi_HandleRegistration_Handler,
		},
		{
			MethodName: "HandleLogin",
			Handler:    _Ekadashi_HandleLogin_Handler,
		},
		{
			MethodName: "ShowEkadashi",
			Handler:    _Ekadashi_ShowEkadashi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}