// Code generated by protoc-gen-go. DO NOT EDIT.
// source: record.proto

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

type FetchRecordRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchRecordRequest) Reset()         { *m = FetchRecordRequest{} }
func (m *FetchRecordRequest) String() string { return proto.CompactTextString(m) }
func (*FetchRecordRequest) ProtoMessage()    {}
func (*FetchRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf94fd919e302a1d, []int{0}
}

func (m *FetchRecordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchRecordRequest.Unmarshal(m, b)
}
func (m *FetchRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchRecordRequest.Marshal(b, m, deterministic)
}
func (m *FetchRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchRecordRequest.Merge(m, src)
}
func (m *FetchRecordRequest) XXX_Size() int {
	return xxx_messageInfo_FetchRecordRequest.Size(m)
}
func (m *FetchRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchRecordRequest proto.InternalMessageInfo

type FetchRecordResponse struct {
	Error                Error    `protobuf:"varint,1,opt,name=error,proto3,enum=common.Error" json:"error,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchRecordResponse) Reset()         { *m = FetchRecordResponse{} }
func (m *FetchRecordResponse) String() string { return proto.CompactTextString(m) }
func (*FetchRecordResponse) ProtoMessage()    {}
func (*FetchRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf94fd919e302a1d, []int{1}
}

func (m *FetchRecordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchRecordResponse.Unmarshal(m, b)
}
func (m *FetchRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchRecordResponse.Marshal(b, m, deterministic)
}
func (m *FetchRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchRecordResponse.Merge(m, src)
}
func (m *FetchRecordResponse) XXX_Size() int {
	return xxx_messageInfo_FetchRecordResponse.Size(m)
}
func (m *FetchRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchRecordResponse proto.InternalMessageInfo

func (m *FetchRecordResponse) GetError() Error {
	if m != nil {
		return m.Error
	}
	return Error_NO_ERROR
}

func (m *FetchRecordResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *FetchRecordResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*FetchRecordRequest)(nil), "record.FetchRecordRequest")
	proto.RegisterType((*FetchRecordResponse)(nil), "record.FetchRecordResponse")
}

func init() { proto.RegisterFile("record.proto", fileDescriptor_bf94fd919e302a1d) }

var fileDescriptor_bf94fd919e302a1d = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4a, 0x4d, 0xce,
	0x2f, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0xa4, 0x78, 0x92, 0xf3,
	0x73, 0x73, 0xf3, 0xf3, 0x20, 0xa2, 0x4a, 0x22, 0x5c, 0x42, 0x6e, 0xa9, 0x25, 0xc9, 0x19, 0x41,
	0x60, 0xc9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0xa5, 0x04, 0x2e, 0x61, 0x14, 0xd1, 0xe2,
	0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x65, 0x2e, 0xd6, 0xd4, 0xa2, 0xa2, 0xfc, 0x22, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0x3e, 0x23, 0x5e, 0x3d, 0xa8, 0x51, 0xae, 0x20, 0xc1, 0x20, 0x88, 0x9c, 0x90,
	0x10, 0x17, 0x4b, 0x72, 0x7e, 0x4a, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d,
	0x24, 0xc0, 0xc5, 0x5c, 0x5a, 0x94, 0x23, 0xc1, 0x0c, 0x16, 0x02, 0x31, 0x9d, 0x14, 0xa3, 0xe4,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0x40, 0x66, 0xe8, 0x97, 0x14, 0x25, 0xa6, 0x64, 0xe6, 0xa5,
	0x3b, 0x7a, 0xea, 0x97, 0x94, 0xa7, 0x26, 0xe9, 0x83, 0x9d, 0x96, 0xc4, 0x06, 0xa6, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x67, 0x46, 0x1a, 0x0e, 0xc7, 0x00, 0x00, 0x00,
}