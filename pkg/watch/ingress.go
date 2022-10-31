package watch

import (
	"cd_platform/common"
	"cd_platform/util"

	"context"

	networkv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func (s *Service) GetIngressByName(ctx context.Context, ns string, name string) (*networkv1.Ingress, error) {
	ret, err := s.Mid.K8sclient.IngressLister.Ingresses(ns).Get(name)
	if err != nil {
		util.Logger.Errorf("watch.GetIngressByName err: %s", err)
		return nil, err
	}
	return ret, nil
}

func (s *Service) GetIngressByLabel(ctx context.Context, cond *common.SelectorCondList) ([]*networkv1.Ingress, error) {
	m := make(map[string]string)
	for i := 0; i < len(cond.Cond); i++ {
		m[cond.Cond[i].Key] = cond.Cond[i].Value
	}
	slt := labels.SelectorFromSet(m)

	ret, err := s.Mid.K8sclient.IngressLister.List(slt)
	if err != nil {
		util.Logger.Errorf("watch.GetIngressByLabel err: %s", err)
		return nil, err
	}

	return ret, nil
}
