package sit

import (
	"cd_platform/util"
	"context"
)

func (s *Service) CreateSitService(ctx context.Context, project string, raw []byte) error {
	return s.Exec.CreateService(ctx, util.ProjectToSit(project), raw)
}

func (s *Service) DeleteSitService(ctx context.Context, project string) error {
	return s.Exec.DeleteService(ctx, util.ProjectToSit(project))
}
