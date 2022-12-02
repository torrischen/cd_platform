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
	DeleteSpecifiedIngressRule(ctx context.Context, project string, path string) error
	UpdateDeploymentImage(ctx context.Context, project string, application string, image string) error
	PatchDeploymentReplica(ctx context.Context, project string, application string, replica int32) error
	CreateProjectIngress(ctx context.Context, project string) error
	RestartDeployment(ctx context.Context, project string, application string) error
	SetDeploymentEnv(ctx context.Context, project string, application string, envs []corev1.EnvVar) error
	CreateApplicationConfigmap(ctx context.Context, project string, application string) error
	AddConfigToConfigmap(ctx context.Context, project string, application string, configs []common.Config) error
	DeleteConfigmap(ctx context.Context, project string, application string) error
	DeleteSpecifiedConfig(ctx context.Context, project string, application string, configName string) error
	UpdateSpecifiedConfig(ctx context.Context, project string, application string, configName string, newVal string) error
	CrdCreateProject(ctx context.Context, project string) error
	CrdAddApplicationToProject(ctx context.Context, project string, application string) error
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
