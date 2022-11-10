package pkg

import (
	"cd_platform/util"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) CreateDeployment(ctx context.Context, project string, raw []byte) error {
	uns, err := util.RawJsonToUnstructured(raw)
	if err != nil {
		util.Logger.Errorf("exec.CreateDeployment err: %s", err)
		return err
	}

	if _, err := s.Mid.K8sclient.DynamicClient.Resource(schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}).Namespace(util.ProjectToNS(project)).Create(context.TODO(), uns, metav1.CreateOptions{}); err != nil {
		util.Logger.Errorf("exec.CreateDeployment err: %s", err)
		return err
	}

	return nil
}

func (s *Service) DeleteDeployment(ctx context.Context, project string) error {
	if err := s.Mid.K8sclient.ClientSet.AppsV1().Deployments(util.ProjectToNS(project)).Delete(context.TODO(), project, metav1.DeleteOptions{}); err != nil {
		util.Logger.Errorf("exec.DeleteDeployment err: %s", err)
		return err
	}

	return nil
}
