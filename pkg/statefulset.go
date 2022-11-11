package pkg

import (
	"cd_platform/util"

	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (s *Service) CreateStatefulset(ctx context.Context, project string, raw []byte) error {
	uns, err := util.RawJsonToUnstructured(raw)
	if err != nil {
		util.Logger.Errorf("exec.CreateStatefulSet err: %s", err)
		return err
	}

	if _, err := s.Mid.K8sclient.DynamicClient.Resource(schema.GroupVersionResource{
		Version:  "v1",
		Resource: "services",
	}).Namespace(util.ProjectToNS(project)).Create(ctx, uns, metav1.CreateOptions{}); err != nil {
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
