// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models.proto

/*
Package fflagcheckapi is a generated protocol buffer package.

It is generated from these files:
	models.proto

It has these top-level messages:
	FlagQuery
	FlagResult
*/
package fflagcheckapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type FlagQuery struct {
	AccountId string `protobuf:"bytes,1,opt,name=accountId" json:"accountId,omitempty"`
	FlagName  string `protobuf:"bytes,2,opt,name=flagName" json:"flagName,omitempty"`
}

func (m *FlagQuery) Reset()                    { *m = FlagQuery{} }
func (m *FlagQuery) String() string            { return proto.CompactTextString(m) }
func (*FlagQuery) ProtoMessage()               {}
func (*FlagQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FlagQuery) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *FlagQuery) GetFlagName() string {
	if m != nil {
		return m.FlagName
	}
	return ""
}

type FlagResult struct {
	Found bool `protobuf:"varint,1,opt,name=found" json:"found,omitempty"`
	Value bool `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *FlagResult) Reset()                    { *m = FlagResult{} }
func (m *FlagResult) String() string            { return proto.CompactTextString(m) }
func (*FlagResult) ProtoMessage()               {}
func (*FlagResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FlagResult) GetFound() bool {
	if m != nil {
		return m.Found
	}
	return false
}

func (m *FlagResult) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto.RegisterType((*FlagQuery)(nil), "fflagcheckapi.FlagQuery")
	proto.RegisterType((*FlagResult)(nil), "fflagcheckapi.FlagResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FeatureFlag service

type FeatureFlagClient interface {
	GetFlag(ctx context.Context, in *FlagQuery, opts ...grpc.CallOption) (*FlagResult, error)
}

type featureFlagClient struct {
	cc *grpc.ClientConn
}

func NewFeatureFlagClient(cc *grpc.ClientConn) FeatureFlagClient {
	return &featureFlagClient{cc}
}

func (c *featureFlagClient) GetFlag(ctx context.Context, in *FlagQuery, opts ...grpc.CallOption) (*FlagResult, error) {
	out := new(FlagResult)
	err := grpc.Invoke(ctx, "/fflagcheckapi.FeatureFlag/GetFlag", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FeatureFlag service

type FeatureFlagServer interface {
	GetFlag(context.Context, *FlagQuery) (*FlagResult, error)
}

func RegisterFeatureFlagServer(s *grpc.Server, srv FeatureFlagServer) {
	s.RegisterService(&_FeatureFlag_serviceDesc, srv)
}

func _FeatureFlag_GetFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlagQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureFlagServer).GetFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fflagcheckapi.FeatureFlag/GetFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureFlagServer).GetFlag(ctx, req.(*FlagQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeatureFlag_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fflagcheckapi.FeatureFlag",
	HandlerType: (*FeatureFlagServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFlag",
			Handler:    _FeatureFlag_GetFlag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "models.proto",
}

func init() { proto.RegisterFile("models.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xcd, 0x4f, 0x49,
	0xcd, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4d, 0x4b, 0xcb, 0x49, 0x4c, 0x4f,
	0xce, 0x48, 0x4d, 0xce, 0x4e, 0x2c, 0xc8, 0x54, 0x72, 0xe5, 0xe2, 0x74, 0xcb, 0x49, 0x4c, 0x0f,
	0x2c, 0x4d, 0x2d, 0xaa, 0x14, 0x92, 0xe1, 0xe2, 0x4c, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xf1,
	0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x08, 0x08, 0x49, 0x71, 0x71, 0x80, 0xb4,
	0xfa, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x81, 0x25, 0xe1, 0x7c, 0x25, 0x0b, 0x2e, 0x2e, 0x90, 0x31,
	0x41, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x42, 0x22, 0x5c, 0xac, 0x69, 0xf9, 0xa5, 0x79, 0x10, 0x33,
	0x38, 0x82, 0x20, 0x1c, 0x90, 0x68, 0x59, 0x62, 0x4e, 0x29, 0x44, 0x33, 0x47, 0x10, 0x84, 0x63,
	0xe4, 0xcf, 0xc5, 0xed, 0x96, 0x9a, 0x58, 0x52, 0x5a, 0x94, 0x0a, 0x32, 0x40, 0xc8, 0x81, 0x8b,
	0xdd, 0x3d, 0xb5, 0x04, 0xcc, 0x94, 0xd0, 0x43, 0x71, 0xaa, 0x1e, 0xdc, 0x9d, 0x52, 0x92, 0x58,
	0x64, 0x20, 0x56, 0x2b, 0x31, 0x24, 0xb1, 0x81, 0xfd, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x76, 0x13, 0xde, 0x63, 0xf7, 0x00, 0x00, 0x00,
}
