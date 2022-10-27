package watch

import (
	"cd_platform/api"
	"cd_platform/mid"
	"cd_platform/util"
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
)

type WatchService interface {
	GetDeploymentByName(ctx context.Context, ns string, name string) (*appsv1.Deployment, error)
	GetDeploymentByLabel(ctx context.Context, cond *api.SelectorCondList) ([]*appsv1.Deployment, error)
	GetPodByName(ctx context.Context, ns string, name string) (*corev1.Pod, error)
	GetPodByLabel(ctx context.Context, cond *api.SelectorCondList) ([]*corev1.Pod, error)
}

type Service struct {
	Mid *mid.Middle
}

func NewService(mid *mid.Middle) *Service {
	return &Service{
		Mid: mid,
	}
}

func (s *Service) GetDeploymentByName(ctx context.Context, ns string, name string) (*appsv1.Deployment, error) {
	ret, err := s.Mid.K8sclient.DeploymentLister.Deployments(ns).Get(name)
	if err != nil {
		util.Logger.Errorf("watch.GetDeploymentByName err: %s", err)
		return nil, err
	}
	return ret, nil
}

func (s *Service) GetDeploymentByLabel(ctx context.Context, cond *api.SelectorCondList) ([]*appsv1.Deployment, error) {
	selector := labels.NewSelector()
	for i := 0; i < len(cond.Cond); i++ {
		r, err := labels.NewRequirement(cond.Cond[i].Key, selection.Operator(cond.Cond[i].Operation), cond.Cond[i].Value)
		if err != nil {
			util.Logger.Errorf("watch.GetDeploymentByLabel err: %s", err)
			continue
		}
		selector.Add(*r)
	}

	ret, err := s.Mid.K8sclient.DeploymentLister.List(selector)
	if err != nil {
		util.Logger.Errorf("watch.GetDeploymentByLabel err: %s", err)
		return nil, err
	}

	return ret, nil
}

func (s *Service) GetPodByName(ctx context.Context, ns string, name string) (*corev1.Pod, error) {
	ret, err := s.Mid.K8sclient.PodLister.Pods(ns).Get(name)
	if err != nil {
		util.Logger.Errorf("watch.GetPodByName err: %s", err)
		return nil, err
	}
	return ret, nil
}

func (s *Service) GetPodByLabel(ctx context.Context, cond *api.SelectorCondList) ([]*corev1.Pod, error) {
	selector := labels.NewSelector()
	for i := 0; i < len(cond.Cond); i++ {
		r, err := labels.NewRequirement(cond.Cond[i].Key, selection.Operator(cond.Cond[i].Operation), cond.Cond[i].Value)
		if err != nil {
			util.Logger.Errorf("watch.GetPodByLabel err: %s", err)
			continue
		}
		selector.Add(*r)
	}

	ret, err := s.Mid.K8sclient.PodLister.List(selector)
	if err != nil {
		util.Logger.Errorf("watch.GetPodByLabel err: %s", err)
		return nil, err
	}

	return ret, nil
}
