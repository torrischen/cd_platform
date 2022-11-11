package pkg

import (
	"cd_platform/common"
	"cd_platform/mid"
	"context"
)

var ExService *Service

type ExecService interface {
	CreateDeployment(ctx context.Context, project string, raw []byte) error
	CreateService(ctx context.Context, project string, raw []byte) error
	CreateStatefulset(ctx context.Context, project string, raw []byte) error
	CreateNamespace(ctx context.Context, project string) error
	DeleteDeployment(ctx context.Context, project string) error
	DeleteService(ctx context.Context, project string) error
	DeleteStatefulset(ctx context.Context, project string) error
	DeleteNamespace(ctx context.Context, project string) error
	InsertIngressRule(ctx context.Context, rule *common.IngressRule) error
}

type Service struct {
	Mid *mid.Middle
}

func NewService(mid *mid.Middle) {
	ExService = &Service{
		Mid: mid,
	}
}
