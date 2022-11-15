package sit

import (
	"cd_platform/common"
	"cd_platform/mid"
	"cd_platform/pkg"
	"context"
)

type SitService interface {
	CreateSitDeployment(ctx context.Context, application string, raw []byte) error
	CreateSitService(ctx context.Context, application string, raw []byte) error
	CreateSitStatefulset(ctx context.Context, application string, raw []byte) error
	CreateSitNamespace(ctx context.Context, application string) error
	DeleteSitDeployment(ctx context.Context, application string) error
	DeleteSitService(ctx context.Context, application string) error
	DeleteSitStatefulset(ctx context.Context, application string) error
	DeleteSitNamespace(ctx context.Context, application string) error
	InsertSitIngressRule(ctx context.Context, application string, rule *common.IngressRule) error
	DeleteSitIngressRule(ctx context.Context, application string) error
	UpdateSitDeployment(ctx context.Context, application string, image string) error
}

type Service struct {
	Exec pkg.ExecService
}

func NewService(mid *mid.Middle) *Service {
	return &Service{
		Exec: pkg.ExService,
	}
}
