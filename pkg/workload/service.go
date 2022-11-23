package workload

import (
	"cd_platform/util"
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) CreateService(ctx context.Context, project string, service *corev1.Service) error {
	if _, err := s.Mid.K8sclient.ClientSet.CoreV1().Services(util.ProjectToNS(project)).Create(ctx, service, metav1.CreateOptions{}); err != nil {
		util.Logger.Errorf("exec.CreateService err: %s", err)
		return err
	}

	return nil
}

func (s *Service) DeleteService(ctx context.Context, project string, application string) error {
	if err := s.Mid.K8sclient.ClientSet.CoreV1().Services(util.ProjectToNS(project)).Delete(ctx, application, metav1.DeleteOptions{}); err != nil {
		util.Logger.Errorf("exec.DeleteService err: %s", err)
		return err
	}

	return nil
}
