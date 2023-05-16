// Code generated by protoc-gen-go. DO NOT EDIT.
// source: template_workflows.proto

/*
Package templateWorkflows is a generated protocol buffer package.

It is generated from these files:
	template_workflows.proto

It has these top-level messages:
	Series
	CreateModuleFromTemplate
*/
package templateWorkflows

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import com_mindtickle_api_governance_common "github.com/MindTickle/governance-protos/pb/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Series struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Milestone  string `protobuf:"bytes,2,opt,name=milestone" json:"milestone,omitempty"`
	IsNew      bool   `protobuf:"varint,3,opt,name=is_new,json=isNew" json:"is_new,omitempty"`
	StringName string `protobuf:"bytes,4,opt,name=string_name,json=stringName" json:"string_name,omitempty"`
}

func (m *Series) Reset()                    { *m = Series{} }
func (m *Series) String() string            { return proto.CompactTextString(m) }
func (*Series) ProtoMessage()               {}
func (*Series) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Series) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Series) GetMilestone() string {
	if m != nil {
		return m.Milestone
	}
	return ""
}

func (m *Series) GetIsNew() bool {
	if m != nil {
		return m.IsNew
	}
	return false
}

func (m *Series) GetStringName() string {
	if m != nil {
		return m.StringName
	}
	return ""
}

type CreateModuleFromTemplate struct {
	TemplateId string                                     `protobuf:"bytes,1,opt,name=template_id,json=templateId" json:"template_id,omitempty"`
	ModuleName string                                     `protobuf:"bytes,2,opt,name=module_name,json=moduleName" json:"module_name,omitempty"`
	Series     *Series                                    `protobuf:"bytes,3,opt,name=series" json:"series,omitempty"`
	Meta       *com_mindtickle_api_governance_common.Meta `protobuf:"bytes,4,opt,name=meta" json:"meta,omitempty"`
}

func (m *CreateModuleFromTemplate) Reset()                    { *m = CreateModuleFromTemplate{} }
func (m *CreateModuleFromTemplate) String() string            { return proto.CompactTextString(m) }
func (*CreateModuleFromTemplate) ProtoMessage()               {}
func (*CreateModuleFromTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateModuleFromTemplate) GetTemplateId() string {
	if m != nil {
		return m.TemplateId
	}
	return ""
}

func (m *CreateModuleFromTemplate) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *CreateModuleFromTemplate) GetSeries() *Series {
	if m != nil {
		return m.Series
	}
	return nil
}

func (m *CreateModuleFromTemplate) GetMeta() *com_mindtickle_api_governance_common.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*Series)(nil), "com.mindtickle.api.governance.templateWorkflows.Series")
	proto.RegisterType((*CreateModuleFromTemplate)(nil), "com.mindtickle.api.governance.templateWorkflows.CreateModuleFromTemplate")
}

func init() { proto.RegisterFile("template_workflows.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0x03, 0x31,
	0x14, 0xc4, 0xd9, 0x5a, 0x17, 0xfb, 0x0a, 0x1e, 0x22, 0xc2, 0x22, 0x82, 0xa5, 0xa7, 0x22, 0x98,
	0x85, 0x7a, 0xf0, 0x22, 0x3d, 0x28, 0x08, 0x1e, 0x5a, 0x61, 0x2d, 0x08, 0x5e, 0x4a, 0xba, 0x79,
	0xd6, 0x47, 0x37, 0xc9, 0x92, 0xa4, 0xee, 0xa7, 0xf6, 0x3b, 0xc8, 0x26, 0xfd, 0x73, 0xf0, 0x20,
	0x9e, 0x16, 0xe6, 0xed, 0x6f, 0x98, 0x99, 0x40, 0xe6, 0x51, 0xd5, 0x95, 0xf0, 0xb8, 0x68, 0x8c,
	0x5d, 0x7f, 0x54, 0xa6, 0x71, 0xbc, 0xb6, 0xc6, 0x1b, 0x96, 0x97, 0x46, 0x71, 0x45, 0x5a, 0x7a,
	0x2a, 0xd7, 0x15, 0x72, 0x51, 0x13, 0x5f, 0x99, 0x2f, 0xb4, 0x5a, 0xe8, 0x12, 0xf9, 0x8e, 0x7b,
	0xdb, 0x61, 0x17, 0x67, 0xa5, 0x51, 0xca, 0xe8, 0x3c, 0x7e, 0xa2, 0xcb, 0x50, 0x43, 0xfa, 0x8a,
	0x96, 0xd0, 0xb1, 0x53, 0xe8, 0x90, 0xcc, 0x92, 0x41, 0x32, 0xea, 0x15, 0x1d, 0x92, 0xec, 0x12,
	0x7a, 0x8a, 0x2a, 0x74, 0xde, 0x68, 0xcc, 0x3a, 0x41, 0x3e, 0x08, 0xec, 0x1c, 0x52, 0x72, 0x0b,
	0x8d, 0x4d, 0x76, 0x34, 0x48, 0x46, 0x27, 0xc5, 0x31, 0xb9, 0x19, 0x36, 0xec, 0x0a, 0xfa, 0xce,
	0x5b, 0xd2, 0xab, 0x85, 0x16, 0x0a, 0xb3, 0x6e, 0xc0, 0x20, 0x4a, 0x33, 0xa1, 0x70, 0xf8, 0x9d,
	0x40, 0xf6, 0x68, 0x51, 0x78, 0x9c, 0x1a, 0xb9, 0xa9, 0xf0, 0xc9, 0x1a, 0x35, 0xdf, 0x46, 0x6d,
	0xe9, 0x7d, 0xdd, 0x7d, 0x16, 0xd8, 0x49, 0xcf, 0xb2, 0xfd, 0x41, 0x05, 0x2c, 0xda, 0xc7, 0x54,
	0x10, 0xa5, 0xd6, 0x9e, 0xbd, 0x40, 0xea, 0x42, 0x9d, 0x10, 0xab, 0x3f, 0xbe, 0xe3, 0xff, 0x5c,
	0x89, 0xc7, 0x35, 0x8a, 0xad, 0x0d, 0x9b, 0x40, 0x57, 0xa1, 0x17, 0xa1, 0x49, 0x7f, 0x7c, 0xfd,
	0x87, 0xdd, 0x76, 0xda, 0x29, 0x7a, 0x51, 0x04, 0xee, 0x61, 0xf2, 0x7e, 0xbf, 0x22, 0xff, 0xb9,
	0x59, 0xb6, 0xb7, 0x7c, 0x4a, 0x5a, 0xce, 0x03, 0x99, 0x1f, 0xa8, 0x9b, 0xf0, 0x12, 0x2e, 0xaf,
	0x97, 0xf9, 0xaf, 0x38, 0xcb, 0x34, 0x1c, 0x6f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x54,
	0xb1, 0x68, 0x08, 0x02, 0x00, 0x00,
}
