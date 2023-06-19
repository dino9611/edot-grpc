// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/chart.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChartClient is the client API for Chart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChartClient interface {
	GetChart(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Chart, error)
}

type chartClient struct {
	cc grpc.ClientConnInterface
}

func NewChartClient(cc grpc.ClientConnInterface) ChartClient {
	return &chartClient{cc}
}

func (c *chartClient) GetChart(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Chart, error) {
	out := new(Chart)
	err := c.cc.Invoke(ctx, "/Chart/GetChart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChartServer is the server API for Chart service.
// All implementations must embed UnimplementedChartServer
// for forward compatibility
type ChartServer interface {
	GetChart(context.Context, *emptypb.Empty) (*Chart, error)
	mustEmbedUnimplementedChartServer()
}

// UnimplementedChartServer must be embedded to have forward compatible implementations.
type UnimplementedChartServer struct {
}

func (UnimplementedChartServer) GetChart(context.Context, *emptypb.Empty) (*Chart, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChart not implemented")
}
func (UnimplementedChartServer) mustEmbedUnimplementedChartServer() {}

// UnsafeChartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChartServer will
// result in compilation errors.
type UnsafeChartServer interface {
	mustEmbedUnimplementedChartServer()
}

func RegisterChartServer(s grpc.ServiceRegistrar, srv ChartServer) {
	s.RegisterService(&Chart_ServiceDesc, srv)
}

func _Chart_GetChart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChartServer).GetChart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chart/GetChart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChartServer).GetChart(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Chart_ServiceDesc is the grpc.ServiceDesc for Chart service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chart_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Chart",
	HandlerType: (*ChartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChart",
			Handler:    _Chart_GetChart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/chart.proto",
}
