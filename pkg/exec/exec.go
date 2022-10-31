package exec

import (
	"cd_platform/common"
	"cd_platform/mid"
	"context"
	appsv1 "k8s.io/api/apps/v1"
)

type ExecService interface {
	CreateDeployment(ctx context.Context, args *common.NewDeploymentArgs) (*appsv1.Deployment, error)
}

type Service struct {
	Mid *mid.Middle
}

func NewService(mid *mid.Middle) *Service {
	return &Service{
		Mid: mid,
	}
}
