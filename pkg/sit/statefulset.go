package sit

import (
	"cd_platform/util"
	"context"
	appsv1 "k8s.io/api/apps/v1"
)

func (s *Service) CreateSitStatefulset(ctx context.Context, application string, sts *appsv1.StatefulSet) error {
	return s.Exec.CreateStatefulset(ctx, util.ToSit(application), sts)
}

func (s *Service) DeleteSitStatefulset(ctx context.Context, application string) error {
	return s.Exec.DeleteStatefulset(ctx, util.ToSit(application), application)
}
