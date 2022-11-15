package sit

import (
	"cd_platform/util"
	"context"
)

func (s *Service) CreateSitService(ctx context.Context, application string, raw []byte) error {
	return s.Exec.CreateService(ctx, util.ToSit(application), raw)
}

func (s *Service) DeleteSitService(ctx context.Context, application string) error {
	return s.Exec.DeleteService(ctx, util.ToSit(application), application)
}
