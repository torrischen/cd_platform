package pkg

import (
	"cd_platform/common"
	"cd_platform/util"
	"context"
	"strings"

	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) InsertIngressRule(ctx context.Context, rule *common.IngressRule) error {
	ing, err := s.Mid.K8sclient.IngressLister.Ingresses("default").Get("cd-ingress")
	if err != nil {
		util.Logger.Errorf("ExecService.InsertIngressRule err: %s", err)
		return err
	}

	pathType := networkv1.PathType("Prefix")
	newIngRule := networkv1.HTTPIngressPath{
		Path:     "/api/" + rule.Project + "/" + rule.Application + "/" + rule.Path,
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
	_, err = s.Mid.K8sclient.ClientSet.NetworkingV1().Ingresses("default").Update(ctx, ing, metav1.UpdateOptions{})
	if err != nil {
		util.Logger.Errorf("ExecService.InsertIngressRule err: %s", err)
		return err
	}

	return nil
}

func (s *Service) DeleteIngressRule(ctx context.Context, project string, application string) error {
	ing, err := s.Mid.K8sclient.IngressLister.Ingresses("default").Get("cd-ingress")
	if err != nil {
		util.Logger.Errorf("ExecService.DeleteIngressRule err: %s", err)
		return err
	}

	newIngRule := ing.Spec.Rules[0].HTTP.Paths
	for i := 0; i < len(ing.Spec.Rules[0].HTTP.Paths); i++ {
		if strings.Contains(ing.Spec.Rules[0].HTTP.Paths[i].Path, "/api/"+project+"/"+application) {
			newIngRule = append(newIngRule[:i], newIngRule[i+1:]...)
		}
	}
	ing.Spec.Rules[0].HTTP.Paths = newIngRule

	_, err = s.Mid.K8sclient.ClientSet.NetworkingV1().Ingresses("default").Update(ctx, ing, metav1.UpdateOptions{})
	if err != nil {
		util.Logger.Errorf("ExecService.DeleteIngressRule err: %s", err)
		return err
	}

	return nil
}
