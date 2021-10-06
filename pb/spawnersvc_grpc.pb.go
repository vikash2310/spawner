// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// SpawnerServiceClient is the client API for SpawnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpawnerServiceClient interface {
	// Spawn required cluster
	CreateCluster(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterResponse, error)
	// Spawn required instance
	AddNode(ctx context.Context, in *NodeSpawnRequest, opts ...grpc.CallOption) (*NodeSpawnResponse, error)
	// Status of cluster
	ClusterStatus(ctx context.Context, in *ClusterStatusRequest, opts ...grpc.CallOption) (*ClusterStatusResponse, error)
	// Delete Cluster
	DeleteCluster(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*ClusterDeleteResponse, error)
	// Delete Node
	DeleteNode(ctx context.Context, in *NodeDeleteRequest, opts ...grpc.CallOption) (*NodeDeleteResponse, error)
}

type spawnerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpawnerServiceClient(cc grpc.ClientConnInterface) SpawnerServiceClient {
	return &spawnerServiceClient{cc}
}

func (c *spawnerServiceClient) CreateCluster(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterResponse, error) {
	out := new(ClusterResponse)
	err := c.cc.Invoke(ctx, "/pb.SpawnerService/CreateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) AddNode(ctx context.Context, in *NodeSpawnRequest, opts ...grpc.CallOption) (*NodeSpawnResponse, error) {
	out := new(NodeSpawnResponse)
	err := c.cc.Invoke(ctx, "/pb.SpawnerService/AddNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) ClusterStatus(ctx context.Context, in *ClusterStatusRequest, opts ...grpc.CallOption) (*ClusterStatusResponse, error) {
	out := new(ClusterStatusResponse)
	err := c.cc.Invoke(ctx, "/pb.SpawnerService/ClusterStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) DeleteCluster(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*ClusterDeleteResponse, error) {
	out := new(ClusterDeleteResponse)
	err := c.cc.Invoke(ctx, "/pb.SpawnerService/DeleteCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) DeleteNode(ctx context.Context, in *NodeDeleteRequest, opts ...grpc.CallOption) (*NodeDeleteResponse, error) {
	out := new(NodeDeleteResponse)
	err := c.cc.Invoke(ctx, "/pb.SpawnerService/DeleteNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpawnerServiceServer is the server API for SpawnerService service.
// All implementations must embed UnimplementedSpawnerServiceServer
// for forward compatibility
type SpawnerServiceServer interface {
	// Spawn required cluster
	CreateCluster(context.Context, *ClusterRequest) (*ClusterResponse, error)
	// Spawn required instance
	AddNode(context.Context, *NodeSpawnRequest) (*NodeSpawnResponse, error)
	// Status of cluster
	ClusterStatus(context.Context, *ClusterStatusRequest) (*ClusterStatusResponse, error)
	// Delete Cluster
	DeleteCluster(context.Context, *ClusterDeleteRequest) (*ClusterDeleteResponse, error)
	// Delete Node
	DeleteNode(context.Context, *NodeDeleteRequest) (*NodeDeleteResponse, error)
	mustEmbedUnimplementedSpawnerServiceServer()
}

// UnimplementedSpawnerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSpawnerServiceServer struct {
}

func (UnimplementedSpawnerServiceServer) CreateCluster(context.Context, *ClusterRequest) (*ClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCluster not implemented")
}
func (UnimplementedSpawnerServiceServer) AddNode(context.Context, *NodeSpawnRequest) (*NodeSpawnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNode not implemented")
}
func (UnimplementedSpawnerServiceServer) ClusterStatus(context.Context, *ClusterStatusRequest) (*ClusterStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterStatus not implemented")
}
func (UnimplementedSpawnerServiceServer) DeleteCluster(context.Context, *ClusterDeleteRequest) (*ClusterDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCluster not implemented")
}
func (UnimplementedSpawnerServiceServer) DeleteNode(context.Context, *NodeDeleteRequest) (*NodeDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNode not implemented")
}
func (UnimplementedSpawnerServiceServer) mustEmbedUnimplementedSpawnerServiceServer() {}

// UnsafeSpawnerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpawnerServiceServer will
// result in compilation errors.
type UnsafeSpawnerServiceServer interface {
	mustEmbedUnimplementedSpawnerServiceServer()
}

func RegisterSpawnerServiceServer(s grpc.ServiceRegistrar, srv SpawnerServiceServer) {
	s.RegisterService(&SpawnerService_ServiceDesc, srv)
}

func _SpawnerService_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpawnerService/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).CreateCluster(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_AddNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeSpawnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).AddNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpawnerService/AddNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).AddNode(ctx, req.(*NodeSpawnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_ClusterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).ClusterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpawnerService/ClusterStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).ClusterStatus(ctx, req.(*ClusterStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).DeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpawnerService/DeleteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).DeleteCluster(ctx, req.(*ClusterDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_DeleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).DeleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SpawnerService/DeleteNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).DeleteNode(ctx, req.(*NodeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpawnerService_ServiceDesc is the grpc.ServiceDesc for SpawnerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpawnerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SpawnerService",
	HandlerType: (*SpawnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCluster",
			Handler:    _SpawnerService_CreateCluster_Handler,
		},
		{
			MethodName: "AddNode",
			Handler:    _SpawnerService_AddNode_Handler,
		},
		{
			MethodName: "ClusterStatus",
			Handler:    _SpawnerService_ClusterStatus_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _SpawnerService_DeleteCluster_Handler,
		},
		{
			MethodName: "DeleteNode",
			Handler:    _SpawnerService_DeleteNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spawnersvc.proto",
}
