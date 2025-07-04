// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: demo.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Stats_ProcessDemo_FullMethodName = "/pb.Stats/ProcessDemo"
	Stats_ReplayDemo_FullMethodName  = "/pb.Stats/ReplayDemo"
)

// StatsClient is the client API for Stats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatsClient interface {
	ProcessDemo(ctx context.Context, in *ProcessDemoReq, opts ...grpc.CallOption) (*ProcessDemoResponse, error)
	ReplayDemo(ctx context.Context, in *ProcessDemoReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Pos], error)
}

type statsClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsClient(cc grpc.ClientConnInterface) StatsClient {
	return &statsClient{cc}
}

func (c *statsClient) ProcessDemo(ctx context.Context, in *ProcessDemoReq, opts ...grpc.CallOption) (*ProcessDemoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProcessDemoResponse)
	err := c.cc.Invoke(ctx, Stats_ProcessDemo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) ReplayDemo(ctx context.Context, in *ProcessDemoReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Pos], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Stats_ServiceDesc.Streams[0], Stats_ReplayDemo_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ProcessDemoReq, Pos]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Stats_ReplayDemoClient = grpc.ServerStreamingClient[Pos]

// StatsServer is the server API for Stats service.
// All implementations must embed UnimplementedStatsServer
// for forward compatibility.
type StatsServer interface {
	ProcessDemo(context.Context, *ProcessDemoReq) (*ProcessDemoResponse, error)
	ReplayDemo(*ProcessDemoReq, grpc.ServerStreamingServer[Pos]) error
	mustEmbedUnimplementedStatsServer()
}

// UnimplementedStatsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStatsServer struct{}

func (UnimplementedStatsServer) ProcessDemo(context.Context, *ProcessDemoReq) (*ProcessDemoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessDemo not implemented")
}
func (UnimplementedStatsServer) ReplayDemo(*ProcessDemoReq, grpc.ServerStreamingServer[Pos]) error {
	return status.Errorf(codes.Unimplemented, "method ReplayDemo not implemented")
}
func (UnimplementedStatsServer) mustEmbedUnimplementedStatsServer() {}
func (UnimplementedStatsServer) testEmbeddedByValue()               {}

// UnsafeStatsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatsServer will
// result in compilation errors.
type UnsafeStatsServer interface {
	mustEmbedUnimplementedStatsServer()
}

func RegisterStatsServer(s grpc.ServiceRegistrar, srv StatsServer) {
	// If the following call pancis, it indicates UnimplementedStatsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Stats_ServiceDesc, srv)
}

func _Stats_ProcessDemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessDemoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).ProcessDemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stats_ProcessDemo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).ProcessDemo(ctx, req.(*ProcessDemoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_ReplayDemo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProcessDemoReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StatsServer).ReplayDemo(m, &grpc.GenericServerStream[ProcessDemoReq, Pos]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Stats_ReplayDemoServer = grpc.ServerStreamingServer[Pos]

// Stats_ServiceDesc is the grpc.ServiceDesc for Stats service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stats_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Stats",
	HandlerType: (*StatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessDemo",
			Handler:    _Stats_ProcessDemo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReplayDemo",
			Handler:       _Stats_ReplayDemo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "demo.proto",
}
