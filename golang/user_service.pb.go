// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_service.proto

package messaging

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// AuthedProfile is a public user profile with authorization attached.
type AuthedProfile struct {
	Auth    *Ident   `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	Profile *Profile `protobuf:"bytes,2,opt,name=profile" json:"profile,omitempty"`
}

func (m *AuthedProfile) Reset()                    { *m = AuthedProfile{} }
func (m *AuthedProfile) String() string            { return proto.CompactTextString(m) }
func (*AuthedProfile) ProtoMessage()               {}
func (*AuthedProfile) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *AuthedProfile) GetAuth() *Ident {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *AuthedProfile) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthedProfile)(nil), "messaging.AuthedProfile")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	// CreateUser creates a user.
	CreateUser(ctx context.Context, in *Ident, opts ...grpc.CallOption) (*Profile, error)
	// SetProfile updates a user's profile.
	SetProfile(ctx context.Context, in *AuthedProfile, opts ...grpc.CallOption) (*Profile, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *Ident, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := grpc.Invoke(ctx, "/messaging.UserService/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetProfile(ctx context.Context, in *AuthedProfile, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := grpc.Invoke(ctx, "/messaging.UserService/SetProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	// CreateUser creates a user.
	CreateUser(context.Context, *Ident) (*Profile, error)
	// SetProfile updates a user's profile.
	SetProfile(context.Context, *AuthedProfile) (*Profile, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ident)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*Ident))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthedProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging.UserService/SetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetProfile(ctx, req.(*AuthedProfile))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messaging.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "SetProfile",
			Handler:    _UserService_SetProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}

func init() { proto.RegisterFile("user_service.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x2d, 0x4e, 0x2d,
	0x8a, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0xcc, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0xcf, 0xcc, 0x4b, 0x97, 0xe2, 0x02, 0x49, 0x43, 0x84, 0x95,
	0x92, 0xb9, 0x78, 0x1d, 0x4b, 0x4b, 0x32, 0x52, 0x53, 0x02, 0x8a, 0xf2, 0xd3, 0x32, 0x73, 0x52,
	0x85, 0x54, 0xb8, 0x58, 0x12, 0x4b, 0x4b, 0x32, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x04,
	0xf4, 0xe0, 0xda, 0xf4, 0x3c, 0x53, 0x52, 0xf3, 0x4a, 0x82, 0xc0, 0xb2, 0x42, 0x3a, 0x5c, 0xec,
	0x05, 0x10, 0x0d, 0x12, 0x4c, 0x60, 0x85, 0x42, 0x48, 0x0a, 0xa1, 0x46, 0x05, 0xc1, 0x94, 0x18,
	0x35, 0x32, 0x72, 0x71, 0x87, 0x16, 0xa7, 0x16, 0x05, 0x43, 0x5c, 0x24, 0x64, 0xc2, 0xc5, 0xe5,
	0x5c, 0x94, 0x9a, 0x58, 0x92, 0x0a, 0x12, 0x14, 0xc2, 0xb0, 0x43, 0x0a, 0x8b, 0x61, 0x4a, 0x0c,
	0x42, 0x36, 0x5c, 0x5c, 0xc1, 0xa9, 0x25, 0x30, 0x77, 0x4a, 0x20, 0xa9, 0x41, 0xf1, 0x01, 0x76,
	0xdd, 0x49, 0x6c, 0x60, 0xff, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x36, 0x4f, 0xee, 0x56,
	0x1c, 0x01, 0x00, 0x00,
}