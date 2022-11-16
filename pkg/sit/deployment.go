package sit

import (
	"cd_platform/util"
	"context"
	appsv1 "k8s.io/api/apps/v1"
)

func (s *Service) CreateSitDeployment(ctx context.Context, application string, deployment *appsv1.Deployment) error {
	return s.Exec.CreateDeployment(ctx, util.ToSit(application), deployment)
}

func (s *Service) DeleteSitDeployment(ctx context.Context, application string) error {
	return s.Exec.DeleteDeployment(ctx, util.ToSit(application), application)
}

func (s *Service) UpdateSitDeployment(ctx context.Context, application string, image string) error {
	return s.Exec.UpdateDeployment(ctx, util.ToSit(application), util.ToSit(application), image)
}
