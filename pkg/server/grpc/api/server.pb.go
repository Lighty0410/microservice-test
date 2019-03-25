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

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
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

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Session struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
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

func (m *Session) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RegisterRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type LoginRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{3}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type LoginResponse struct {
	Response             *Session `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{4}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetResponse() *Session {
	if m != nil {
		return m.Response
	}
	return nil
}

type ShowEkadashiRequest struct {
	Request              *Session `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowEkadashiRequest) Reset()         { *m = ShowEkadashiRequest{} }
func (m *ShowEkadashiRequest) String() string { return proto.CompactTextString(m) }
func (*ShowEkadashiRequest) ProtoMessage()    {}
func (*ShowEkadashiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{5}
}

func (m *ShowEkadashiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowEkadashiRequest.Unmarshal(m, b)
}
func (m *ShowEkadashiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowEkadashiRequest.Marshal(b, m, deterministic)
}
func (m *ShowEkadashiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowEkadashiRequest.Merge(m, src)
}
func (m *ShowEkadashiRequest) XXX_Size() int {
	return xxx_messageInfo_ShowEkadashiRequest.Size(m)
}
func (m *ShowEkadashiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowEkadashiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShowEkadashiRequest proto.InternalMessageInfo

func (m *ShowEkadashiRequest) GetRequest() *Session {
	if m != nil {
		return m.Request
	}
	return nil
}

type ShowEkadashiResponse struct {
	Ekadashi             int64    `protobuf:"varint,1,opt,name=ekadashi,proto3" json:"ekadashi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowEkadashiResponse) Reset()         { *m = ShowEkadashiResponse{} }
func (m *ShowEkadashiResponse) String() string { return proto.CompactTextString(m) }
func (*ShowEkadashiResponse) ProtoMessage()    {}
func (*ShowEkadashiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{6}
}

func (m *ShowEkadashiResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowEkadashiResponse.Unmarshal(m, b)
}
func (m *ShowEkadashiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowEkadashiResponse.Marshal(b, m, deterministic)
}
func (m *ShowEkadashiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowEkadashiResponse.Merge(m, src)
}
func (m *ShowEkadashiResponse) XXX_Size() int {
	return xxx_messageInfo_ShowEkadashiResponse.Size(m)
}
func (m *ShowEkadashiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowEkadashiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowEkadashiResponse proto.InternalMessageInfo

func (m *ShowEkadashiResponse) GetEkadashi() int64 {
	if m != nil {
		return m.Ekadashi
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{7}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "grpc.User")
	proto.RegisterType((*Session)(nil), "grpc.Session")
	proto.RegisterType((*RegisterRequest)(nil), "grpc.RegisterRequest")
	proto.RegisterType((*LoginRequest)(nil), "grpc.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "grpc.LoginResponse")
	proto.RegisterType((*ShowEkadashiRequest)(nil), "grpc.ShowEkadashiRequest")
	proto.RegisterType((*ShowEkadashiResponse)(nil), "grpc.ShowEkadashiResponse")
	proto.RegisterType((*Empty)(nil), "grpc.Empty")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xb1, 0x4e, 0xc3, 0x40,
	0x0c, 0x6d, 0x21, 0x25, 0xc1, 0x4d, 0x85, 0xe4, 0x16, 0xa9, 0xdc, 0x00, 0xe8, 0x16, 0x60, 0x89,
	0x20, 0x48, 0x0c, 0x0c, 0x6c, 0x15, 0x0b, 0xd3, 0x55, 0x7c, 0x40, 0xa0, 0x56, 0x1a, 0x55, 0xcd,
	0x85, 0x73, 0x4a, 0xc5, 0xa7, 0xf1, 0x77, 0xa8, 0x77, 0x97, 0xa8, 0x85, 0x0e, 0x6c, 0xf6, 0xb3,
	0x9f, 0xfd, 0xfc, 0x64, 0x88, 0x99, 0xcc, 0x27, 0x99, 0xa4, 0x32, 0xba, 0xd6, 0x18, 0xe4, 0xa6,
	0x7a, 0x97, 0x0f, 0x10, 0xbc, 0x32, 0x19, 0x44, 0x08, 0xca, 0x6c, 0x49, 0xe3, 0xee, 0x65, 0xf7,
	0xfa, 0x58, 0xd9, 0x18, 0x05, 0x44, 0x55, 0xc6, 0xbc, 0xd6, 0x66, 0x36, 0x3e, 0xb0, 0x78, 0x9b,
	0xcb, 0x0b, 0x08, 0xa7, 0xc4, 0x5c, 0xe8, 0x12, 0x47, 0xd0, 0xab, 0xf5, 0x82, 0x4a, 0xcf, 0x75,
	0x89, 0xbc, 0x83, 0x13, 0x45, 0x79, 0xc1, 0x35, 0x19, 0x45, 0x1f, 0x2b, 0xe2, 0x1a, 0xcf, 0x21,
	0x58, 0x31, 0x19, 0xdb, 0xd7, 0x4f, 0x21, 0xd9, 0x08, 0x48, 0x36, 0xdb, 0x95, 0xc5, 0x65, 0x02,
	0xf1, 0x8b, 0xce, 0x8b, 0xf2, 0xbf, 0xfd, 0x8f, 0x30, 0xf0, 0xfd, 0x5c, 0xe9, 0x92, 0x09, 0x6f,
	0x20, 0x32, 0x3e, 0xf6, 0xa4, 0x81, 0x23, 0x79, 0xa9, 0xaa, 0x2d, 0xcb, 0x27, 0x18, 0x4e, 0xe7,
	0x7a, 0x3d, 0x59, 0x64, 0xb3, 0x8c, 0xe7, 0x45, 0xb3, 0xf2, 0x0a, 0x42, 0xe3, 0xc2, 0xfd, 0x03,
	0x9a, 0xaa, 0x4c, 0x61, 0xb4, 0xcb, 0xf7, 0x12, 0x04, 0x44, 0xe4, 0x31, 0x3b, 0xe1, 0x50, 0xb5,
	0xb9, 0x0c, 0xa1, 0x37, 0x59, 0x56, 0xf5, 0x57, 0xfa, 0xdd, 0x85, 0xa8, 0x61, 0xe2, 0x2d, 0x44,
	0x8d, 0x51, 0x78, 0xea, 0xb6, 0xfd, 0x32, 0x4e, 0xf4, 0x1d, 0x6c, 0xc9, 0xb2, 0x83, 0x29, 0xf4,
	0xec, 0xdd, 0x88, 0x0e, 0xdf, 0x36, 0x4d, 0x0c, 0x77, 0x30, 0x7f, 0x6d, 0x07, 0x9f, 0x21, 0xde,
	0xd6, 0x8b, 0x67, 0xfe, 0xae, 0xbf, 0x1e, 0x08, 0xb1, 0xaf, 0xd4, 0x0c, 0x7a, 0x3b, 0xb2, 0xdf,
	0x73, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x80, 0xd9, 0x12, 0x4d, 0x02, 0x00, 0x00,
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
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Empty, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	ShowEkadashi(ctx context.Context, in *ShowEkadashiRequest, opts ...grpc.CallOption) (*ShowEkadashiResponse, error)
}

type ekadashiClient struct {
	cc *grpc.ClientConn
}

func NewEkadashiClient(cc *grpc.ClientConn) EkadashiClient {
	return &ekadashiClient{cc}
}

func (c *ekadashiClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/grpc.Ekadashi/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekadashiClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/grpc.Ekadashi/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekadashiClient) ShowEkadashi(ctx context.Context, in *ShowEkadashiRequest, opts ...grpc.CallOption) (*ShowEkadashiResponse, error) {
	out := new(ShowEkadashiResponse)
	err := c.cc.Invoke(ctx, "/grpc.Ekadashi/ShowEkadashi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EkadashiServer is the server API for Ekadashi service.
type EkadashiServer interface {
	Register(context.Context, *RegisterRequest) (*Empty, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	ShowEkadashi(context.Context, *ShowEkadashiRequest) (*ShowEkadashiResponse, error)
}

func RegisterEkadashiServer(s *grpc.Server, srv EkadashiServer) {
	s.RegisterService(&_Ekadashi_serviceDesc, srv)
}

func _Ekadashi_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkadashiServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Ekadashi/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkadashiServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekadashi_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkadashiServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Ekadashi/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkadashiServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekadashi_ShowEkadashi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowEkadashiRequest)
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
		return srv.(EkadashiServer).ShowEkadashi(ctx, req.(*ShowEkadashiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ekadashi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Ekadashi",
	HandlerType: (*EkadashiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Ekadashi_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Ekadashi_Login_Handler,
		},
		{
			MethodName: "ShowEkadashi",
			Handler:    _Ekadashi_ShowEkadashi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
