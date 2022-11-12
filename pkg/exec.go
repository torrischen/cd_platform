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
	DeleteDeployment(ctx context.Context, project string, application string) error
	DeleteService(ctx context.Context, project string, application string) error
	DeleteStatefulset(ctx context.Context, project string, application string) error
	DeleteNamespace(ctx context.Context, project string) error
	InsertIngressRule(ctx context.Context, project string, rule *common.IngressRule) error
	DeleteIngressRule(ctx context.Context, project string, application string) error
	UpdateDeployment(ctx context.Context, project string, application string, image string) error
}

type Service struct {
	Mid *mid.Middle
}

func NewService(mid *mid.Middle) {
	ExService = &Service{
		Mid: mid,
	}
}
