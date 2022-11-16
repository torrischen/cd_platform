package sit

import (
	"cd_platform/common"
	"cd_platform/mid"
	"cd_platform/pkg"
	"context"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type SitService interface {
	CreateSitDeployment(ctx context.Context, application string, deployment *appsv1.Deployment) error
	CreateSitService(ctx context.Context, application string, service *corev1.Service) error
	CreateSitStatefulset(ctx context.Context, application string, set *appsv1.StatefulSet) error
	CreateSitNamespace(ctx context.Context, application string) error
	DeleteSitDeployment(ctx context.Context, application string) error
	DeleteSitService(ctx context.Context, application string) error
	DeleteSitStatefulset(ctx context.Context, application string) error
	DeleteSitNamespace(ctx context.Context, application string) error
	InsertSitIngressRule(ctx context.Context, application string, rule *common.IngressRule) error
	DeleteSitIngressRule(ctx context.Context, application string) error
	UpdateSitDeployment(ctx context.Context, application string, image string) error
}

type Service struct {
	Exec pkg.ExecService
}

func NewService(mid *mid.Middle) *Service {
	return &Service{
		Exec: pkg.ExService,
	}
}
