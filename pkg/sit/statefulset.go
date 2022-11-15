package sit

import (
	"cd_platform/util"
	"context"
)

func (s *Service) CreateSitStatefulset(ctx context.Context, application string, raw []byte) error {
	return s.Exec.CreateStatefulset(ctx, util.ToSit(application), raw)
}

func (s *Service) DeleteSitStatefulset(ctx context.Context, application string) error {
	return s.Exec.DeleteStatefulset(ctx, util.ToSit(application), application)
}
