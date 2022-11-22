package sit

import (
	"cd_platform/util"
	"context"
)

func (s *Service) CreateSitNamespace(ctx context.Context, project string, application string) error {
	return s.Exec.CreateNamespace(ctx, util.ToSit(project)+"-"+application)
}

func (s *Service) DeleteSitNamespace(ctx context.Context, project string, application string) error {
	return s.Exec.DeleteNamespace(ctx, util.ToSit(project)+"-"+application)
}
