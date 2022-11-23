package watch

import (
	"cd_platform/common"
	"cd_platform/util"

	"context"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func (s *Service) GetDeploymentByName(ctx context.Context, ns string, name string) (*appsv1.Deployment, error) {
	ret, err := s.Mid.K8sclient.DeploymentLister.Deployments(ns).Get(name)
	if err != nil {
		util.Logger.Errorf("watch.GetDeploymentByName err: %s", err)
		return nil, err
	}
	return ret, nil
}

func (s *Service) GetDeploymentByLabel(ctx context.Context, cond *common.SelectorCondList) ([]*appsv1.Deployment, error) {
	m := make(map[string]string)
	for i := 0; i < len(cond.Cond); i++ {
		m[cond.Cond[i].Key] = cond.Cond[i].Value
	}
	slt := labels.SelectorFromSet(m)

	ret, err := s.Mid.K8sclient.DeploymentLister.List(slt)
	if err != nil {
		util.Logger.Errorf("watch.GetDeploymentByLabel err: %s", err)
		return nil, err
	}

	return ret, nil
}
