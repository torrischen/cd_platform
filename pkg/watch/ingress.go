package watch

import (
	"cd_platform/api"
	"cd_platform/util"

	"context"

	networkv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
)

func (s *Service) GetIngressByName(ctx context.Context, ns string, name string) (*networkv1.Ingress, error) {
	ret, err := s.Mid.K8sclient.IngressLister.Ingresses(ns).Get(name)
	if err != nil {
		util.Logger.Errorf("watch.GetIngressByName err: %s", err)
		return nil, err
	}
	return ret, nil
}

func (s *Service) GetIngressByLabel(ctx context.Context, cond *api.SelectorCondList) ([]*networkv1.Ingress, error) {
	selector := labels.NewSelector()
	for i := 0; i < len(cond.Cond); i++ {
		r, err := labels.NewRequirement(cond.Cond[i].Key, selection.Operator(cond.Cond[i].Operation), cond.Cond[i].Value)
		if err != nil {
			util.Logger.Errorf("watch.GetIngressByLabel err: %s", err)
			continue
		}
		selector.Add(*r)
	}

	ret, err := s.Mid.K8sclient.IngressLister.List(selector)
	if err != nil {
		util.Logger.Errorf("watch.GetIngressByLabel err: %s", err)
		return nil, err
	}

	return ret, nil
}
