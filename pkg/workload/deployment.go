package workload

import (
	"cd_platform/util"
	"context"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) CreateDeployment(ctx context.Context, project string, deployment *appsv1.Deployment) error {
	if _, err := s.Mid.K8sclient.ClientSet.AppsV1().Deployments(util.ProjectToNS(project)).Create(ctx, deployment, metav1.CreateOptions{}); err != nil {
		util.Logger.Errorf("exec.CreateDeployment err: %s", err)
		return err
	}

	return nil
}

func (s *Service) DeleteDeployment(ctx context.Context, project string, application string) error {
	if err := s.Mid.K8sclient.ClientSet.AppsV1().Deployments(util.ProjectToNS(project)).Delete(ctx, application, metav1.DeleteOptions{}); err != nil {
		util.Logger.Errorf("exec.DeleteDeployment err: %s", err)
		return err
	}

	return nil
}

func (s *Service) UpdateDeploymentImage(ctx context.Context, project string, application string, image string) error {
	dep, err := s.Mid.K8sclient.ClientSet.AppsV1().Deployments(util.ProjectToNS(project)).Get(ctx, application, metav1.GetOptions{})
	if err != nil {
		util.Logger.Errorf("exec.UpdateDeploymentImage err: %s", err)
		return err
	}

	dep.Spec.Template.Spec.Containers[0].Image = image
	_, err = s.Mid.K8sclient.ClientSet.AppsV1().Deployments(util.ProjectToNS(project)).Update(ctx, dep, metav1.UpdateOptions{})
	if err != nil {
		util.Logger.Errorf("exec.UpdateDeploymentImage err: %s", err)
		return err
	}

	return nil
}

func (s *Service) PatchDeploymentReplica(ctx context.Context, project string, application string, replica int32) error {
	newpatchmap := map[string]interface{}{
		"spec": map[string]interface{}{
			"replicas": replica,
		},
	}
	replicapatch, err := json.Marshal(newpatchmap)
	if err != nil {
		util.Logger.Errorf("exec.PatchDeploymentReplica err: %s", err)
		return err
	}

	_, err = s.Mid.K8sclient.ClientSet.AppsV1().Deployments(util.ProjectToNS(project)).Patch(ctx, application, types.MergePatchType, replicapatch, metav1.PatchOptions{})
	if err != nil {
		util.Logger.Errorf("exec.PatchDeploymentReplica err: %s", err)
		return err
	}

	return nil
}
