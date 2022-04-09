package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/pkg/errors"
	"gitlab.com/netbook-devs/spawner-service/pkg/service/common"
	"gitlab.com/netbook-devs/spawner-service/pkg/service/constants"
	"gitlab.com/netbook-devs/spawner-service/pkg/service/labels"
	proto "gitlab.com/netbook-devs/spawner-service/proto/netbookdevs/spawnerservice"
)

func (svc AWSController) createClusterInternal(ctx context.Context, session *Session, clusterName string, req *proto.ClusterRequest) (*eks.Cluster, error) {

	var subnetIds []*string
	region := session.Region

	awsRegionNetworkStack, err := GetRegionWkspNetworkStack(session)
	if err != nil {
		svc.logger.Errorw("error getting network stack for region", "region", region, "error", err)
		return nil, err
	}

	if awsRegionNetworkStack.Vpc != nil && len(awsRegionNetworkStack.Subnets) > 0 {
		for _, subn := range awsRegionNetworkStack.Subnets {
			subnetIds = append(subnetIds, subn.SubnetId)
		}
		svc.logger.Infow("got network stack for region", "vpc", awsRegionNetworkStack.Vpc.VpcId, "subnets", subnetIds)
	} else {
		awsRegionNetworkStack, err = CreateRegionWkspNetworkStack(session)
		if err != nil {
			svc.logger.Errorw("error creating network stack for region with no clusters", "region", region, "error", err)
			svc.logger.Warnw("rolling back network stack changes as creation failed", "region", region)
			delErr := DeleteRegionWkspNetworkStack(session, *awsRegionNetworkStack)
			if delErr != nil {
				svc.logger.Errorw("error deleting network stack for region", "region", region, "error", delErr)
			}

			return nil, err
		}
		for _, subn := range awsRegionNetworkStack.Subnets {
			subnetIds = append(subnetIds, subn.SubnetId)
		}
		svc.logger.Infow("created network stack for region", "vpc", awsRegionNetworkStack.Vpc.VpcId, "subnets", subnetIds)
	}

	tags := labels.DefaultTags()
	tags[constants.ClusterNameLabel] = &clusterName
	//override with additional labels from request
	for k, v := range req.Labels {
		v := v
		tags[k] = &v
	}

	iamClient := session.getIAMClient()
	roleName := AWS_CLUSTER_ROLE_NAME

	eksRole, newRole, err := svc.createRoleOrGetExisting(ctx, iamClient, roleName, "eks cluster and service access role", EKS_ASSUME_ROLE_DOC)

	if err != nil {
		svc.logger.Errorf("failed to create role %w", err)
		return nil, err
	}

	if newRole {
		err = svc.attachPolicy(ctx, iamClient, roleName, EKS_CLUSTER_POLICY_ARN)
		if err != nil {
			svc.logger.Errorf("failed to attach policy '%s' to role '%s' %w", EKS_CLUSTER_POLICY_ARN, roleName, err)
			return nil, err
		}

		err = svc.attachPolicy(ctx, iamClient, roleName, EKS_SERVICE_POLICY_ARN)
		if err != nil {
			svc.logger.Errorf("failed to attach policy '%s' to role '%s' %w", EKS_SERVICE_POLICY_ARN, roleName, err)
			return nil, err
		}
	}

	clusterInput := &eks.CreateClusterInput{
		Name: &clusterName,
		ResourcesVpcConfig: &eks.VpcConfigRequest{
			SubnetIds:             subnetIds,
			EndpointPublicAccess:  common.BoolPtr(true),
			EndpointPrivateAccess: common.BoolPtr(false),
		},
		Tags:    tags,
		Version: &constants.KubeVersion,
		RoleArn: eksRole.Arn,
	}

	client := session.getEksClient()
	createClusterOutput, err := client.CreateClusterWithContext(ctx, clusterInput)
	if err != nil {
		svc.logger.Errorw("failed to create cluster", "error", err)
		return nil, err
	}

	return createClusterOutput.Cluster, nil

}

//isExist check if the given cluster exist in EKS
//
// This function currently uses non paginated version of ListClusters, safe to assume that we would not have clusters more than 100 in single account.
func isExist(ctx context.Context, client *eks.EKS, name string) (bool, error) {

	res, err := client.ListClustersWithContext(ctx, &eks.ListClustersInput{})
	if err != nil {
		return false, err
	}

	for _, cluster := range res.Clusters {
		if name == *cluster {
			return true, nil
		}
	}
	return false, nil
}

//CreateCluster Create new cluster with given specification, no op if cluster already exist
func (ctrl AWSController) CreateCluster(ctx context.Context, req *proto.ClusterRequest) (*proto.ClusterResponse, error) {

	var clusterName string
	if clusterName = req.ClusterName; len(clusterName) == 0 {
		clusterName = fmt.Sprintf("%s-%s", req.Provider, req.Region)
	}

	region := req.Region
	accountName := req.AccountName
	session, err := NewSession(ctx, region, accountName)

	if err != nil {
		return nil, err
	}
	eksClient := session.getEksClient()

	ctrl.logger.Debugf("checking cluster status for '%s', region '%s'", clusterName, region)
	exist, err := isExist(ctx, eksClient, clusterName)

	if err != nil {
		return nil, errors.Wrap(err, "CreateCluster failed")
	}

	if exist {
		ctrl.logger.Infof("cluster '%s', already exist", clusterName)
		return nil, errors.New("cluster already exist")
	}

	ctrl.logger.Debugf("cluster '%s' does not exist, creating ...", clusterName)
	cluster, err := ctrl.createClusterInternal(ctx, session, clusterName, req)
	if err != nil {
		ctrl.logger.Errorw("failed to create cluster ", "cluster", clusterName, "error", err)
		return nil, err
	}

	ctrl.logger.Infow("cluster is in creating state, it might take some time, please check AWS console for status", "cluster", clusterName)

	return &proto.ClusterResponse{
		ClusterName: *cluster.Name,
	}, nil
}
