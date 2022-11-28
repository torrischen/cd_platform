package workload

import (
	"cd_platform/common"
	"cd_platform/util"
	"context"
	"errors"
	"strings"

	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) CreateProjectIngress(ctx context.Context, project string) error {
	pathType := networkv1.PathType("Prefix")
	dftpath := networkv1.HTTPIngressPath{
		Path:     "/default",
		PathType: &pathType,
		Backend: networkv1.IngressBackend{
			Service: &networkv1.IngressServiceBackend{
				Name: "default",
				Port: networkv1.ServiceBackendPort{
					Number: 1,
				},
			},
		},
	}

	ingrule := networkv1.IngressRule{
		IngressRuleValue: networkv1.IngressRuleValue{
			HTTP: &networkv1.HTTPIngressRuleValue{
				Paths: []networkv1.HTTPIngressPath{dftpath},
			},
		},
	}

	ing := &networkv1.Ingress{
		Spec: networkv1.IngressSpec{
			Rules: []networkv1.IngressRule{
				ingrule,
			},
		},
	}
	ing.Name = util.ProjectToNS(project)

	_, err := s.Mid.K8sclient.ClientSet.NetworkingV1().Ingresses(util.ProjectToNS(project)).Create(ctx, ing, metav1.CreateOptions{})
	if err != nil {
		util.Logger.Errorf("ExecService.CreateProjectIngress err: %s", err)
		return err
	}

	return nil
}

func (s *Service) InsertIngressRule(ctx context.Context, rule *common.IngressRule) error {
	//ing, err := s.Mid.K8sclient.ClientSet.NetworkingV1().Ingresses(util.ProjectToNS(rule.Project)).Get(ctx, util.ProjectToNS(rule.Project), metav1.GetOptions{})
	if rule.Port > 65535 || rule.Port < 1 {
		return errors.New("port out of range")
	}
	ing, err := s.Mid.K8sclient.IngressLister.Ingresses(util.ProjectToNS(rule.Project)).Get(util.ProjectToNS(rule.Project))
	if err != nil {
		util.Logger.Errorf("ExecService.InsertIngressRule err: %s", err)
		return err
	}

	pathType := networkv1.PathType("Prefix")
	newIngRule := networkv1.HTTPIngressPath{
		Path:     rule.Path,
		PathType: &pathType,
		Backend: networkv1.IngressBackend{
			Service: &networkv1.IngressServiceBackend{
				Name: rule.Application,
				Port: networkv1.ServiceBackendPort{
					Number: rule.Port,
				},
			},
		},
	}

	ing.Spec.Rules[0].HTTP.Paths = append(ing.Spec.Rules[0].HTTP.Paths, newIngRule)
	_, err = s.Mid.K8sclient.ClientSet.NetworkingV1().Ingresses(util.ProjectToNS(rule.Project)).Update(ctx, ing, metav1.UpdateOptions{})
	if err != nil {
		util.Logger.Errorf("ExecService.InsertIngressRule err: %s", err)
		return err
	}

	return nil
}

func (s *Service) DeleteIngressRule(ctx context.Context, project string, application string) error {
	ing, err := s.Mid.K8sclient.IngressLister.Ingresses(util.ProjectToNS(project)).Get(util.ProjectToNS(project))
	if err != nil {
		util.Logger.Errorf("ExecService.DeleteIngressRule err: %s", err)
		return err
	}

	newIngRule := ing.Spec.Rules[0].HTTP.Paths
	for i := 0; i < len(newIngRule); i++ {
		if strings.Contains(newIngRule[i].Path, "/"+project+"/"+application) {
			newIngRule = append(newIngRule[:i], newIngRule[i+1:]...)
			i--
		}
	}
	ing.Spec.Rules[0].HTTP.Paths = newIngRule

	_, err = s.Mid.K8sclient.ClientSet.NetworkingV1().Ingresses(util.ProjectToNS(project)).Update(ctx, ing, metav1.UpdateOptions{})
	if err != nil {
		util.Logger.Errorf("ExecService.DeleteIngressRule err: %s", err)
		return err
	}

	return nil
}

func (s *Service) DeleteSpecifiedIngressRule(ctx context.Context, project string, path string) error {
	ing, err := s.Mid.K8sclient.IngressLister.Ingresses(util.ProjectToNS(project)).Get(util.ProjectToNS(project))
	if err != nil {
		util.Logger.Errorf("ExecService.DeleteSpecifiedIngressRule err: %s", err)
		return err
	}

	newIngRule := ing.Spec.Rules[0].HTTP.Paths
	for i := 0; i < len(newIngRule); i++ {
		if newIngRule[i].Path == path {
			newIngRule = append(newIngRule[:i], newIngRule[i+1:]...)
			i--
		}
	}
	ing.Spec.Rules[0].HTTP.Paths = newIngRule

	_, err = s.Mid.K8sclient.ClientSet.NetworkingV1().Ingresses(util.ProjectToNS(project)).Update(ctx, ing, metav1.UpdateOptions{})
	if err != nil {
		util.Logger.Errorf("ExecService.DeleteSpecifiedIngressRule err: %s", err)
		return err
	}

	return nil
}
