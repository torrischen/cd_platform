package watch

import (
	"cd_platform/common"
	"cd_platform/util"

	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func (s *Service) GetServiceByName(ctx context.Context, ns string, name string) (*corev1.Service, error) {
	ret, err := s.Mid.K8sclient.ServiceLister.Services(ns).Get(name)
	if err != nil {
		util.Logger.Errorf("watch.GetServiceByName err: %s", err)
		return nil, err
	}
	return ret, nil
}

func (s *Service) GetServiceByLabel(ctx context.Context, cond *common.SelectorCondList) ([]*corev1.Service, error) {
	m := make(map[string]string)
	for i := 0; i < len(cond.Cond); i++ {
		m[cond.Cond[i].Key] = cond.Cond[i].Value
	}
	slt := labels.SelectorFromSet(m)

	ret, err := s.Mid.K8sclient.ServiceLister.List(slt)
	if err != nil {
		util.Logger.Errorf("watch.GetServiceByLabel err: %s", err)
		return nil, err
	}

	return ret, nil
}
