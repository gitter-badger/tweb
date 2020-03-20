// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session.proto

package proto

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

type Session struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{0}
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

func (m *Session) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{1}
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

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Error                Error    `protobuf:"varint,1,opt,name=error,proto3,enum=common.Error" json:"error,omitempty"`
	Session              *Session `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{2}
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

func (m *LoginResponse) GetError() Error {
	if m != nil {
		return m.Error
	}
	return Error_NO_ERROR
}

func (m *LoginResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type LogoutRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{3}
}

func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (m *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(m, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

type LogoutResponse struct {
	Error                Error    `protobuf:"varint,1,opt,name=error,proto3,enum=common.Error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{4}
}

func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (m *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(m, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

func (m *LogoutResponse) GetError() Error {
	if m != nil {
		return m.Error
	}
	return Error_NO_ERROR
}

func init() {
	proto.RegisterType((*Session)(nil), "session.Session")
	proto.RegisterType((*LoginRequest)(nil), "session.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "session.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "session.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "session.LogoutResponse")
}

func init() { proto.RegisterFile("session.proto", fileDescriptor_3a6be1b361fa6f14) }

var fileDescriptor_3a6be1b361fa6f14 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xa9, 0xb8, 0xae, 0x3b, 0xb6, 0xab, 0x04, 0x0f, 0xcb, 0x5e, 0xac, 0xf1, 0xb2, 0x78,
	0x68, 0x61, 0xc5, 0x1f, 0xb0, 0x82, 0x82, 0xe0, 0x29, 0xe2, 0xc5, 0x93, 0xad, 0x3b, 0xd4, 0x22,
	0xcd, 0xd4, 0x4c, 0xca, 0xfe, 0x7d, 0x49, 0x9a, 0xf4, 0xbe, 0xa7, 0xf0, 0xde, 0x97, 0xbc, 0x99,
	0x17, 0xc8, 0x18, 0x99, 0x5b, 0xd2, 0x45, 0x6f, 0xc8, 0x92, 0x98, 0x07, 0xb9, 0x4e, 0xbf, 0xa9,
	0xeb, 0xa2, 0x2d, 0x77, 0x30, 0x7f, 0x1f, 0x81, 0xb8, 0x86, 0x99, 0xa5, 0x5f, 0xd4, 0xab, 0x24,
	0x4f, 0x36, 0x0b, 0x35, 0x0a, 0x91, 0xc3, 0xe9, 0xc0, 0x68, 0x56, 0x27, 0x79, 0xb2, 0xb9, 0xd8,
	0xa6, 0x45, 0x78, 0xfd, 0xc1, 0x68, 0x94, 0x27, 0xf2, 0x05, 0xd2, 0x37, 0x6a, 0x5a, 0xad, 0xf0,
	0x6f, 0x40, 0xb6, 0x62, 0x0d, 0xe7, 0xce, 0xd7, 0x55, 0x87, 0x21, 0x6a, 0xd2, 0x8e, 0xf5, 0x15,
	0xf3, 0x81, 0xcc, 0xde, 0x27, 0x2e, 0xd4, 0xa4, 0xe5, 0x17, 0x64, 0x21, 0x87, 0x7b, 0xd2, 0x8c,
	0xe2, 0x0e, 0x66, 0x68, 0x0c, 0x19, 0x9f, 0xb2, 0xdc, 0x66, 0x71, 0xf6, 0xb3, 0x33, 0xd5, 0xc8,
	0xc4, 0x3d, 0xc4, 0x66, 0x61, 0xc5, 0xab, 0x22, 0x16, 0x0f, 0xc5, 0x54, 0xbc, 0x20, 0x2f, 0xfd,
	0x04, 0x1a, 0x6c, 0x58, 0x55, 0x3e, 0xc2, 0x32, 0x1a, 0x47, 0xcc, 0x7c, 0xba, 0xfd, 0xbc, 0x69,
	0x5a, 0xfb, 0x33, 0xd4, 0x8e, 0x96, 0xd6, 0x54, 0xfb, 0x56, 0x37, 0xbb, 0xd7, 0xd2, 0x1e, 0xb0,
	0x2e, 0xfd, 0xbf, 0xd6, 0x67, 0xfe, 0x78, 0xf8, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x88, 0x45, 0x95,
	0xc1, 0x86, 0x01, 0x00, 0x00,
}
