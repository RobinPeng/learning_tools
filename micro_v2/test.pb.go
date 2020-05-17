// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test_agent

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ReqMsg struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqMsg) Reset()         { *m = ReqMsg{} }
func (m *ReqMsg) String() string { return proto.CompactTextString(m) }
func (*ReqMsg) ProtoMessage()    {}
func (*ReqMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *ReqMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqMsg.Unmarshal(m, b)
}
func (m *ReqMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqMsg.Marshal(b, m, deterministic)
}
func (m *ReqMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqMsg.Merge(m, src)
}
func (m *ReqMsg) XXX_Size() int {
	return xxx_messageInfo_ReqMsg.Size(m)
}
func (m *ReqMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ReqMsg proto.InternalMessageInfo

func (m *ReqMsg) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type ResMsg struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResMsg) Reset()         { *m = ResMsg{} }
func (m *ResMsg) String() string { return proto.CompactTextString(m) }
func (*ResMsg) ProtoMessage()    {}
func (*ResMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{2}
}

func (m *ResMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResMsg.Unmarshal(m, b)
}
func (m *ResMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResMsg.Marshal(b, m, deterministic)
}
func (m *ResMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResMsg.Merge(m, src)
}
func (m *ResMsg) XXX_Size() int {
	return xxx_messageInfo_ResMsg.Size(m)
}
func (m *ResMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ResMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ResMsg proto.InternalMessageInfo

func (m *ResMsg) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ResMsg) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "test_agent.Error")
	proto.RegisterType((*ReqMsg)(nil), "test_agent.ReqMsg")
	proto.RegisterType((*ResMsg)(nil), "test_agent.ResMsg")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x31, 0xaf, 0x82, 0x30,
	0x14, 0x85, 0x1f, 0x4f, 0x20, 0x72, 0x59, 0xf4, 0x4e, 0x44, 0x17, 0xd2, 0xc4, 0xc8, 0x22, 0x03,
	0x0e, 0x8e, 0x4e, 0x0c, 0x0e, 0x3a, 0x34, 0x3a, 0x13, 0xc4, 0x0b, 0x71, 0xa0, 0xc5, 0xb6, 0xfe,
	0x7f, 0xd3, 0x62, 0xa2, 0x71, 0xfb, 0x72, 0x7a, 0x7a, 0xbe, 0x16, 0xc0, 0x90, 0x36, 0xf9, 0xa0,
	0xa4, 0x91, 0xe8, 0xb8, 0xaa, 0x3b, 0x12, 0x86, 0x6d, 0x20, 0x28, 0x95, 0x92, 0x0a, 0x11, 0xfc,
	0x46, 0xde, 0x28, 0xf1, 0x52, 0x2f, 0x0b, 0xb8, 0x63, 0x9c, 0xc1, 0xa4, 0xd7, 0x5d, 0xf2, 0x9f,
	0x7a, 0x59, 0xc4, 0x2d, 0xb2, 0x15, 0x84, 0x9c, 0x1e, 0x47, 0xdd, 0xe1, 0x12, 0xa2, 0xa7, 0x26,
	0x55, 0x89, 0xba, 0x1f, 0x2f, 0x45, 0x7c, 0x6a, 0x83, 0x53, 0xdd, 0x13, 0x2b, 0x6d, 0x4d, 0xdb,
	0xda, 0x1a, 0x02, 0xb2, 0xfb, 0xae, 0x12, 0x17, 0xf3, 0xfc, 0xe3, 0xce, 0x9d, 0x98, 0x8f, 0xe7,
	0xd6, 0x7f, 0x17, 0xad, 0x7c, 0xcb, 0x1c, 0x17, 0x7b, 0xf0, 0xcf, 0xa4, 0x0d, 0xee, 0x20, 0xe6,
	0x43, 0x73, 0xd1, 0xa4, 0x0e, 0xa2, 0x95, 0x88, 0xdf, 0x23, 0xe3, 0x73, 0x16, 0x3f, 0x99, 0x75,
	0xb3, 0xbf, 0x6b, 0xe8, 0x3e, 0xbc, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x4f, 0x13, 0x89,
	0xfe, 0x00, 0x00, 0x00,
}