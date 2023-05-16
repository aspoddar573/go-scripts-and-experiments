// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common_messages.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common_messages.proto

It has these top-level messages:
	RequestContext
*/
package common

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

type RequestContext struct {
	TenantId string `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId" json:"tenant_id,omitempty"`
}

func (m *RequestContext) Reset()                    { *m = RequestContext{} }
func (m *RequestContext) String() string            { return proto.CompactTextString(m) }
func (*RequestContext) ProtoMessage()               {}
func (*RequestContext) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RequestContext) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestContext)(nil), "com.mindtickle.storage.pb.common.RequestContext")
}

func init() { proto.RegisterFile("common_messages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x52, 0x48, 0xce, 0xcf, 0xd5, 0xcb, 0xcd, 0xcc, 0x4b, 0x29, 0xc9, 0x4c, 0xce, 0xce,
	0x49, 0xd5, 0x2b, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x48, 0xd2, 0x83, 0x68, 0x50,
	0xd2, 0xe5, 0xe2, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x71, 0xce, 0xcf, 0x2b, 0x49, 0xad,
	0x28, 0x11, 0x92, 0xe6, 0xe2, 0x2c, 0x49, 0xcd, 0x4b, 0xcc, 0x2b, 0x89, 0xcf, 0x4c, 0x91, 0x60,
	0x56, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x80, 0x08, 0x78, 0xa6, 0x38, 0xa5, 0x70, 0xa9, 0x61, 0x37,
	0x12, 0x6c, 0x5f, 0x31, 0xc2, 0x60, 0x27, 0x6e, 0x67, 0x30, 0x1d, 0x00, 0x12, 0x8f, 0xd2, 0x4d,
	0xcf, 0x2c, 0xc9, 0x28, 0x05, 0xcb, 0xe9, 0xfb, 0x66, 0xe6, 0xa5, 0x84, 0x80, 0xf5, 0xea, 0xa3,
	0xe8, 0xd5, 0x2f, 0x48, 0xd2, 0x87, 0xe8, 0x4d, 0x62, 0x03, 0x8b, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xf5, 0x11, 0xf7, 0x0c, 0xd6, 0x00, 0x00, 0x00,
}