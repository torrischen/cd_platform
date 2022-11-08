package sit

import (
	"cd_platform/mid"
	"context"
)

type SitService interface {
	CreateDeployment(ctx context.Context, project string, raw []byte) error
}

type Service struct {
	Mid *mid.Middle
}

func NewService(mid *mid.Middle) *Service {
	return &Service{
		Mid: mid,
	}
}
