package exec

import (
	"cd_platform/util"

	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
)

func (s *Service) CreateStatefulset(ctx context.Context, project string, raw []byte) error {
	obj, _, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(raw, nil, nil)
	if err != nil {
		util.Logger.Errorf("exec.CreateStatefulSet err: %s", err)
		return err
	}
	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		util.Logger.Errorf("exec.CreateStatefulSet err: %s", err)
		return err
	}

	if _, err := s.Mid.K8sclient.DynamicClient.Resource(schema.GroupVersionResource{
		Version:  "v1",
		Resource: "services",
	}).Namespace(util.ProjectToNS(project)).Create(context.TODO(), &unstructured.Unstructured{Object: m}, metav1.CreateOptions{}); err != nil {
		util.Logger.Errorf("exec.CreateStatefulSet err: %s", err)
		return err
	}

	return nil
}