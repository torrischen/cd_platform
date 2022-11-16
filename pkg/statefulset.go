package pkg

import (
	"cd_platform/util"
	appsv1 "k8s.io/api/apps/v1"

	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) CreateStatefulset(ctx context.Context, project string, sts *appsv1.StatefulSet) error {
	if _, err := s.Mid.K8sclient.ClientSet.AppsV1().StatefulSets(util.ProjectToNS(project)).Create(ctx, sts, metav1.CreateOptions{}); err != nil {
		util.Logger.Errorf("exec.CreateStatefulSet err: %s", err)
		return err
	}

	return nil
}

func (s *Service) DeleteStatefulset(ctx context.Context, project string, application string) error {
	if err := s.Mid.K8sclient.ClientSet.AppsV1().StatefulSets(util.ProjectToNS(project)).Delete(ctx, application, metav1.DeleteOptions{}); err != nil {
		util.Logger.Errorf("exec.DeleteStatefulset err: %s", err)
		return err
	}

	return nil
}
