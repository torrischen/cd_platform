package workload

import (
	"cd_platform/common"
	"cd_platform/mid"
	"context"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

var exService *Service

type ExecService interface {
	CreateDeployment(ctx context.Context, project string, deployment *appsv1.Deployment) error
	CreateService(ctx context.Context, project string, service *corev1.Service) error
	CreateStatefulset(ctx context.Context, project string, sts *appsv1.StatefulSet) error
	CreateNamespace(ctx context.Context, project string) error
	DeleteDeployment(ctx context.Context, project string, application string) error
	DeleteService(ctx context.Context, project string, application string) error
	DeleteStatefulset(ctx context.Context, project string, application string) error
	DeleteNamespace(ctx context.Context, project string) error
	InsertIngressRule(ctx context.Context, rule *common.IngressRule) error
	DeleteIngressRule(ctx context.Context, project string, application string) error
	DeleteSpecifiedIngressRule(ctx context.Context, path string) error
	UpdateDeployment(ctx context.Context, project string, application string, image string) error
	PatchDeploymentReplica(ctx context.Context, project string, application string, replica int32) error
}

type Service struct {
	Mid *mid.Middle
}

func Init(mid *mid.Middle) {
	exService = &Service{
		Mid: mid,
	}
}

func NewService() *Service {
	return exService
}
