package watch

import (
	"cd_platform/common"
	"cd_platform/mid"
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
)

var watchService *Service

type WatchService interface {
	GetDeploymentByName(ctx context.Context, ns string, name string) (*appsv1.Deployment, error)
	GetDeploymentByLabel(ctx context.Context, cond *common.SelectorCondList) ([]*appsv1.Deployment, error)
	GetPodByName(ctx context.Context, ns string, name string) (*corev1.Pod, error)
	GetPodByLabel(ctx context.Context, cond *common.SelectorCondList) ([]*corev1.Pod, error)
	GetStatefulSetByName(ctx context.Context, ns string, name string) (*appsv1.StatefulSet, error)
	GetStatefulSetByLabel(ctx context.Context, cond *common.SelectorCondList) ([]*appsv1.StatefulSet, error)
	GetServiceByName(ctx context.Context, ns string, name string) (*corev1.Service, error)
	GetServiceByLabel(ctx context.Context, cond *common.SelectorCondList) ([]*corev1.Service, error)
	GetIngressByName(ctx context.Context, ns string, name string) (*networkv1.Ingress, error)
	GetIngressByLabel(ctx context.Context, cond *common.SelectorCondList) ([]*networkv1.Ingress, error)
	GetNamespaceByName(ctx context.Context, ns string) (*corev1.Namespace, error)
	GetPodByApplication(ctx context.Context, project string, application string) ([]*common.PodDetail, error)
	GetPodByProject(ctx context.Context, project string) ([]*common.PodByApplication, error)
	GetIngressByApplication(ctx context.Context, project string, application string) ([]*common.IngressRule, error)
	GetPodLog(ctx context.Context, project string, podname string) ([]byte, error)
	GetDeploymentListByProject(ctx context.Context, project string) ([]string, error)
	GetDeploymentEnvs(ctx context.Context, project string, application string) ([]corev1.EnvVar, error)
	GetDeploymentYaml(ctx context.Context, project string, application string) (string, error)
	GetApplicationConfigList(ctx context.Context, project string, application string) (*common.ConfigList, error)
	GetApplicationConfigDetail(ctx context.Context, project string, application string) ([]*common.ConfigDetail, error)
}

type Service struct {
	Mid *mid.Middle
}

func Init(mid *mid.Middle) {
	watchService = &Service{
		Mid: mid,
	}
}

func NewService() *Service {
	return watchService
}
