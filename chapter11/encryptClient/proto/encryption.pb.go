// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/encryption.proto

/*
Package encryption is a generated protocol buffer package.

It is generated from these files:
	proto/encryption.proto

It has these top-level messages:
	Request
	Response
*/
package encryption

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

type Request struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Key     string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Request) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Response struct {
	Result string `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("proto/encryption.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xcd, 0x4b, 0x2e, 0xaa, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0x0b, 0x28,
	0x99, 0x72, 0xb1, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x70, 0xb1, 0xe7, 0xa6,
	0x16, 0x17, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42, 0x02,
	0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x4c, 0x60, 0x51, 0x10, 0x53, 0x49, 0x89, 0x8b, 0x23, 0x28,
	0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8c, 0x8b, 0xad, 0x28, 0xb5, 0xb8, 0x34, 0xa7,
	0x04, 0xaa, 0x00, 0xca, 0x33, 0xf2, 0xe7, 0xe2, 0x74, 0x85, 0x58, 0x97, 0x5a, 0x24, 0xa4, 0xc0,
	0xc5, 0x0e, 0xe5, 0x08, 0x71, 0xe8, 0x41, 0x6d, 0x94, 0xe2, 0xd4, 0x83, 0x19, 0xa2, 0xc4, 0x00,
	0x52, 0xe1, 0x92, 0x8a, 0x4f, 0x45, 0x12, 0x1b, 0xd8, 0xc9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x73, 0x83, 0xda, 0xf4, 0xcc, 0x00, 0x00, 0x00,
}
