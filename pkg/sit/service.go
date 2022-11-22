package sit

import (
	"cd_platform/util"
	"context"
	corev1 "k8s.io/api/core/v1"
)

func (s *Service) CreateSitService(ctx context.Context, project string, service *corev1.Service) error {
	return s.Exec.CreateService(ctx, util.ToSit(project)+"-"+service.Name, service)
}

func (s *Service) DeleteSitService(ctx context.Context, project string, application string) error {
	return s.Exec.DeleteService(ctx, util.ToSit(project)+"-"+application, application)
}
