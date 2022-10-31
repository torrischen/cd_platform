package watch

import (
	"cd_platform/common"
	"cd_platform/util"

	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func (s *Service) GetPodByName(ctx context.Context, ns string, name string) (*corev1.Pod, error) {
	ret, err := s.Mid.K8sclient.PodLister.Pods(ns).Get(name)
	if err != nil {
		util.Logger.Errorf("watch.GetPodByName err: %s", err)
		return nil, err
	}
	return ret, nil
}

func (s *Service) GetPodByLabel(ctx context.Context, cond *common.SelectorCondList) ([]*corev1.Pod, error) {
	m := make(map[string]string)
	for i := 0; i < len(cond.Cond); i++ {
		m[cond.Cond[i].Key] = cond.Cond[i].Value
	}
	slt := labels.SelectorFromSet(m)

	ret, err := s.Mid.K8sclient.PodLister.List(slt)
	if err != nil {
		util.Logger.Errorf("watch.GetPodByLabel err: %s", err)
		return nil, err
	}

	return ret, nil
}
