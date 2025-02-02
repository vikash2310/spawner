package service

import (
	"context"

	proto "gitlab.com/netbook-devs/spawner-service/proto/netbookai/spawner"
)

//Controller defines the set of the functions which each provider controller must implement to support spawner functionality
type Controller interface {
	CreateCluster(ctx context.Context, req *proto.ClusterRequest) (*proto.ClusterResponse, error)
	GetCluster(ctx context.Context, req *proto.GetClusterRequest) (*proto.ClusterSpec, error)
	GetClusters(ctx context.Context, req *proto.GetClustersRequest) (*proto.GetClustersResponse, error)
	AddToken(ctx context.Context, req *proto.AddTokenRequest) (*proto.AddTokenResponse, error)
	GetToken(ctx context.Context, req *proto.GetTokenRequest) (*proto.GetTokenResponse, error)
	ClusterStatus(ctx context.Context, req *proto.ClusterStatusRequest) (*proto.ClusterStatusResponse, error)
	AddNode(ctx context.Context, req *proto.NodeSpawnRequest) (*proto.NodeSpawnResponse, error)
	DeleteCluster(ctx context.Context, req *proto.ClusterDeleteRequest) (*proto.ClusterDeleteResponse, error)
	DeleteNode(ctx context.Context, req *proto.NodeDeleteRequest) (*proto.NodeDeleteResponse, error)
	CreateVolume(ctx context.Context, req *proto.CreateVolumeRequest) (*proto.CreateVolumeResponse, error)
	DeleteVolume(ctx context.Context, req *proto.DeleteVolumeRequest) (*proto.DeleteVolumeResponse, error)
	CreateSnapshot(ctx context.Context, req *proto.CreateSnapshotRequest) (*proto.CreateSnapshotResponse, error)
	CreateSnapshotAndDelete(ctx context.Context, req *proto.CreateSnapshotAndDeleteRequest) (*proto.CreateSnapshotAndDeleteResponse, error)
	GetWorkspacesCost(context.Context, *proto.GetWorkspacesCostRequest) (*proto.GetWorkspacesCostResponse, error)
	GetKubeConfig(ctx context.Context, in *proto.GetKubeConfigRequest) (*proto.GetKubeConfigResponse, error)
	TagNodeInstance(ctx context.Context, req *proto.TagNodeInstanceRequest) (*proto.TagNodeInstanceResponse, error)
	GetCostByTime(ctx context.Context, req *proto.GetCostByTimeRequest) (*proto.GetCostByTimeResponse, error)
}
