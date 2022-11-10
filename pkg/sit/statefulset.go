package sit

import (
	"cd_platform/util"
	"context"
)

func (s *Service) CreateSitStatefulset(ctx context.Context, project string, raw []byte) error {
	return s.Exec.CreateStatefulset(ctx, util.ProjectToSit(project), raw)
}

func (s *Service) DeleteSitStatefulset(ctx context.Context, project string) error {
	return s.Exec.DeleteStatefulset(ctx, util.ProjectToSit(project))
}
