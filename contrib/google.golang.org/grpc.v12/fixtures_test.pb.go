// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016 Datadog, Inc.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fixtures_test.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	fixtures_test.proto

It has these top-level messages:
	FixtureRequest
	FixtureReply
*/
package grpc

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc1 "google.golang.org/grpc"
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

// The request message containing the user's name.
type FixtureRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *FixtureRequest) Reset()                    { *m = FixtureRequest{} }
func (m *FixtureRequest) String() string            { return proto.CompactTextString(m) }
func (*FixtureRequest) ProtoMessage()               {}
func (*FixtureRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FixtureRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type FixtureReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *FixtureReply) Reset()                    { *m = FixtureReply{} }
func (m *FixtureReply) String() string            { return proto.CompactTextString(m) }
func (*FixtureReply) ProtoMessage()               {}
func (*FixtureReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FixtureReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*FixtureRequest)(nil), "grpc.FixtureRequest")
	proto.RegisterType((*FixtureReply)(nil), "grpc.FixtureReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for Fixture service

type FixtureClient interface {
	Ping(ctx context.Context, in *FixtureRequest, opts ...grpc1.CallOption) (*FixtureReply, error)
}

type fixtureClient struct {
	cc *grpc1.ClientConn
}

func NewFixtureClient(cc *grpc1.ClientConn) FixtureClient {
	return &fixtureClient{cc}
}

func (c *fixtureClient) Ping(ctx context.Context, in *FixtureRequest, opts ...grpc1.CallOption) (*FixtureReply, error) {
	out := new(FixtureReply)
	err := grpc1.Invoke(ctx, "/grpc.Fixture/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Fixture service

type FixtureServer interface {
	Ping(context.Context, *FixtureRequest) (*FixtureReply, error)
}

func RegisterFixtureServer(s *grpc1.Server, srv FixtureServer) {
	s.RegisterService(&_Fixture_serviceDesc, srv)
}

func _Fixture_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(FixtureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixtureServer).Ping(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Fixture/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixtureServer).Ping(ctx, req.(*FixtureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Fixture_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.Fixture",
	HandlerType: (*FixtureServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Fixture_Ping_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "fixtures_test.proto",
}

func init() { proto.RegisterFile("fixtures_test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xcb, 0xac, 0x28,
	0x29, 0x2d, 0x4a, 0x2d, 0x8e, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x49, 0x2f, 0x2a, 0x48, 0x56, 0x52, 0xe1, 0xe2, 0x73, 0x83, 0x48, 0x06, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x81, 0xd9, 0x4a, 0x1a, 0x5c, 0x3c, 0x70, 0x55, 0x05, 0x39, 0x95, 0x42, 0x12, 0x5c,
	0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0x30, 0x65, 0x30, 0xae, 0x91, 0x2d, 0x17, 0x3b, 0x54,
	0xa5, 0x90, 0x11, 0x17, 0x4b, 0x40, 0x66, 0x5e, 0xba, 0x90, 0x88, 0x1e, 0xc8, 0x26, 0x3d, 0x54,
	0x6b, 0xa4, 0x84, 0xd0, 0x44, 0x0b, 0x72, 0x2a, 0x95, 0x18, 0x9c, 0x74, 0xb8, 0x24, 0x33, 0xf3,
	0x21, 0x32, 0xa9, 0x15, 0x89, 0xb9, 0x05, 0x39, 0xa9, 0xc5, 0x7a, 0x20, 0x37, 0x83, 0x44, 0x9c,
	0x78, 0x43, 0x52, 0x8b, 0x4b, 0xdc, 0x83, 0x02, 0x9c, 0x03, 0x40, 0x1e, 0x08, 0x60, 0x4c, 0x62,
	0x03, 0xfb, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x18, 0x42, 0x90, 0x4d, 0xe0, 0x00, 0x00,
	0x00,
}
