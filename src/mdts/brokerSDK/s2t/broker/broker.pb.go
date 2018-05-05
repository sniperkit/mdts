// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mdts/brokerSDK/s2t/broker/broker.proto

/*
Package broker is a generated protocol buffer package.

It is generated from these files:
	mdts/brokerSDK/s2t/broker/broker.proto

It has these top-level messages:
*/
package broker

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import brokermsg "mdts/protocols/brokermsg"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for S2TBroker service

type S2TBrokerClient interface {
	// Request: service -> dts -> broker -> dts -> third
	TransforDataToThird(ctx context.Context, in *brokermsg.ParamToThird, opts ...grpc.CallOption) (*brokermsg.ResultToThird, error)
	// Response: service <- dts <- broker <- dts <- third
	TransforDataFromThird(ctx context.Context, in *brokermsg.ParamFromThird, opts ...grpc.CallOption) (*brokermsg.ResultFromThird, error)
}

type s2TBrokerClient struct {
	cc *grpc.ClientConn
}

func NewS2TBrokerClient(cc *grpc.ClientConn) S2TBrokerClient {
	return &s2TBrokerClient{cc}
}

func (c *s2TBrokerClient) TransforDataToThird(ctx context.Context, in *brokermsg.ParamToThird, opts ...grpc.CallOption) (*brokermsg.ResultToThird, error) {
	out := new(brokermsg.ResultToThird)
	err := grpc.Invoke(ctx, "/broker.S2TBroker/TransforDataToThird", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s2TBrokerClient) TransforDataFromThird(ctx context.Context, in *brokermsg.ParamFromThird, opts ...grpc.CallOption) (*brokermsg.ResultFromThird, error) {
	out := new(brokermsg.ResultFromThird)
	err := grpc.Invoke(ctx, "/broker.S2TBroker/TransforDataFromThird", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for S2TBroker service

type S2TBrokerServer interface {
	// Request: service -> dts -> broker -> dts -> third
	TransforDataToThird(context.Context, *brokermsg.ParamToThird) (*brokermsg.ResultToThird, error)
	// Response: service <- dts <- broker <- dts <- third
	TransforDataFromThird(context.Context, *brokermsg.ParamFromThird) (*brokermsg.ResultFromThird, error)
}

func RegisterS2TBrokerServer(s *grpc.Server, srv S2TBrokerServer) {
	s.RegisterService(&_S2TBroker_serviceDesc, srv)
}

func _S2TBroker_TransforDataToThird_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(brokermsg.ParamToThird)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S2TBrokerServer).TransforDataToThird(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/broker.S2TBroker/TransforDataToThird",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S2TBrokerServer).TransforDataToThird(ctx, req.(*brokermsg.ParamToThird))
	}
	return interceptor(ctx, in, info, handler)
}

func _S2TBroker_TransforDataFromThird_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(brokermsg.ParamFromThird)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S2TBrokerServer).TransforDataFromThird(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/broker.S2TBroker/TransforDataFromThird",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S2TBrokerServer).TransforDataFromThird(ctx, req.(*brokermsg.ParamFromThird))
	}
	return interceptor(ctx, in, info, handler)
}

var _S2TBroker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "broker.S2TBroker",
	HandlerType: (*S2TBrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransforDataToThird",
			Handler:    _S2TBroker_TransforDataToThird_Handler,
		},
		{
			MethodName: "TransforDataFromThird",
			Handler:    _S2TBroker_TransforDataFromThird_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mdts/brokerSDK/s2t/broker/broker.proto",
}

func init() { proto.RegisterFile("mdts/brokerSDK/s2t/broker/broker.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcb, 0x4d, 0x29, 0x29,
	0xd6, 0x4f, 0x2a, 0xca, 0xcf, 0x4e, 0x2d, 0x0a, 0x76, 0xf1, 0xd6, 0x2f, 0x36, 0x2a, 0x81, 0xf2,
	0xa0, 0x94, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x10, 0x1b, 0x84, 0x27, 0xa5, 0x01, 0x56, 0x0f,
	0x16, 0x4b, 0xce, 0xcf, 0x81, 0xe9, 0xcc, 0x2d, 0x4e, 0x47, 0xb0, 0x20, 0x3a, 0x8c, 0x56, 0x32,
	0x72, 0x71, 0x06, 0x1b, 0x85, 0x38, 0x81, 0x85, 0x85, 0xbc, 0xb8, 0x84, 0x43, 0x8a, 0x12, 0xf3,
	0x8a, 0xd3, 0xf2, 0x8b, 0x5c, 0x12, 0x4b, 0x12, 0x43, 0xf2, 0x43, 0x32, 0x32, 0x8b, 0x52, 0x84,
	0xc4, 0xf5, 0x10, 0xda, 0x02, 0x12, 0x8b, 0x12, 0x73, 0xa1, 0x12, 0x52, 0x12, 0x48, 0x12, 0x41,
	0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x50, 0x19, 0x25, 0x06, 0xa1, 0x00, 0x2e, 0x51, 0x64, 0xb3, 0xdc,
	0x8a, 0xf2, 0x73, 0x21, 0xa6, 0x49, 0xa2, 0x9b, 0x06, 0x97, 0x92, 0x92, 0xc2, 0x30, 0x0f, 0x2e,
	0xa7, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0xb2, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x85, 0x82,
	0xa9, 0x0e, 0x01, 0x00, 0x00,
}
