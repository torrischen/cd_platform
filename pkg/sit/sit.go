package sit

import (
	"cd_platform/common"
	"cd_platform/pkg"
	"context"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type SitService interface {
	CreateSitDeployment(ctx context.Context, project string, deployment *appsv1.Deployment) error
	CreateSitService(ctx context.Context, project string, service *corev1.Service) error
	CreateSitStatefulset(ctx context.Context, project string, set *appsv1.StatefulSet) error
	CreateSitNamespace(ctx context.Context, project string, application string) error
	DeleteSitDeployment(ctx context.Context, project string, application string) error
	DeleteSitService(ctx context.Context, project string, application string) error
	DeleteSitStatefulset(ctx context.Context, project string, application string) error
	DeleteSitNamespace(ctx context.Context, project string, application string) error
	InsertSitIngressRule(ctx context.Context, rule *common.IngressRule) error
	DeleteSitIngressRule(ctx context.Context, project string, application string) error
	DeleteSpecifiedSitIngressRule(ctx context.Context, path string) error
	UpdateSitDeployment(ctx context.Context, project string, application string, image string) error
}

type Service struct {
	Exec pkg.ExecService
}

func NewService() *Service {
	return &Service{
		Exec: pkg.ExService,
	}
}
