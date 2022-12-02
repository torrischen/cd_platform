package workload

import (
	crdv1 "cd_platform/common/v1"
	"cd_platform/util"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (s *Service) CrdCreateProject(ctx context.Context, project string) error {
	proj := crdv1.Project{}
	proj.Name = project
	proj.Spec.Application = []string{}

	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(proj)
	if err != nil {
		util.Logger.Errorf("workload.CrdCreateProject err : %s", err)
		return err
	}

	un := &unstructured.Unstructured{Object: m}

	_, err = s.Mid.K8sclient.DynamicClient.Resource(schema.GroupVersionResource{
		Group:    "cytcrd.nainaiguan.com",
		Version:  "v1",
		Resource: "projects",
	}).Namespace(util.ProjectToNS(project)).Create(ctx, un, metav1.CreateOptions{})
	if err != nil {
		util.Logger.Errorf("workload.CrdCreateProject err : %s", err)
		return err
	}

	return nil
}

func (s *Service) CrdAddApplicationToProject(ctx context.Context, project string, application string) error {
	obj, err := s.Mid.K8sclient.ProjectLister.ByNamespace(util.ProjectToNS(project)).Get(project)
	if err != nil {
		util.Logger.Errorf("workload.CrdAddApplicationToProject err : %s", err)
		return err
	}

	un, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		util.Logger.Errorf("workload.CrdAddApplicationToProject err : %s", err)
		return err
	}

	proj := &crdv1.Project{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(un, proj)
	if err != nil {
		util.Logger.Errorf("workload.CrdAddApplicationToProject err : %s", err)
		return err
	}

	proj.Spec.Application = append(proj.Spec.Application, application)

	newun, err := runtime.DefaultUnstructuredConverter.ToUnstructured(proj)
	_, err = s.Mid.K8sclient.DynamicClient.Resource(schema.GroupVersionResource{
		Group:    "cytcrd.nainaiguan.com",
		Version:  "v1",
		Resource: "projects",
	}).Namespace(util.ProjectToNS(project)).Create(ctx, &unstructured.Unstructured{Object: newun}, metav1.CreateOptions{})
	if err != nil {
		util.Logger.Errorf("workload.CrdAddApplicationToProject err : %s", err)
		return err
	}

	return nil
}
