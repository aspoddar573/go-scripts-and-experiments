// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common.proto

It has these top-level messages:
	EmptyMsg
	PaginationRequest
	PaginationResponse
	MetaData
	PageInfo
	DateTimeWithTimezone
	DateTime
	ModuleEdge
	UserEdge
	Module
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import com_mindtickle_api_governance_enums "github.com/MindTickle/governance-protos/pb/enums"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SortOrder int32

const (
	SortOrder_ASC  SortOrder = 0
	SortOrder_DESC SortOrder = 1
)

var SortOrder_name = map[int32]string{
	0: "ASC",
	1: "DESC",
}
var SortOrder_value = map[string]int32{
	"ASC":  0,
	"DESC": 1,
}

func (x SortOrder) String() string {
	return proto.EnumName(SortOrder_name, int32(x))
}
func (SortOrder) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// has to be same as platform relevance...
type ModuleRelevance int32

const (
	ModuleRelevance_UNMARKED ModuleRelevance = 0
	ModuleRelevance_OPT      ModuleRelevance = 1
	ModuleRelevance_REQ      ModuleRelevance = 2
	ModuleRelevance_NONE     ModuleRelevance = 3
)

var ModuleRelevance_name = map[int32]string{
	0: "UNMARKED",
	1: "OPT",
	2: "REQ",
	3: "NONE",
}
var ModuleRelevance_value = map[string]int32{
	"UNMARKED": 0,
	"OPT":      1,
	"REQ":      2,
	"NONE":     3,
}

func (x ModuleRelevance) String() string {
	return proto.EnumName(ModuleRelevance_name, int32(x))
}
func (ModuleRelevance) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type EmptyMsg struct {
}

func (m *EmptyMsg) Reset()                    { *m = EmptyMsg{} }
func (m *EmptyMsg) String() string            { return proto.CompactTextString(m) }
func (*EmptyMsg) ProtoMessage()               {}
func (*EmptyMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PaginationRequest struct {
	From int32 `protobuf:"varint,1,opt,name=from" json:"from,omitempty"`
	Size int32 `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
}

func (m *PaginationRequest) Reset()                    { *m = PaginationRequest{} }
func (m *PaginationRequest) String() string            { return proto.CompactTextString(m) }
func (*PaginationRequest) ProtoMessage()               {}
func (*PaginationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PaginationRequest) GetFrom() int32 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *PaginationRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type PaginationResponse struct {
	TotalCount int32 `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	Cursor     int32 `protobuf:"varint,2,opt,name=cursor" json:"cursor,omitempty"`
	HasMore    bool  `protobuf:"varint,3,opt,name=has_more,json=hasMore" json:"has_more,omitempty"`
}

func (m *PaginationResponse) Reset()                    { *m = PaginationResponse{} }
func (m *PaginationResponse) String() string            { return proto.CompactTextString(m) }
func (*PaginationResponse) ProtoMessage()               {}
func (*PaginationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PaginationResponse) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *PaginationResponse) GetCursor() int32 {
	if m != nil {
		return m.Cursor
	}
	return 0
}

func (m *PaginationResponse) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

type MetaData struct {
	CreatedOn int64  `protobuf:"varint,1,opt,name=created_on,json=createdOn" json:"created_on,omitempty"`
	CreatedBy string `protobuf:"bytes,2,opt,name=created_by,json=createdBy" json:"created_by,omitempty"`
	UpdatedOn int64  `protobuf:"varint,3,opt,name=updated_on,json=updatedOn" json:"updated_on,omitempty"`
	UpdatedBy string `protobuf:"bytes,4,opt,name=updated_by,json=updatedBy" json:"updated_by,omitempty"`
}

func (m *MetaData) Reset()                    { *m = MetaData{} }
func (m *MetaData) String() string            { return proto.CompactTextString(m) }
func (*MetaData) ProtoMessage()               {}
func (*MetaData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MetaData) GetCreatedOn() int64 {
	if m != nil {
		return m.CreatedOn
	}
	return 0
}

func (m *MetaData) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *MetaData) GetUpdatedOn() int64 {
	if m != nil {
		return m.UpdatedOn
	}
	return 0
}

func (m *MetaData) GetUpdatedBy() string {
	if m != nil {
		return m.UpdatedBy
	}
	return ""
}

type PageInfo struct {
	HasNextPage     bool  `protobuf:"varint,1,opt,name=has_next_page,json=hasNextPage" json:"has_next_page,omitempty"`
	HasPreviousPage bool  `protobuf:"varint,2,opt,name=has_previous_page,json=hasPreviousPage" json:"has_previous_page,omitempty"`
	EndCursor       int32 `protobuf:"varint,3,opt,name=end_cursor,json=endCursor" json:"end_cursor,omitempty"`
	StartCursor     int32 `protobuf:"varint,4,opt,name=start_cursor,json=startCursor" json:"start_cursor,omitempty"`
	Total           int32 `protobuf:"varint,5,opt,name=total" json:"total,omitempty"`
}

func (m *PageInfo) Reset()                    { *m = PageInfo{} }
func (m *PageInfo) String() string            { return proto.CompactTextString(m) }
func (*PageInfo) ProtoMessage()               {}
func (*PageInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PageInfo) GetHasNextPage() bool {
	if m != nil {
		return m.HasNextPage
	}
	return false
}

func (m *PageInfo) GetHasPreviousPage() bool {
	if m != nil {
		return m.HasPreviousPage
	}
	return false
}

func (m *PageInfo) GetEndCursor() int32 {
	if m != nil {
		return m.EndCursor
	}
	return 0
}

func (m *PageInfo) GetStartCursor() int32 {
	if m != nil {
		return m.StartCursor
	}
	return 0
}

func (m *PageInfo) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type DateTimeWithTimezone struct {
	Time       *DateTime `protobuf:"bytes,1,opt,name=time" json:"time,omitempty"`
	TimeZoneId string    `protobuf:"bytes,2,opt,name=time_zone_id,json=timeZoneId" json:"time_zone_id,omitempty"`
}

func (m *DateTimeWithTimezone) Reset()                    { *m = DateTimeWithTimezone{} }
func (m *DateTimeWithTimezone) String() string            { return proto.CompactTextString(m) }
func (*DateTimeWithTimezone) ProtoMessage()               {}
func (*DateTimeWithTimezone) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DateTimeWithTimezone) GetTime() *DateTime {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *DateTimeWithTimezone) GetTimeZoneId() string {
	if m != nil {
		return m.TimeZoneId
	}
	return ""
}

type DateTime struct {
	Epoch int64 `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
}

func (m *DateTime) Reset()                    { *m = DateTime{} }
func (m *DateTime) String() string            { return proto.CompactTextString(m) }
func (*DateTime) ProtoMessage()               {}
func (*DateTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DateTime) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

type ModuleEdge struct {
	Cursor   int32  `protobuf:"varint,1,opt,name=cursor" json:"cursor,omitempty"`
	ModuleId string `protobuf:"bytes,2,opt,name=module_id,json=moduleId" json:"module_id,omitempty"`
}

func (m *ModuleEdge) Reset()                    { *m = ModuleEdge{} }
func (m *ModuleEdge) String() string            { return proto.CompactTextString(m) }
func (*ModuleEdge) ProtoMessage()               {}
func (*ModuleEdge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ModuleEdge) GetCursor() int32 {
	if m != nil {
		return m.Cursor
	}
	return 0
}

func (m *ModuleEdge) GetModuleId() string {
	if m != nil {
		return m.ModuleId
	}
	return ""
}

type UserEdge struct {
	Cursor int32  `protobuf:"varint,1,opt,name=cursor" json:"cursor,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *UserEdge) Reset()                    { *m = UserEdge{} }
func (m *UserEdge) String() string            { return proto.CompactTextString(m) }
func (*UserEdge) ProtoMessage()               {}
func (*UserEdge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UserEdge) GetCursor() int32 {
	if m != nil {
		return m.Cursor
	}
	return 0
}

func (m *UserEdge) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Module struct {
	ModuleId     string                                          `protobuf:"bytes,1,opt,name=module_id,json=moduleId" json:"module_id,omitempty"`
	ModuleName   string                                          `protobuf:"bytes,2,opt,name=module_name,json=moduleName" json:"module_name,omitempty"`
	DisplayOrder int64                                           `protobuf:"varint,3,opt,name=display_order,json=displayOrder" json:"display_order,omitempty"`
	Milestone    string                                          `protobuf:"bytes,4,opt,name=milestone" json:"milestone,omitempty"`
	ModuleType   com_mindtickle_api_governance_enums.ModuleType  `protobuf:"varint,5,opt,name=module_type,json=moduleType,enum=com.mindtickle.api.governance.enums.ModuleType" json:"module_type,omitempty"`
	ModuleState  com_mindtickle_api_governance_enums.ModuleState `protobuf:"varint,6,opt,name=module_state,json=moduleState,enum=com.mindtickle.api.governance.enums.ModuleState" json:"module_state,omitempty"`
	MetaData     *MetaData                                       `protobuf:"bytes,7,opt,name=meta_data,json=metaData" json:"meta_data,omitempty"`
}

func (m *Module) Reset()                    { *m = Module{} }
func (m *Module) String() string            { return proto.CompactTextString(m) }
func (*Module) ProtoMessage()               {}
func (*Module) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Module) GetModuleId() string {
	if m != nil {
		return m.ModuleId
	}
	return ""
}

func (m *Module) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *Module) GetDisplayOrder() int64 {
	if m != nil {
		return m.DisplayOrder
	}
	return 0
}

func (m *Module) GetMilestone() string {
	if m != nil {
		return m.Milestone
	}
	return ""
}

func (m *Module) GetModuleType() com_mindtickle_api_governance_enums.ModuleType {
	if m != nil {
		return m.ModuleType
	}
	return com_mindtickle_api_governance_enums.ModuleType_COURSE
}

func (m *Module) GetModuleState() com_mindtickle_api_governance_enums.ModuleState {
	if m != nil {
		return m.ModuleState
	}
	return com_mindtickle_api_governance_enums.ModuleState_MODULE_STATE_DRAFT
}

func (m *Module) GetMetaData() *MetaData {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func init() {
	proto.RegisterType((*EmptyMsg)(nil), "com.mindtickle.api.governance.common.EmptyMsg")
	proto.RegisterType((*PaginationRequest)(nil), "com.mindtickle.api.governance.common.PaginationRequest")
	proto.RegisterType((*PaginationResponse)(nil), "com.mindtickle.api.governance.common.PaginationResponse")
	proto.RegisterType((*MetaData)(nil), "com.mindtickle.api.governance.common.MetaData")
	proto.RegisterType((*PageInfo)(nil), "com.mindtickle.api.governance.common.PageInfo")
	proto.RegisterType((*DateTimeWithTimezone)(nil), "com.mindtickle.api.governance.common.DateTimeWithTimezone")
	proto.RegisterType((*DateTime)(nil), "com.mindtickle.api.governance.common.DateTime")
	proto.RegisterType((*ModuleEdge)(nil), "com.mindtickle.api.governance.common.ModuleEdge")
	proto.RegisterType((*UserEdge)(nil), "com.mindtickle.api.governance.common.UserEdge")
	proto.RegisterType((*Module)(nil), "com.mindtickle.api.governance.common.Module")
	proto.RegisterEnum("com.mindtickle.api.governance.common.SortOrder", SortOrder_name, SortOrder_value)
	proto.RegisterEnum("com.mindtickle.api.governance.common.ModuleRelevance", ModuleRelevance_name, ModuleRelevance_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 721 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcb, 0x8e, 0xdb, 0x36,
	0x14, 0x8d, 0x6c, 0x8f, 0x2d, 0x5f, 0x3b, 0x8d, 0x43, 0x04, 0xad, 0xfb, 0x48, 0xeb, 0xaa, 0x5d,
	0x04, 0x03, 0x54, 0x6e, 0x93, 0xe5, 0xac, 0xc6, 0x8f, 0x85, 0x11, 0xf8, 0x51, 0xd9, 0x41, 0x81,
	0x6c, 0x04, 0x5a, 0xba, 0xb1, 0x84, 0x9a, 0xa4, 0x2a, 0x52, 0x83, 0x28, 0xe8, 0xba, 0x1f, 0xd4,
	0xff, 0x2b, 0x50, 0xf0, 0xa1, 0x79, 0x6c, 0xda, 0x64, 0x45, 0xf2, 0xdc, 0x73, 0x0e, 0x8f, 0x2f,
	0xaf, 0x0c, 0xc3, 0x44, 0x30, 0x26, 0x78, 0x58, 0x94, 0x42, 0x09, 0xf2, 0x63, 0x22, 0x58, 0xc8,
	0x72, 0x9e, 0xaa, 0x3c, 0xf9, 0xfd, 0x8c, 0x21, 0x2d, 0xf2, 0xf0, 0x24, 0x6e, 0xb0, 0xe4, 0x94,
	0x27, 0x18, 0x5a, 0xee, 0x57, 0xc4, 0xae, 0x53, 0xe4, 0x15, 0x93, 0x56, 0x19, 0x00, 0xf8, 0x4b,
	0x56, 0xa8, 0x7a, 0x2d, 0x4f, 0xc1, 0x15, 0x3c, 0xdd, 0xd1, 0x53, 0xce, 0xa9, 0xca, 0x05, 0x8f,
	0xf0, 0x8f, 0x0a, 0xa5, 0x22, 0x04, 0x3a, 0xef, 0x4a, 0xc1, 0xc6, 0xde, 0xc4, 0x7b, 0x71, 0x11,
	0x99, 0xbd, 0xc6, 0x64, 0xfe, 0x01, 0xc7, 0x2d, 0x8b, 0xe9, 0x7d, 0x90, 0x01, 0xb9, 0x2f, 0x96,
	0x85, 0xe0, 0x12, 0xc9, 0x77, 0x30, 0x50, 0x42, 0xd1, 0x73, 0x9c, 0x88, 0x8a, 0x2b, 0x67, 0x02,
	0x06, 0x9a, 0x6b, 0x84, 0x7c, 0x0e, 0xdd, 0xa4, 0x2a, 0xa5, 0x28, 0x9d, 0x99, 0x3b, 0x91, 0x2f,
	0xc1, 0xcf, 0xa8, 0x8c, 0x99, 0x28, 0x71, 0xdc, 0x9e, 0x78, 0x2f, 0xfc, 0xa8, 0x97, 0x51, 0xb9,
	0x16, 0x25, 0x06, 0x7f, 0x79, 0xe0, 0xaf, 0x51, 0xd1, 0x05, 0x55, 0x94, 0x3c, 0x07, 0x48, 0x4a,
	0xa4, 0x0a, 0xd3, 0x58, 0x70, 0xe3, 0xdf, 0x8e, 0xfa, 0x0e, 0xd9, 0xf2, 0xfb, 0xe5, 0x63, 0x6d,
	0xae, 0xe8, 0xdf, 0x96, 0x67, 0xb5, 0x2e, 0x57, 0x45, 0xda, 0xa8, 0xdb, 0x56, 0xed, 0x10, 0xab,
	0x6e, 0xca, 0xc7, 0x7a, 0xdc, 0xb1, 0x6a, 0x87, 0xcc, 0xea, 0xe0, 0x6f, 0x0f, 0xfc, 0x1d, 0x3d,
	0xe1, 0x8a, 0xbf, 0x13, 0x24, 0x80, 0xc7, 0x3a, 0x30, 0xc7, 0xf7, 0x2a, 0x2e, 0xe8, 0x09, 0x4d,
	0x16, 0x3f, 0x1a, 0x64, 0x54, 0x6e, 0xf0, 0xbd, 0xd2, 0x3c, 0x72, 0x09, 0x4f, 0x35, 0xa7, 0x28,
	0xf1, 0x26, 0x17, 0x95, 0xb4, 0xbc, 0x96, 0xe1, 0x3d, 0xc9, 0xa8, 0xdc, 0x39, 0xdc, 0x70, 0x9f,
	0x03, 0x20, 0x4f, 0x63, 0xd7, 0x9c, 0xb6, 0x69, 0x4e, 0x1f, 0x79, 0x3a, 0xb7, 0xfd, 0xf9, 0x1e,
	0x86, 0x52, 0xd1, 0x52, 0x35, 0x84, 0x8e, 0x21, 0x0c, 0x0c, 0xe6, 0x28, 0xcf, 0xe0, 0xc2, 0x34,
	0x7a, 0x7c, 0x61, 0x6a, 0xf6, 0x10, 0xfc, 0x09, 0xcf, 0x16, 0x54, 0xe1, 0x21, 0x67, 0xf8, 0x5b,
	0xae, 0x32, 0xbd, 0x7e, 0x10, 0x1c, 0xc9, 0x0c, 0x3a, 0x2a, 0x67, 0x36, 0xf6, 0xe0, 0x65, 0x18,
	0x7e, 0xcc, 0x44, 0x85, 0x8d, 0x53, 0x64, 0xb4, 0x64, 0x02, 0x43, 0xbd, 0xc6, 0xda, 0x30, 0xce,
	0x53, 0xd7, 0x6f, 0xd0, 0xd8, 0x5b, 0xc1, 0x71, 0x95, 0x06, 0x13, 0xf0, 0x1b, 0x8d, 0xce, 0x87,
	0x85, 0x48, 0x32, 0xf7, 0x6a, 0xf6, 0x10, 0x5c, 0x03, 0xac, 0x45, 0x5a, 0x9d, 0x71, 0x99, 0x9e,
	0xf0, 0xde, 0x78, 0x78, 0x0f, 0xc6, 0xe3, 0x6b, 0xe8, 0x33, 0xc3, 0xba, 0xbb, 0xc6, 0xb7, 0xc0,
	0x2a, 0x0d, 0xae, 0xc0, 0x7f, 0x23, 0xb1, 0xfc, 0x4f, 0x83, 0x2f, 0xa0, 0x57, 0x49, 0x2c, 0xef,
	0xe4, 0x5d, 0x7d, 0x5c, 0xa5, 0xc1, 0x3f, 0x2d, 0xe8, 0xda, 0x00, 0x0f, 0x2f, 0xf1, 0x1e, 0x5e,
	0xa2, 0x27, 0xdb, 0x15, 0x39, 0x65, 0xd8, 0xfc, 0x54, 0x0b, 0x6d, 0x28, 0x43, 0xf2, 0x03, 0x3c,
	0x4e, 0x73, 0x59, 0x9c, 0x69, 0x1d, 0x8b, 0x32, 0xc5, 0xd2, 0x8d, 0xd7, 0xd0, 0x81, 0x5b, 0x8d,
	0x91, 0x6f, 0xa0, 0xcf, 0xf2, 0x33, 0x4a, 0x25, 0x38, 0x36, 0x03, 0x76, 0x0b, 0x90, 0xdd, 0xed,
	0x1d, 0xaa, 0x2e, 0xd0, 0xbc, 0xe3, 0x67, 0x2f, 0xa7, 0xff, 0xf3, 0x34, 0xf6, 0xeb, 0xb6, 0x3f,
	0xe1, 0x50, 0x17, 0xd8, 0x84, 0xd2, 0x7b, 0xb2, 0x87, 0xa1, 0x73, 0x94, 0x8a, 0x2a, 0x1c, 0x77,
	0x8d, 0xe5, 0xcf, 0x9f, 0x60, 0xb9, 0xd7, 0xba, 0xc8, 0xe5, 0x32, 0x07, 0xf2, 0x1a, 0xfa, 0x0c,
	0x15, 0x8d, 0x53, 0xaa, 0xe8, 0xb8, 0xf7, 0x29, 0xf3, 0xd3, 0x7c, 0xc6, 0x91, 0xcf, 0xdc, 0xee,
	0xf2, 0x5b, 0xe8, 0xef, 0x45, 0xa9, 0x6c, 0x7b, 0x7a, 0xd0, 0xbe, 0xde, 0xcf, 0x47, 0x8f, 0x88,
	0x0f, 0x9d, 0xc5, 0x72, 0x3f, 0x1f, 0x79, 0x97, 0x57, 0xf0, 0xc4, 0x06, 0x89, 0xf0, 0x8c, 0x37,
	0xda, 0x8c, 0x0c, 0xc1, 0x7f, 0xb3, 0x59, 0x5f, 0x47, 0xaf, 0x97, 0x8b, 0xd1, 0x23, 0xad, 0xd9,
	0xee, 0x0e, 0x23, 0x4f, 0x6f, 0xa2, 0xe5, 0xaf, 0xa3, 0x96, 0x16, 0x6f, 0xb6, 0x9b, 0xe5, 0xa8,
	0x3d, 0x7b, 0xf5, 0xf6, 0x97, 0x53, 0xae, 0xb2, 0xea, 0xa8, 0x03, 0x4c, 0xd7, 0x39, 0x4f, 0x0f,
	0x26, 0xde, 0xf4, 0x2e, 0xda, 0x4f, 0xe6, 0x7f, 0x51, 0x4e, 0x8b, 0xe3, 0xd4, 0x86, 0x3c, 0x76,
	0x0d, 0xf2, 0xea, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xf9, 0x91, 0xd1, 0x73, 0x05, 0x00,
	0x00,
}
