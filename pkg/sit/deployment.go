package sit

import (
	"cd_platform/util"
	"context"
)

func (s *Service) CreateSitDeployment(ctx context.Context, project string, raw []byte) error {
	return s.Exec.CreateDeployment(ctx, util.ProjectToSit(project), raw)
}

func (s *Service) DeleteSitDeployment(ctx context.Context, project string, application string) error {
	return s.Exec.DeleteDeployment(ctx, util.ProjectToSit(project), application)
}
