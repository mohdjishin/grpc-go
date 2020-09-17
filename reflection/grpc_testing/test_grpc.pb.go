// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc_testing

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	StreamingSearch(ctx context.Context, opts ...grpc.CallOption) (SearchService_StreamingSearchClient, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

var searchServiceSearchStreamDesc = &grpc.StreamDesc{
	StreamName: "Search",
}

func (c *searchServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/grpc.testing.SearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var searchServiceStreamingSearchStreamDesc = &grpc.StreamDesc{
	StreamName:    "StreamingSearch",
	ServerStreams: true,
	ClientStreams: true,
}

func (c *searchServiceClient) StreamingSearch(ctx context.Context, opts ...grpc.CallOption) (SearchService_StreamingSearchClient, error) {
	stream, err := c.cc.NewStream(ctx, searchServiceStreamingSearchStreamDesc, "/grpc.testing.SearchService/StreamingSearch", opts...)
	if err != nil {
		return nil, err
	}
	x := &searchServiceStreamingSearchClient{stream}
	return x, nil
}

type SearchService_StreamingSearchClient interface {
	Send(*SearchRequest) error
	Recv() (*SearchResponse, error)
	grpc.ClientStream
}

type searchServiceStreamingSearchClient struct {
	grpc.ClientStream
}

func (x *searchServiceStreamingSearchClient) Send(m *SearchRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *searchServiceStreamingSearchClient) Recv() (*SearchResponse, error) {
	m := new(SearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SearchServiceService is the service API for SearchService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterSearchServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type SearchServiceService struct {
	Search          func(context.Context, *SearchRequest) (*SearchResponse, error)
	StreamingSearch func(SearchService_StreamingSearchServer) error
}

func (s *SearchServiceService) search(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/grpc.testing.SearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SearchServiceService) streamingSearch(_ interface{}, stream grpc.ServerStream) error {
	return s.StreamingSearch(&searchServiceStreamingSearchServer{stream})
}

type SearchService_StreamingSearchServer interface {
	Send(*SearchResponse) error
	Recv() (*SearchRequest, error)
	grpc.ServerStream
}

type searchServiceStreamingSearchServer struct {
	grpc.ServerStream
}

func (x *searchServiceStreamingSearchServer) Send(m *SearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *searchServiceStreamingSearchServer) Recv() (*SearchRequest, error) {
	m := new(SearchRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegisterSearchServiceService registers a service implementation with a gRPC server.
func RegisterSearchServiceService(s grpc.ServiceRegistrar, srv *SearchServiceService) {
	srvCopy := *srv
	if srvCopy.Search == nil {
		srvCopy.Search = func(context.Context, *SearchRequest) (*SearchResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
		}
	}
	if srvCopy.StreamingSearch == nil {
		srvCopy.StreamingSearch = func(SearchService_StreamingSearchServer) error {
			return status.Errorf(codes.Unimplemented, "method StreamingSearch not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "grpc.testing.SearchService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Search",
				Handler:    srvCopy.search,
			},
		},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "StreamingSearch",
				Handler:       srvCopy.streamingSearch,
				ServerStreams: true,
				ClientStreams: true,
			},
		},
		Metadata: "reflection/grpc_testing/test.proto",
	}

	s.RegisterService(&sd, nil)
}