package watch

import (
	"cd_platform/common"
	"cd_platform/util"
	"strings"

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

func (s *Service) GetIngressByApplication(ctx context.Context, project string, application string) ([]*common.IngressRule, error) {
	ing, err := s.GetIngressByName(ctx, util.ProjectToNS(project), util.ProjectToNS(project))
	if err != nil {
		util.Logger.Errorf("watch.GetIngressByApplication err: %s", err)
		return nil, err
	}

	ret := make([]*common.IngressRule, 0)
	rules := ing.Spec.Rules[0].HTTP.Paths
	for i := 0; i < len(rules); i++ {
		if strings.Contains(rules[i].Path, "/api/"+project+"/"+application) {
			tmp := &common.IngressRule{}
			tmp.Project = project
			tmp.Application = application
			tmp.Path = rules[i].Path
			tmp.Port = rules[i].Backend.Service.Port.Number
			ret = append(ret, tmp)
		}
	}

	return ret, nil
}
