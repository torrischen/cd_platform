package sit

import (
	"cd_platform/util"
	"context"
)

func (s *Service) CreateSitDeployment(ctx context.Context, application string, raw []byte) error {
	return s.Exec.CreateDeployment(ctx, util.ToSit(application), raw)
}

func (s *Service) DeleteSitDeployment(ctx context.Context, application string) error {
	return s.Exec.DeleteDeployment(ctx, util.ToSit(application), application)
}

func (s *Service) UpdateSitDeployment(ctx context.Context, application string, image string) error {
	return s.Exec.UpdateDeployment(ctx, util.ToSit(application), util.ToSit(application), image)
}
