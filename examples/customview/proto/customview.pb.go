// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/gomatcha/matcha/examples/customview/proto/customview.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type View struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Enabled              bool     `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *View) Reset()         { *m = View{} }
func (m *View) String() string { return proto.CompactTextString(m) }
func (*View) ProtoMessage()    {}
func (*View) Descriptor() ([]byte, []int) {
	return fileDescriptor_b938b504d6f3b152, []int{0}
}
func (m *View) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_View.Unmarshal(m, b)
}
func (m *View) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_View.Marshal(b, m, deterministic)
}
func (m *View) XXX_Merge(src proto.Message) {
	xxx_messageInfo_View.Merge(m, src)
}
func (m *View) XXX_Size() int {
	return xxx_messageInfo_View.Size(m)
}
func (m *View) XXX_DiscardUnknown() {
	xxx_messageInfo_View.DiscardUnknown(m)
}

var xxx_messageInfo_View proto.InternalMessageInfo

func (m *View) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func (m *View) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type Event struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_b938b504d6f3b152, []int{1}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto.RegisterType((*View)(nil), "matcha.examples.customview.View")
	proto.RegisterType((*Event)(nil), "matcha.examples.customview.Event")
}

func init() {
	proto.RegisterFile("github.com/gomatcha/matcha/examples/customview/proto/customview.proto", fileDescriptor_b938b504d6f3b152)
}

var fileDescriptor_b938b504d6f3b152 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0xd4,
	0x87, 0x52, 0xa9, 0x15, 0x89, 0xb9, 0x05, 0x39, 0xa9, 0xc5, 0xfa, 0xc9, 0xa5, 0xc5, 0x25, 0xf9,
	0xb9, 0x65, 0x99, 0xa9, 0xe5, 0xfa, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x48, 0x02, 0x7a, 0x60, 0x01,
	0x21, 0x29, 0x88, 0x16, 0x3d, 0x98, 0x16, 0x3d, 0x84, 0x0a, 0x25, 0x33, 0x2e, 0x96, 0xb0, 0xcc,
	0xd4, 0x72, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0x8e, 0x20, 0x08, 0x47, 0x48, 0x82, 0x8b, 0x3d, 0x35, 0x2f, 0x31, 0x29, 0x27, 0x35, 0x45, 0x82,
	0x09, 0x2c, 0x0e, 0xe3, 0x2a, 0xc9, 0x72, 0xb1, 0xba, 0x96, 0xa5, 0xe6, 0x95, 0x60, 0xd7, 0xe8,
	0xe4, 0xc1, 0x25, 0x93, 0x99, 0xaf, 0x07, 0x73, 0xb3, 0x1e, 0xba, 0x93, 0x9c, 0xf8, 0x9d, 0xc1,
	0x22, 0x20, 0xab, 0x03, 0x40, 0x02, 0x51, 0xac, 0x60, 0xf1, 0x45, 0x4c, 0xe8, 0x12, 0x49, 0x6c,
	0x60, 0x09, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0xd7, 0x29, 0x58, 0x0c, 0x01, 0x00,
	0x00,
}