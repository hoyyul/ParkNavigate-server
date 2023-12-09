// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.0--rc2
// source: navigation.proto

package navigation_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NavigationServiceClient is the client API for NavigationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NavigationServiceClient interface {
	ShowParkings(ctx context.Context, in *NavigationRequest, opts ...grpc.CallOption) (*NavigationResponse, error)
}

type navigationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNavigationServiceClient(cc grpc.ClientConnInterface) NavigationServiceClient {
	return &navigationServiceClient{cc}
}

func (c *navigationServiceClient) ShowParkings(ctx context.Context, in *NavigationRequest, opts ...grpc.CallOption) (*NavigationResponse, error) {
	out := new(NavigationResponse)
	err := c.cc.Invoke(ctx, "/navigation_pb.NavigationService/ShowParkings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NavigationServiceServer is the server API for NavigationService service.
// All implementations must embed UnimplementedNavigationServiceServer
// for forward compatibility
type NavigationServiceServer interface {
	ShowParkings(context.Context, *NavigationRequest) (*NavigationResponse, error)
	mustEmbedUnimplementedNavigationServiceServer()
}

// UnimplementedNavigationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNavigationServiceServer struct {
}

func (UnimplementedNavigationServiceServer) ShowParkings(context.Context, *NavigationRequest) (*NavigationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowParkings not implemented")
}
func (UnimplementedNavigationServiceServer) mustEmbedUnimplementedNavigationServiceServer() {}

// UnsafeNavigationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NavigationServiceServer will
// result in compilation errors.
type UnsafeNavigationServiceServer interface {
	mustEmbedUnimplementedNavigationServiceServer()
}

func RegisterNavigationServiceServer(s grpc.ServiceRegistrar, srv NavigationServiceServer) {
	s.RegisterService(&NavigationService_ServiceDesc, srv)
}

func _NavigationService_ShowParkings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NavigationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NavigationServiceServer).ShowParkings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/navigation_pb.NavigationService/ShowParkings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NavigationServiceServer).ShowParkings(ctx, req.(*NavigationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NavigationService_ServiceDesc is the grpc.ServiceDesc for NavigationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NavigationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "navigation_pb.NavigationService",
	HandlerType: (*NavigationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShowParkings",
			Handler:    _NavigationService_ShowParkings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "navigation.proto",
}
