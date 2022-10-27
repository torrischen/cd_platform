package watch

import (
	"cd_platform/api"
	"cd_platform/util"

	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
)

func (s *Service) GetServiceByName(ctx context.Context, ns string, name string) (*corev1.Service, error) {
	ret, err := s.Mid.K8sclient.ServiceLister.Services(ns).Get(name)
	if err != nil {
		util.Logger.Errorf("watch.GetServiceByName err: %s", err)
		return nil, err
	}
	return ret, nil
}

func (s *Service) GetServiceByLabel(ctx context.Context, cond *api.SelectorCondList) ([]*corev1.Service, error) {
	selector := labels.NewSelector()
	for i := 0; i < len(cond.Cond); i++ {
		r, err := labels.NewRequirement(cond.Cond[i].Key, selection.Operator(cond.Cond[i].Operation), cond.Cond[i].Value)
		if err != nil {
			util.Logger.Errorf("watch.GetServiceByLabel err: %s", err)
			continue
		}
		selector.Add(*r)
	}

	ret, err := s.Mid.K8sclient.ServiceLister.List(selector)
	if err != nil {
		util.Logger.Errorf("watch.GetServiceByLabel err: %s", err)
		return nil, err
	}

	return ret, nil
}
