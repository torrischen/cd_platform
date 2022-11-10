package sit

import (
	"cd_platform/util"
	"context"
)

func (s *Service) CreateSitNamespace(ctx context.Context, project string) error {
	return s.Exec.CreateNamespace(ctx, util.ProjectToSit(project))
}

func (s *Service) DeleteSitNamespace(ctx context.Context, project string) error {
	return s.Exec.DeleteNamespace(ctx, util.ProjectToSit(project))
}
