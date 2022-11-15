package sit

import (
	"cd_platform/util"
	"context"
)

func (s *Service) CreateSitNamespace(ctx context.Context, application string) error {
	return s.Exec.CreateNamespace(ctx, util.ToSit(application))
}

func (s *Service) DeleteSitNamespace(ctx context.Context, application string) error {
	return s.Exec.DeleteNamespace(ctx, util.ToSit(application))
}
