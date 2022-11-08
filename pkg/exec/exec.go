package exec

import (
	"cd_platform/mid"
	"context"
)

type ExecService interface {
	CreateDeployment(ctx context.Context, project string, raw []byte) error
	CreateService(ctx context.Context, project string, raw []byte) error
	CreateStatefulset(ctx context.Context, project string, raw []byte) error
	CreateNamespace(ctx context.Context, project string, raw []byte) error
}

type Service struct {
	Mid *mid.Middle
}

func NewService(mid *mid.Middle) *Service {
	return &Service{
		Mid: mid,
	}
}
