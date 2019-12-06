// Code generated by protoc-gen-go. DO NOT EDIT.
// source: packages/gossip/transport/proto/handshake.proto

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

type HandshakeRequest struct {
	// protocol version number
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// string form of the sender address (e.g. "192.0.2.1:25", "[2001:db8::1]:80")
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	// string form of the recipient address
	To string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	// unix time
	Timestamp            int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandshakeRequest) Reset()         { *m = HandshakeRequest{} }
func (m *HandshakeRequest) String() string { return proto.CompactTextString(m) }
func (*HandshakeRequest) ProtoMessage()    {}
func (*HandshakeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9e96b60881ea276, []int{0}
}

func (m *HandshakeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandshakeRequest.Unmarshal(m, b)
}
func (m *HandshakeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandshakeRequest.Marshal(b, m, deterministic)
}
func (m *HandshakeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeRequest.Merge(m, src)
}
func (m *HandshakeRequest) XXX_Size() int {
	return xxx_messageInfo_HandshakeRequest.Size(m)
}
func (m *HandshakeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeRequest proto.InternalMessageInfo

func (m *HandshakeRequest) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HandshakeRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *HandshakeRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *HandshakeRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type HandshakeResponse struct {
	// hash of the ping packet
	ReqHash              []byte   `protobuf:"bytes,1,opt,name=req_hash,json=reqHash,proto3" json:"req_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandshakeResponse) Reset()         { *m = HandshakeResponse{} }
func (m *HandshakeResponse) String() string { return proto.CompactTextString(m) }
func (*HandshakeResponse) ProtoMessage()    {}
func (*HandshakeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9e96b60881ea276, []int{1}
}

func (m *HandshakeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandshakeResponse.Unmarshal(m, b)
}
func (m *HandshakeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandshakeResponse.Marshal(b, m, deterministic)
}
func (m *HandshakeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeResponse.Merge(m, src)
}
func (m *HandshakeResponse) XXX_Size() int {
	return xxx_messageInfo_HandshakeResponse.Size(m)
}
func (m *HandshakeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeResponse proto.InternalMessageInfo

func (m *HandshakeResponse) GetReqHash() []byte {
	if m != nil {
		return m.ReqHash
	}
	return nil
}

func init() {
	proto.RegisterType((*HandshakeRequest)(nil), "proto.HandshakeRequest")
	proto.RegisterType((*HandshakeResponse)(nil), "proto.HandshakeResponse")
}

func init() {
	proto.RegisterFile("packages/gossip/transport/proto/handshake.proto", fileDescriptor_f9e96b60881ea276)
}

var fileDescriptor_f9e96b60881ea276 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0xb4, 0x50, 0x6a, 0x01, 0x02, 0x4f, 0x41, 0x62, 0x88, 0x3a, 0x65, 0x8a, 0x07,
	0x7e, 0x00, 0x82, 0xa9, 0xb3, 0x47, 0x16, 0x74, 0x6d, 0x8f, 0xd8, 0x2a, 0xf6, 0x39, 0x77, 0x57,
	0x7e, 0x3f, 0xc2, 0x52, 0x05, 0x1b, 0xd3, 0xbb, 0xf7, 0x6e, 0xf8, 0xf4, 0x19, 0x57, 0x60, 0x7f,
	0x84, 0x09, 0xc5, 0x4d, 0x24, 0x12, 0x8b, 0x53, 0x86, 0x2c, 0x85, 0x58, 0x5d, 0x61, 0x52, 0x72,
	0x01, 0xf2, 0x41, 0x02, 0x1c, 0x71, 0xac, 0xdd, 0x5e, 0xd4, 0xd8, 0x64, 0x73, 0xb7, 0x3d, 0x7f,
	0x3c, 0xce, 0x27, 0x14, 0xb5, 0x9d, 0x59, 0x7d, 0x21, 0x4b, 0xa4, 0xdc, 0x35, 0x7d, 0x33, 0xdc,
	0xf8, 0x73, 0xb5, 0xd6, 0x2c, 0x3f, 0x98, 0x52, 0xd7, 0xf6, 0xcd, 0xb0, 0xf6, 0xf5, 0xb6, 0xb7,
	0xa6, 0x55, 0xea, 0x16, 0x75, 0x69, 0x95, 0xec, 0xa3, 0x59, 0x6b, 0x4c, 0x28, 0x0a, 0xa9, 0x74,
	0xcb, 0xbe, 0x19, 0x16, 0xfe, 0x77, 0xd8, 0x8c, 0xe6, 0xfe, 0x0f, 0x4f, 0x0a, 0x65, 0x41, 0xfb,
	0x60, 0xae, 0x18, 0xe7, 0xf7, 0x00, 0x12, 0x2a, 0xf1, 0xda, 0xaf, 0x18, 0xe7, 0x2d, 0x48, 0x78,
	0x7d, 0x79, 0x7b, 0x9e, 0xa2, 0x86, 0xd3, 0x6e, 0xdc, 0x53, 0x72, 0x91, 0x14, 0x3e, 0xf1, 0x30,
	0x21, 0xff, 0x68, 0x86, 0x98, 0x12, 0xf2, 0x7f, 0xe6, 0xbb, 0xcb, 0x1a, 0x4f, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x2a, 0x53, 0xc3, 0xbb, 0x23, 0x01, 0x00, 0x00,
}