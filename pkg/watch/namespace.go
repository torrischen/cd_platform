package watch

import (
	"cd_platform/util"
	"context"
	corev1 "k8s.io/api/core/v1"
)

func (s *Service) GetNamespaceByName(ctx context.Context, ns string) (*corev1.Namespace, error) {
	ret, err := s.Mid.K8sclient.NSLister.Get(ns)
	if err != nil {
		util.Logger.Errorf("watch.GetNamespaceByName err: %s", err)
		return nil, err
	}
	return ret, nil
}
