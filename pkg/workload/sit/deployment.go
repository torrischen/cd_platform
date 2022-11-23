package sit

import (
	"cd_platform/util"
	"context"
	appsv1 "k8s.io/api/apps/v1"
)

func (s *Service) CreateSitDeployment(ctx context.Context, project string, deployment *appsv1.Deployment) error {
	return s.Exec.CreateDeployment(ctx, util.ToSit(project)+"-"+deployment.Name, deployment)
}

func (s *Service) DeleteSitDeployment(ctx context.Context, project string, application string) error {
	return s.Exec.DeleteDeployment(ctx, util.ToSit(project)+"-"+application, application)
}

func (s *Service) UpdateSitDeployment(ctx context.Context, project string, application string, image string) error {
	return s.Exec.UpdateDeployment(ctx, util.ToSit(project)+"-"+application, application, image)
}
