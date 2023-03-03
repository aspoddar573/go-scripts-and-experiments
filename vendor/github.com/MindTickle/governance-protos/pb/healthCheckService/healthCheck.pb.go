// Code generated by protoc-gen-go. DO NOT EDIT.
// source: healthCheck.proto

/*
Package healthCheckService is a generated protocol buffer package.

It is generated from these files:
	healthCheck.proto

It has these top-level messages:
	HealthCheckResponse
*/
package healthCheckService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import com_mindtickle_api_governance_common "github.com/MindTickle/governance-protos/pb/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A boolean response for service liveness
type HealthCheckResponse struct {
	Healthy bool `protobuf:"varint,1,opt,name=Healthy" json:"Healthy,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HealthCheckResponse) GetHealthy() bool {
	if m != nil {
		return m.Healthy
	}
	return false
}

func init() {
	proto.RegisterType((*HealthCheckResponse)(nil), "com.mindtickle.api.governance.health.v1.HealthCheckResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HealthCheck service

type HealthCheckClient interface {
	// Check service liveness with a health check RPC call
	HealthCheck(ctx context.Context, in *com_mindtickle_api_governance_common.EmptyMsg, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthCheckClient struct {
	cc *grpc.ClientConn
}

func NewHealthCheckClient(cc *grpc.ClientConn) HealthCheckClient {
	return &healthCheckClient{cc}
}

func (c *healthCheckClient) HealthCheck(ctx context.Context, in *com_mindtickle_api_governance_common.EmptyMsg, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/com.mindtickle.api.governance.health.v1.HealthCheck/HealthCheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HealthCheck service

type HealthCheckServer interface {
	// Check service liveness with a health check RPC call
	HealthCheck(context.Context, *com_mindtickle_api_governance_common.EmptyMsg) (*HealthCheckResponse, error)
}

func RegisterHealthCheckServer(s *grpc.Server, srv HealthCheckServer) {
	s.RegisterService(&_HealthCheck_serviceDesc, srv)
}

func _HealthCheck_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(com_mindtickle_api_governance_common.EmptyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthCheckServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.mindtickle.api.governance.health.v1.HealthCheck/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthCheckServer).HealthCheck(ctx, req.(*com_mindtickle_api_governance_common.EmptyMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthCheck_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.mindtickle.api.governance.health.v1.HealthCheck",
	HandlerType: (*HealthCheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _HealthCheck_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "healthCheck.proto",
}

func init() { proto.RegisterFile("healthCheck.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0x70, 0xce, 0x48, 0x4d, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x4f,
	0xce, 0xcf, 0xd5, 0xcb, 0xcd, 0xcc, 0x4b, 0x29, 0xc9, 0x4c, 0xce, 0xce, 0x49, 0xd5, 0x4b, 0x2c,
	0xc8, 0xd4, 0x4b, 0xcf, 0x2f, 0x4b, 0x2d, 0xca, 0x4b, 0xcc, 0x4b, 0x4e, 0xd5, 0x83, 0x68, 0xd0,
	0x2b, 0x33, 0x94, 0x12, 0x4e, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x87, 0x50, 0x10, 0xdd, 0x4a,
	0xfa, 0x5c, 0xc2, 0x1e, 0x08, 0x23, 0x83, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x24,
	0xb8, 0xd8, 0x21, 0xc2, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x30, 0xae, 0x51, 0x17,
	0x23, 0x17, 0x37, 0x92, 0x0e, 0xa1, 0x6a, 0x54, 0xae, 0x9e, 0x1e, 0x7e, 0xe7, 0x40, 0x2d, 0x77,
	0xcd, 0x2d, 0x28, 0xa9, 0xf4, 0x2d, 0x4e, 0x97, 0xb2, 0xd1, 0x23, 0xd2, 0xf9, 0x7a, 0x58, 0x9c,
	0xe9, 0x64, 0x1f, 0x65, 0x9b, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0x04, 0x32, 0x45, 0xdf, 0x37, 0x33,
	0x2f, 0x25, 0x04, 0x6c, 0x8a, 0x3e, 0xc2, 0x04, 0x5d, 0xb0, 0x3f, 0x8b, 0xf5, 0x0b, 0x92, 0xf4,
	0x91, 0xc2, 0x2e, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x35, 0x89, 0x0d, 0x2c, 0x6b, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x7d, 0x1b, 0x08, 0xd8, 0x58, 0x01, 0x00, 0x00,
}