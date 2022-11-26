package workload

import (
	"cd_platform/util"
	"context"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	"time"

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

func (s *Service) RestartDeployment(ctx context.Context, project string, application string) error {
	patchdata := map[string]interface{}{
		"spec": map[string]interface{}{
			"template": map[string]interface{}{
				"metadata": map[string]interface{}{
					"annotations": map[string]interface{}{
						"restarted_at": time.Now().Format("2006-01-02 15:04:05"),
					},
				},
			},
		},
	}
	patchjson, err := json.Marshal(patchdata)
	if err != nil {
		util.Logger.Errorf("workload.RestartDeployment err: %s", err)
		return err
	}

	_, err = s.Mid.K8sclient.ClientSet.AppsV1().Deployments(util.ProjectToNS(project)).Patch(ctx, application, types.MergePatchType, patchjson, metav1.PatchOptions{})
	if err != nil {
		util.Logger.Errorf("workload.RestartDeployment err: %s", err)
		return err
	}

	return nil
}

func (s *Service) SetDeploymentEnv(ctx context.Context, project string, application string, envs []corev1.EnvVar) error {
	dep, err := s.Mid.K8sclient.DeploymentLister.Deployments(util.ProjectToNS(project)).Get(application)
	if err != nil {
		util.Logger.Errorf("workload.SetDeploymentEnv err: %s", err)
		return err
	}

	dep.Spec.Template.Spec.Containers[0].Env = envs

	_, err = s.Mid.K8sclient.ClientSet.AppsV1().Deployments(util.ProjectToNS(project)).Update(ctx, dep, metav1.UpdateOptions{})
	if err != nil {
		util.Logger.Errorf("workload.SetDeploymentEnv err: %s", err)
		return err
	}

	return nil
}

func (s *Service) CreateApplicationConfigmap(ctx context.Context, project string, application string) error {
	dep, err := s.Mid.K8sclient.DeploymentLister.Deployments(util.ProjectToNS(project)).Get(application)
	if err != nil {
		util.Logger.Errorf("workload.CreateApplicationConfigmap err: %s", err)
		return err
	}

	newvlm := corev1.VolumeMount{
		Name:      application,
		MountPath: "./conf",
	}
	newvl := corev1.Volume{
		Name: application,
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: application,
				},
			},
		},
	}
	newcfg := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: application,
		},
	}
	dep.Spec.Template.Spec.Containers[0].VolumeMounts = append(dep.Spec.Template.Spec.Containers[0].VolumeMounts, newvlm)
	dep.Spec.Template.Spec.Volumes = append(dep.Spec.Template.Spec.Volumes, newvl)

	_, err = s.Mid.K8sclient.ClientSet.CoreV1().ConfigMaps(util.ProjectToNS(project)).Create(ctx, newcfg, metav1.CreateOptions{})
	if err != nil {
		util.Logger.Errorf("workload.CreateApplicationConfigmap err: %s", err)
		return err
	}
	_, err = s.Mid.K8sclient.ClientSet.AppsV1().Deployments(util.ProjectToNS(project)).Update(ctx, dep, metav1.UpdateOptions{})
	if err != nil {
		util.Logger.Errorf("workload.CreateApplicationConfigmap err: %s", err)
		return err
	}

	return nil
}
