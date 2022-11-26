package workload

import (
	"cd_platform/util"
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) CreateApplicationConfigmap(ctx context.Context, project string, application string) error {
	dep, err := s.Mid.K8sclient.DeploymentLister.Deployments(util.ProjectToNS(project)).Get(application)
	if err != nil {
		util.Logger.Errorf("workload.CreateApplicationConfigmap err: %s", err)
		return err
	}

	newvlm := corev1.VolumeMount{
		Name:      application,
		MountPath: "./config",
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

func (s *Service) DeleteConfigmap(ctx context.Context, project string, application string) error {
	return s.Mid.K8sclient.ClientSet.CoreV1().ConfigMaps(util.ProjectToNS(project)).Delete(ctx, application, metav1.DeleteOptions{})
}

func (s *Service) AddConfigToConfigmap(ctx context.Context, project string, application string, configName string, data string) error {
	cfm, err := s.Mid.K8sclient.CMLister.ConfigMaps(util.ProjectToNS(project)).Get(application)
	if err != nil {
		util.Logger.Errorf("workload.AddConfigToConfigmap err: %s", err)
		return err
	}

	if cfm.Data == nil {
		cfm.Data = make(map[string]string)
	}
	cfm.Data[configName] = data

	_, err = s.Mid.K8sclient.ClientSet.CoreV1().ConfigMaps(util.ProjectToNS(project)).Update(ctx, cfm, metav1.UpdateOptions{})
	if err != nil {
		util.Logger.Errorf("workload.AddConfigToConfigmap err: %s", err)
		return err
	}

	return nil
}

func (s *Service) DeleteSpecifiedConfig(ctx context.Context, project string, application string, configName string) error {
	cfm, err := s.Mid.K8sclient.CMLister.ConfigMaps(util.ProjectToNS(project)).Get(application)
	if err != nil {
		util.Logger.Errorf("workload.DeleteSpecifiedConfig err: %s", err)
		return err
	}

	delete(cfm.Data, configName)

	_, err = s.Mid.K8sclient.ClientSet.CoreV1().ConfigMaps(util.ProjectToNS(project)).Update(ctx, cfm, metav1.UpdateOptions{})
	if err != nil {
		util.Logger.Errorf("workload.DeleteSpecifiedConfig err: %s", err)
		return err
	}

	return nil
}

func (s *Service) UpdateSpecifiedConfig(ctx context.Context, project string, application string, configName string, newVal string) error {
	cfm, err := s.Mid.K8sclient.CMLister.ConfigMaps(util.ProjectToNS(project)).Get(application)
	if err != nil {
		util.Logger.Errorf("workload.UpdateSpecifiedConfig getcfm err: %s", err)
		return err
	}

	cfm.Data[configName] = newVal

	_, err = s.Mid.K8sclient.ClientSet.CoreV1().ConfigMaps(util.ProjectToNS(project)).Update(ctx, cfm, metav1.UpdateOptions{})
	if err != nil {
		util.Logger.Errorf("workload.UpdateSpecifiedConfig update err: %s", err)
		return err
	}

	return nil
}
