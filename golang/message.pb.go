// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package messaging

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Message is a message from one user to another
type Message struct {
	// id is the message's unique ID.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// from is the user the message was sent from.
	From *Profile `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	// to is the user the message was sent to.
	To *Profile `protobuf:"bytes,3,opt,name=to" json:"to,omitempty"`
	// body is the message's body.
	Body string `protobuf:"bytes,5,opt,name=body" json:"body,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetFrom() *Profile {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Message) GetTo() *Profile {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Message) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// Thread is a sequence of messages from one user to another
type Thread struct {
	// id is the thread's unique ID.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// title is the message's title.
	Title string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	// messages is the list of messages in the thread.
	Messages []*Message `protobuf:"bytes,2,rep,name=messages" json:"messages,omitempty"`
}

func (m *Thread) Reset()                    { *m = Thread{} }
func (m *Thread) String() string            { return proto.CompactTextString(m) }
func (*Thread) ProtoMessage()               {}
func (*Thread) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Thread) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Thread) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Thread) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "messaging.Message")
	proto.RegisterType((*Thread)(nil), "messaging.Thread")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x70, 0x33, 0xf3, 0xd2,
	0xa5, 0xb8, 0x4a, 0x8b, 0x53, 0x8b, 0x20, 0xc2, 0x4a, 0xf5, 0x5c, 0xec, 0xbe, 0x10, 0x75, 0x42,
	0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x4c, 0x99, 0x29, 0x42,
	0x6a, 0x5c, 0x2c, 0x69, 0x45, 0xf9, 0xb9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x42, 0x7a,
	0x70, 0x03, 0xf4, 0x02, 0x8a, 0xf2, 0xd3, 0x32, 0x73, 0x52, 0x83, 0xc0, 0xf2, 0x42, 0x4a, 0x5c,
	0x4c, 0x25, 0xf9, 0x12, 0xcc, 0x38, 0x55, 0x31, 0x95, 0xe4, 0x0b, 0x09, 0x71, 0xb1, 0x24, 0xe5,
	0xa7, 0x54, 0x4a, 0xb0, 0x82, 0x4d, 0x07, 0xb3, 0xbd, 0x58, 0x38, 0x58, 0x04, 0x58, 0x95, 0xe2,
	0xb8, 0xd8, 0x42, 0x32, 0x8a, 0x52, 0x13, 0x53, 0x30, 0xec, 0x17, 0xe1, 0x62, 0x2d, 0xc9, 0x2c,
	0xc9, 0x49, 0x05, 0x1b, 0xcd, 0x19, 0x04, 0xe1, 0x08, 0xe9, 0x71, 0x71, 0x40, 0x3d, 0x56, 0x2c,
	0xc1, 0xa4, 0xc0, 0x8c, 0x66, 0x27, 0xd4, 0x2f, 0x41, 0x70, 0x35, 0x49, 0x6c, 0x60, 0x7f, 0x1a,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xc8, 0xac, 0x4e, 0x0f, 0x01, 0x00, 0x00,
}
