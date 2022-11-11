package pkg

import (
	"cd_platform/util"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (s *Service) CreateService(ctx context.Context, project string, raw []byte) error {
	uns, err := util.RawJsonToUnstructured(raw)
	if err != nil {
		util.Logger.Errorf("exec.CreateService err: %s", err)
		return err
	}

	if _, err := s.Mid.K8sclient.DynamicClient.Resource(schema.GroupVersionResource{
		Version:  "v1",
		Resource: "services",
	}).Namespace(util.ProjectToNS(project)).Create(ctx, uns, metav1.CreateOptions{}); err != nil {
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
