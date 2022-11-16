package sit

import (
	"cd_platform/util"
	"context"
	corev1 "k8s.io/api/core/v1"
)

func (s *Service) CreateSitService(ctx context.Context, application string, service *corev1.Service) error {
	return s.Exec.CreateService(ctx, util.ToSit(application), service)
}

func (s *Service) DeleteSitService(ctx context.Context, application string) error {
	return s.Exec.DeleteService(ctx, util.ToSit(application), application)
}
