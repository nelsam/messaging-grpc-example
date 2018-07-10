// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package messaging

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ident is a message used to identify a user.  It should be used when a user needs
// to log in or create an account.
type Ident struct {
	// username is the unique name of the user.
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	// password is the user's password.
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *Ident) Reset()                    { *m = Ident{} }
func (m *Ident) String() string            { return proto.CompactTextString(m) }
func (*Ident) ProtoMessage()               {}
func (*Ident) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Ident) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Ident) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// Profile is the set of public details about a user.
type Profile struct {
	// id is the user's unique ID.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// email is the user's email address.
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	// display_name is the user's name to show to other users.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Profile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Profile) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Profile) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func init() {
	proto.RegisterType((*Ident)(nil), "messaging.Ident")
	proto.RegisterType((*Profile)(nil), "messaging.Profile")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0xcf, 0xcc,
	0x4b, 0x57, 0xb2, 0xe7, 0x62, 0xf5, 0x4c, 0x49, 0xcd, 0x2b, 0x11, 0x92, 0xe2, 0xe2, 0x00, 0xa9,
	0xc8, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x41, 0x72, 0x05,
	0x89, 0xc5, 0xc5, 0xe5, 0xf9, 0x45, 0x29, 0x12, 0x4c, 0x10, 0x39, 0x18, 0x5f, 0x29, 0x88, 0x8b,
	0x3d, 0xa0, 0x28, 0x3f, 0x2d, 0x33, 0x27, 0x55, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x05, 0xaa, 0x99,
	0x29, 0x33, 0x45, 0x48, 0x84, 0x8b, 0x35, 0x35, 0x37, 0x31, 0x33, 0x07, 0xaa, 0x07, 0xc2, 0x11,
	0x52, 0xe4, 0xe2, 0x49, 0xc9, 0x2c, 0x2e, 0xc8, 0x49, 0xac, 0x8c, 0x07, 0x5b, 0xc6, 0x0c, 0x96,
	0xe4, 0x86, 0x8a, 0xf9, 0x25, 0xe6, 0xa6, 0x26, 0xb1, 0x81, 0x9d, 0x69, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x43, 0xa5, 0x38, 0xf1, 0xb4, 0x00, 0x00, 0x00,
}