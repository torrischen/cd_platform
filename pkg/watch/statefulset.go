package watch

import (
	"cd_platform/api"
	"cd_platform/util"

	"context"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
)

func (s *Service) GetStatefulSetByName(ctx context.Context, ns string, name string) (*appsv1.StatefulSet, error) {
	ret, err := s.Mid.K8sclient.StatefulSetLister.StatefulSets(ns).Get(name)
	if err != nil {
		util.Logger.Errorf("watch.GetStatefulSetByName err: %s", err)
		return nil, err
	}
	return ret, nil
}

func (s *Service) GetStatefulSetByLabel(ctx context.Context, cond *api.SelectorCondList) ([]*appsv1.StatefulSet, error) {
	selector := labels.NewSelector()
	for i := 0; i < len(cond.Cond); i++ {
		r, err := labels.NewRequirement(cond.Cond[i].Key, selection.Operator(cond.Cond[i].Operation), cond.Cond[i].Value)
		if err != nil {
			util.Logger.Errorf("watch.GetStatefulSetByLabel err: %s", err)
			continue
		}
		selector.Add(*r)
	}

	ret, err := s.Mid.K8sclient.StatefulSetLister.List(selector)
	if err != nil {
		util.Logger.Errorf("watch.GetStatefulSetByLabel err: %s", err)
		return nil, err
	}

	return ret, nil
}
