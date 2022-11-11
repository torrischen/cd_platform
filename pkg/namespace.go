package pkg

import (
	"cd_platform/util"
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) CreateNamespace(ctx context.Context, project string) error {
	newns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: util.ProjectToNS(project),
		},
	}

	if _, err := s.Mid.K8sclient.ClientSet.CoreV1().Namespaces().Create(ctx, newns, metav1.CreateOptions{}); err != nil {
		util.Logger.Errorf("exec.CreateNamespace err: %s", err)
		return err
	}

	return nil
}

func (s *Service) DeleteNamespace(ctx context.Context, project string) error {
	if err := s.Mid.K8sclient.ClientSet.CoreV1().Namespaces().Delete(ctx, util.ProjectToNS(project), metav1.DeleteOptions{}); err != nil {
		util.Logger.Errorf("exec.DeleteNamespace err: %s", err)
		return err
	}

	return nil
}
