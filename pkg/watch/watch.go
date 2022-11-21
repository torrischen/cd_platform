package watch

import (
	"cd_platform/common"
	"cd_platform/mid"

	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
)

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
	GetPodByProject(ctx context.Context, project string) ([]*common.PodByApplication, error)
	GetSitPodByApplication(ctx context.Context, application string) (*common.PodByApplication, error)
	GetIngressByApplication(ctx context.Context, project string, application string) ([]*common.IngressRule, error)
	GetSitIngressByApplication(ctx context.Context, application string) ([]*common.SitIngressRule, error)
	GetPodLog(ctx context.Context, project string, podname string) ([]byte, error)
	GetSitPodLog(ctx context.Context, application string, podname string) ([]byte, error)
}

type Service struct {
	Mid *mid.Middle
}

func NewService(mid *mid.Middle) *Service {
	return &Service{
		Mid: mid,
	}
}
