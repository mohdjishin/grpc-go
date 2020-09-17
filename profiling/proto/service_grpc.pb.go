// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ProfilingClient is the client API for Profiling service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfilingClient interface {
	// Enable allows users to toggle profiling on and off remotely.
	Enable(ctx context.Context, in *EnableRequest, opts ...grpc.CallOption) (*EnableResponse, error)
	// GetStreamStats is used to retrieve an array of stream-level stats from a
	// gRPC client/server.
	GetStreamStats(ctx context.Context, in *GetStreamStatsRequest, opts ...grpc.CallOption) (*GetStreamStatsResponse, error)
}

type profilingClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilingClient(cc grpc.ClientConnInterface) ProfilingClient {
	return &profilingClient{cc}
}

var profilingEnableStreamDesc = &grpc.StreamDesc{
	StreamName: "Enable",
}

func (c *profilingClient) Enable(ctx context.Context, in *EnableRequest, opts ...grpc.CallOption) (*EnableResponse, error) {
	out := new(EnableResponse)
	err := c.cc.Invoke(ctx, "/grpc.go.profiling.v1alpha.Profiling/Enable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var profilingGetStreamStatsStreamDesc = &grpc.StreamDesc{
	StreamName: "GetStreamStats",
}

func (c *profilingClient) GetStreamStats(ctx context.Context, in *GetStreamStatsRequest, opts ...grpc.CallOption) (*GetStreamStatsResponse, error) {
	out := new(GetStreamStatsResponse)
	err := c.cc.Invoke(ctx, "/grpc.go.profiling.v1alpha.Profiling/GetStreamStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilingService is the service API for Profiling service.
// Fields should be assigned to their respective handler implementations only before
// RegisterProfilingService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type ProfilingService struct {
	// Enable allows users to toggle profiling on and off remotely.
	Enable func(context.Context, *EnableRequest) (*EnableResponse, error)
	// GetStreamStats is used to retrieve an array of stream-level stats from a
	// gRPC client/server.
	GetStreamStats func(context.Context, *GetStreamStatsRequest) (*GetStreamStatsResponse, error)
}

func (s *ProfilingService) enable(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/grpc.go.profiling.v1alpha.Profiling/Enable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Enable(ctx, req.(*EnableRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *ProfilingService) getStreamStats(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetStreamStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/grpc.go.profiling.v1alpha.Profiling/GetStreamStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetStreamStats(ctx, req.(*GetStreamStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterProfilingService registers a service implementation with a gRPC server.
func RegisterProfilingService(s grpc.ServiceRegistrar, srv *ProfilingService) {
	srvCopy := *srv
	if srvCopy.Enable == nil {
		srvCopy.Enable = func(context.Context, *EnableRequest) (*EnableResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
		}
	}
	if srvCopy.GetStreamStats == nil {
		srvCopy.GetStreamStats = func(context.Context, *GetStreamStatsRequest) (*GetStreamStatsResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetStreamStats not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "grpc.go.profiling.v1alpha.Profiling",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Enable",
				Handler:    srvCopy.enable,
			},
			{
				MethodName: "GetStreamStats",
				Handler:    srvCopy.getStreamStats,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "profiling/proto/service.proto",
	}

	s.RegisterService(&sd, nil)
}

// ProfilingServer is the service API for Profiling service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended unless you own the service definition.
type ProfilingServer interface {
	// Enable allows users to toggle profiling on and off remotely.
	Enable(context.Context, *EnableRequest) (*EnableResponse, error)
	// GetStreamStats is used to retrieve an array of stream-level stats from a
	// gRPC client/server.
	GetStreamStats(context.Context, *GetStreamStatsRequest) (*GetStreamStatsResponse, error)
}

// UnimplementedProfilingServer can be embedded to have forward compatible implementations of
// ProfilingServer
type UnimplementedProfilingServer struct {
}

func (*UnimplementedProfilingServer) Enable(context.Context, *EnableRequest) (*EnableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (*UnimplementedProfilingServer) GetStreamStats(context.Context, *GetStreamStatsRequest) (*GetStreamStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreamStats not implemented")
}

// RegisterProfilingServer registers a service implementation with a gRPC server.
func RegisterProfilingServer(s grpc.ServiceRegistrar, srv ProfilingServer) {
	str := &ProfilingService{
		Enable:         srv.Enable,
		GetStreamStats: srv.GetStreamStats,
	}
	RegisterProfilingService(s, str)
}