// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/netbookdevs/spawnerservice/spawner.proto

package spawnerservice

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
	HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	// Spawn required cluster
	CreateCluster(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterResponse, error)
	// Create add token to secret manager
	AddToken(ctx context.Context, in *AddTokenRequest, opts ...grpc.CallOption) (*AddTokenResponse, error)
	// Create get token to secret manager
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
	// Add Route53 record for Caddy
	AddRoute53Record(ctx context.Context, in *AddRoute53RecordRequest, opts ...grpc.CallOption) (*AddRoute53RecordResponse, error)
	// Get Cluster
	GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*ClusterSpec, error)
	// Get Clusters
	GetClusters(ctx context.Context, in *GetClustersRequest, opts ...grpc.CallOption) (*GetClustersResponse, error)
	// Spawn required instance
	AddNode(ctx context.Context, in *NodeSpawnRequest, opts ...grpc.CallOption) (*NodeSpawnResponse, error)
	// Status of cluster
	ClusterStatus(ctx context.Context, in *ClusterStatusRequest, opts ...grpc.CallOption) (*ClusterStatusResponse, error)
	// Delete Cluster
	DeleteCluster(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*ClusterDeleteResponse, error)
	// Delete Node
	DeleteNode(ctx context.Context, in *NodeDeleteRequest, opts ...grpc.CallOption) (*NodeDeleteResponse, error)
	// Create Volume
	CreateVolume(ctx context.Context, in *CreateVolumeRequest, opts ...grpc.CallOption) (*CreateVolumeResponse, error)
	// Delete Vol
	DeleteVolume(ctx context.Context, in *DeleteVolumeRequest, opts ...grpc.CallOption) (*DeleteVolumeResponse, error)
	CreateSnapshot(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*CreateSnapshotResponse, error)
	CreateSnapshotAndDelete(ctx context.Context, in *CreateSnapshotAndDeleteRequest, opts ...grpc.CallOption) (*CreateSnapshotAndDeleteResponse, error)
	RegisterWithRancher(ctx context.Context, in *RancherRegistrationRequest, opts ...grpc.CallOption) (*RancherRegistrationResponse, error)
	GetWorkspacesCost(ctx context.Context, in *GetWorkspacesCostRequest, opts ...grpc.CallOption) (*GetWorkspacesCostResponse, error)
	WriteCredential(ctx context.Context, in *WriteCredentialRequest, opts ...grpc.CallOption) (*WriteCredentialResponse, error)
	ReadCredential(ctx context.Context, in *ReadCredentialRequest, opts ...grpc.CallOption) (*ReadCredentialResponse, error)
}

type spawnerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpawnerServiceClient(cc grpc.ClientConnInterface) SpawnerServiceClient {
	return &spawnerServiceClient{cc}
}

func (c *spawnerServiceClient) HealthCheck(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) CreateCluster(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterResponse, error) {
	out := new(ClusterResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/CreateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) AddToken(ctx context.Context, in *AddTokenRequest, opts ...grpc.CallOption) (*AddTokenResponse, error) {
	out := new(AddTokenResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/AddToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) AddRoute53Record(ctx context.Context, in *AddRoute53RecordRequest, opts ...grpc.CallOption) (*AddRoute53RecordResponse, error) {
	out := new(AddRoute53RecordResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/AddRoute53Record", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*ClusterSpec, error) {
	out := new(ClusterSpec)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) GetClusters(ctx context.Context, in *GetClustersRequest, opts ...grpc.CallOption) (*GetClustersResponse, error) {
	out := new(GetClustersResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/GetClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) AddNode(ctx context.Context, in *NodeSpawnRequest, opts ...grpc.CallOption) (*NodeSpawnResponse, error) {
	out := new(NodeSpawnResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/AddNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) ClusterStatus(ctx context.Context, in *ClusterStatusRequest, opts ...grpc.CallOption) (*ClusterStatusResponse, error) {
	out := new(ClusterStatusResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/ClusterStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) DeleteCluster(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*ClusterDeleteResponse, error) {
	out := new(ClusterDeleteResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/DeleteCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) DeleteNode(ctx context.Context, in *NodeDeleteRequest, opts ...grpc.CallOption) (*NodeDeleteResponse, error) {
	out := new(NodeDeleteResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/DeleteNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) CreateVolume(ctx context.Context, in *CreateVolumeRequest, opts ...grpc.CallOption) (*CreateVolumeResponse, error) {
	out := new(CreateVolumeResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/CreateVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) DeleteVolume(ctx context.Context, in *DeleteVolumeRequest, opts ...grpc.CallOption) (*DeleteVolumeResponse, error) {
	out := new(DeleteVolumeResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/DeleteVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) CreateSnapshot(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*CreateSnapshotResponse, error) {
	out := new(CreateSnapshotResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/CreateSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) CreateSnapshotAndDelete(ctx context.Context, in *CreateSnapshotAndDeleteRequest, opts ...grpc.CallOption) (*CreateSnapshotAndDeleteResponse, error) {
	out := new(CreateSnapshotAndDeleteResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/CreateSnapshotAndDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) RegisterWithRancher(ctx context.Context, in *RancherRegistrationRequest, opts ...grpc.CallOption) (*RancherRegistrationResponse, error) {
	out := new(RancherRegistrationResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/RegisterWithRancher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) GetWorkspacesCost(ctx context.Context, in *GetWorkspacesCostRequest, opts ...grpc.CallOption) (*GetWorkspacesCostResponse, error) {
	out := new(GetWorkspacesCostResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/GetWorkspacesCost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) WriteCredential(ctx context.Context, in *WriteCredentialRequest, opts ...grpc.CallOption) (*WriteCredentialResponse, error) {
	out := new(WriteCredentialResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/WriteCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnerServiceClient) ReadCredential(ctx context.Context, in *ReadCredentialRequest, opts ...grpc.CallOption) (*ReadCredentialResponse, error) {
	out := new(ReadCredentialResponse)
	err := c.cc.Invoke(ctx, "/spawnerservice.SpawnerService/ReadCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpawnerServiceServer is the server API for SpawnerService service.
// All implementations must embed UnimplementedSpawnerServiceServer
// for forward compatibility
type SpawnerServiceServer interface {
	HealthCheck(context.Context, *Empty) (*Empty, error)
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	// Spawn required cluster
	CreateCluster(context.Context, *ClusterRequest) (*ClusterResponse, error)
	// Create add token to secret manager
	AddToken(context.Context, *AddTokenRequest) (*AddTokenResponse, error)
	// Create get token to secret manager
	GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error)
	// Add Route53 record for Caddy
	AddRoute53Record(context.Context, *AddRoute53RecordRequest) (*AddRoute53RecordResponse, error)
	// Get Cluster
	GetCluster(context.Context, *GetClusterRequest) (*ClusterSpec, error)
	// Get Clusters
	GetClusters(context.Context, *GetClustersRequest) (*GetClustersResponse, error)
	// Spawn required instance
	AddNode(context.Context, *NodeSpawnRequest) (*NodeSpawnResponse, error)
	// Status of cluster
	ClusterStatus(context.Context, *ClusterStatusRequest) (*ClusterStatusResponse, error)
	// Delete Cluster
	DeleteCluster(context.Context, *ClusterDeleteRequest) (*ClusterDeleteResponse, error)
	// Delete Node
	DeleteNode(context.Context, *NodeDeleteRequest) (*NodeDeleteResponse, error)
	// Create Volume
	CreateVolume(context.Context, *CreateVolumeRequest) (*CreateVolumeResponse, error)
	// Delete Vol
	DeleteVolume(context.Context, *DeleteVolumeRequest) (*DeleteVolumeResponse, error)
	CreateSnapshot(context.Context, *CreateSnapshotRequest) (*CreateSnapshotResponse, error)
	CreateSnapshotAndDelete(context.Context, *CreateSnapshotAndDeleteRequest) (*CreateSnapshotAndDeleteResponse, error)
	RegisterWithRancher(context.Context, *RancherRegistrationRequest) (*RancherRegistrationResponse, error)
	GetWorkspacesCost(context.Context, *GetWorkspacesCostRequest) (*GetWorkspacesCostResponse, error)
	WriteCredential(context.Context, *WriteCredentialRequest) (*WriteCredentialResponse, error)
	ReadCredential(context.Context, *ReadCredentialRequest) (*ReadCredentialResponse, error)
	mustEmbedUnimplementedSpawnerServiceServer()
}

// UnimplementedSpawnerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSpawnerServiceServer struct {
}

func (UnimplementedSpawnerServiceServer) HealthCheck(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedSpawnerServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedSpawnerServiceServer) CreateCluster(context.Context, *ClusterRequest) (*ClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCluster not implemented")
}
func (UnimplementedSpawnerServiceServer) AddToken(context.Context, *AddTokenRequest) (*AddTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToken not implemented")
}
func (UnimplementedSpawnerServiceServer) GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedSpawnerServiceServer) AddRoute53Record(context.Context, *AddRoute53RecordRequest) (*AddRoute53RecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoute53Record not implemented")
}
func (UnimplementedSpawnerServiceServer) GetCluster(context.Context, *GetClusterRequest) (*ClusterSpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (UnimplementedSpawnerServiceServer) GetClusters(context.Context, *GetClustersRequest) (*GetClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusters not implemented")
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
func (UnimplementedSpawnerServiceServer) CreateVolume(context.Context, *CreateVolumeRequest) (*CreateVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVolume not implemented")
}
func (UnimplementedSpawnerServiceServer) DeleteVolume(context.Context, *DeleteVolumeRequest) (*DeleteVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVolume not implemented")
}
func (UnimplementedSpawnerServiceServer) CreateSnapshot(context.Context, *CreateSnapshotRequest) (*CreateSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSnapshot not implemented")
}
func (UnimplementedSpawnerServiceServer) CreateSnapshotAndDelete(context.Context, *CreateSnapshotAndDeleteRequest) (*CreateSnapshotAndDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSnapshotAndDelete not implemented")
}
func (UnimplementedSpawnerServiceServer) RegisterWithRancher(context.Context, *RancherRegistrationRequest) (*RancherRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWithRancher not implemented")
}
func (UnimplementedSpawnerServiceServer) GetWorkspacesCost(context.Context, *GetWorkspacesCostRequest) (*GetWorkspacesCostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspacesCost not implemented")
}
func (UnimplementedSpawnerServiceServer) WriteCredential(context.Context, *WriteCredentialRequest) (*WriteCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteCredential not implemented")
}
func (UnimplementedSpawnerServiceServer) ReadCredential(context.Context, *ReadCredentialRequest) (*ReadCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCredential not implemented")
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

func _SpawnerService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawnerservice.SpawnerService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).HealthCheck(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawnerservice.SpawnerService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/spawnerservice.SpawnerService/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).CreateCluster(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_AddToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).AddToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawnerservice.SpawnerService/AddToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).AddToken(ctx, req.(*AddTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawnerservice.SpawnerService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_AddRoute53Record_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoute53RecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).AddRoute53Record(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawnerservice.SpawnerService/AddRoute53Record",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).AddRoute53Record(ctx, req.(*AddRoute53RecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawnerservice.SpawnerService/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).GetCluster(ctx, req.(*GetClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_GetClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).GetClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawnerservice.SpawnerService/GetClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).GetClusters(ctx, req.(*GetClustersRequest))
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
		FullMethod: "/spawnerservice.SpawnerService/AddNode",
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
		FullMethod: "/spawnerservice.SpawnerService/ClusterStatus",
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
		FullMethod: "/spawnerservice.SpawnerService/DeleteCluster",
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
		FullMethod: "/spawnerservice.SpawnerService/DeleteNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).DeleteNode(ctx, req.(*NodeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_CreateVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).CreateVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawnerservice.SpawnerService/CreateVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).CreateVolume(ctx, req.(*CreateVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_DeleteVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).DeleteVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawnerservice.SpawnerService/DeleteVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).DeleteVolume(ctx, req.(*DeleteVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_CreateSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).CreateSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawnerservice.SpawnerService/CreateSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).CreateSnapshot(ctx, req.(*CreateSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_CreateSnapshotAndDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnapshotAndDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).CreateSnapshotAndDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawnerservice.SpawnerService/CreateSnapshotAndDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).CreateSnapshotAndDelete(ctx, req.(*CreateSnapshotAndDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_RegisterWithRancher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RancherRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).RegisterWithRancher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawnerservice.SpawnerService/RegisterWithRancher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).RegisterWithRancher(ctx, req.(*RancherRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_GetWorkspacesCost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspacesCostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).GetWorkspacesCost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawnerservice.SpawnerService/GetWorkspacesCost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).GetWorkspacesCost(ctx, req.(*GetWorkspacesCostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_WriteCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).WriteCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawnerservice.SpawnerService/WriteCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).WriteCredential(ctx, req.(*WriteCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpawnerService_ReadCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnerServiceServer).ReadCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawnerservice.SpawnerService/ReadCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnerServiceServer).ReadCredential(ctx, req.(*ReadCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpawnerService_ServiceDesc is the grpc.ServiceDesc for SpawnerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpawnerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spawnerservice.SpawnerService",
	HandlerType: (*SpawnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _SpawnerService_HealthCheck_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _SpawnerService_Echo_Handler,
		},
		{
			MethodName: "CreateCluster",
			Handler:    _SpawnerService_CreateCluster_Handler,
		},
		{
			MethodName: "AddToken",
			Handler:    _SpawnerService_AddToken_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _SpawnerService_GetToken_Handler,
		},
		{
			MethodName: "AddRoute53Record",
			Handler:    _SpawnerService_AddRoute53Record_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _SpawnerService_GetCluster_Handler,
		},
		{
			MethodName: "GetClusters",
			Handler:    _SpawnerService_GetClusters_Handler,
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
		{
			MethodName: "CreateVolume",
			Handler:    _SpawnerService_CreateVolume_Handler,
		},
		{
			MethodName: "DeleteVolume",
			Handler:    _SpawnerService_DeleteVolume_Handler,
		},
		{
			MethodName: "CreateSnapshot",
			Handler:    _SpawnerService_CreateSnapshot_Handler,
		},
		{
			MethodName: "CreateSnapshotAndDelete",
			Handler:    _SpawnerService_CreateSnapshotAndDelete_Handler,
		},
		{
			MethodName: "RegisterWithRancher",
			Handler:    _SpawnerService_RegisterWithRancher_Handler,
		},
		{
			MethodName: "GetWorkspacesCost",
			Handler:    _SpawnerService_GetWorkspacesCost_Handler,
		},
		{
			MethodName: "WriteCredential",
			Handler:    _SpawnerService_WriteCredential_Handler,
		},
		{
			MethodName: "ReadCredential",
			Handler:    _SpawnerService_ReadCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/netbookdevs/spawnerservice/spawner.proto",
}
