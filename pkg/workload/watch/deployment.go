package watch

import (
	"cd_platform/common"
	"cd_platform/util"
	"encoding/json"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
	"sort"

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

func (s *Service) GetDeploymentListByProject(ctx context.Context, project string) ([]string, error) {
	deplist, err := s.Mid.K8sclient.DeploymentLister.Deployments(util.ProjectToNS(project)).List(labels.NewSelector())
	if err != nil {
		util.Logger.Errorf("watch.GetDeploymentListByProject err: %s", err)
		return nil, err
	}

	ret := make([]string, 0)
	for i := 0; i < len(deplist); i++ {
		ret = append(ret, deplist[i].Name)
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i] < ret[j]
	})

	return ret, nil
}

func (s *Service) GetDeploymentYaml(ctx context.Context, project string, application string) (string, error) {
	dep, err := s.Mid.K8sclient.DeploymentLister.Deployments(util.ProjectToNS(project)).Get(application)
	if err != nil {
		util.Logger.Errorf("watch.GetDeploymentYaml err: %s", err)
		return "", err
	}

	js, err := json.Marshal(dep)
	if err != nil {
		util.Logger.Errorf("watch.GetDeploymentYaml err: %s", err)
		return "", err
	}

	y, err := yaml.JSONToYAML(js)
	if err != nil {
		util.Logger.Errorf("watch.GetDeploymentYaml err: %s", err)
		return "", err
	}

	return util.ByteToString(y), nil
}

func (s *Service) GetDeploymentEnvs(ctx context.Context, project string, application string) ([]corev1.EnvVar, error) {
	dep, err := s.Mid.K8sclient.DeploymentLister.Deployments(util.ProjectToNS(project)).Get(application)
	if err != nil {
		util.Logger.Errorf("watch.GetDeploymentEnvs err: %s", err)
		return nil, err
	}

	return dep.Spec.Template.Spec.Containers[0].Env, nil
}
