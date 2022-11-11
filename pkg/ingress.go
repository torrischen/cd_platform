package pkg

import (
	"cd_platform/common"
	"cd_platform/util"
	"context"
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
		Path:     "/api/" + rule.Application,
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

	_, err = s.Mid.K8sclient.ClientSet.NetworkingV1().Ingresses("default").Update(context.TODO(), ing, metav1.UpdateOptions{})
	if err != nil {
		util.Logger.Errorf("ExecService.InsertIngressRule err: %s", err)
		return err
	}

	return nil
}