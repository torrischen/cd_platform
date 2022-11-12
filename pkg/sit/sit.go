package sit

import (
	"cd_platform/mid"
	"cd_platform/pkg"
	"context"
)

type SitService interface {
	CreateSitDeployment(ctx context.Context, project string, raw []byte) error
	CreateSitService(ctx context.Context, project string, raw []byte) error
	CreateSitStatefulset(ctx context.Context, project string, raw []byte) error
	CreateSitNamespace(ctx context.Context, project string) error
	DeleteSitDeployment(ctx context.Context, project string, application string) error
	DeleteSitService(ctx context.Context, project string, application string) error
	DeleteSitStatefulset(ctx context.Context, project string, application string) error
	DeleteSitNamespace(ctx context.Context, project string) error
}

type Service struct {
	Exec pkg.ExecService
}

func NewService(mid *mid.Middle) *Service {
	return &Service{
		Exec: pkg.ExService,
	}
}
