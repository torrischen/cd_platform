package exec

import (
	"cd_platform/common"
	"cd_platform/util"
	"context"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) CreateDeployment(ctx context.Context, args *common.NewDeploymentArgs) (*appsv1.Deployment, error) {
	labelM := make(map[string]string)
	for i := 0; i < len(args.Labels); i++ {
		labelM[args.Labels[i].Key] = args.Labels[i].Value
	}

	annoM := make(map[string]string)
	for i := 0; i < len(args.Annotations); i++ {
		annoM[args.Annotations[i].Key] = args.Annotations[i].Value
	}

	newDep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:        args.DeploymentName,
			Namespace:   util.OwnerToNS(args.OwnerName),
			Labels:      labelM,
			Annotations: annoM,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas:                nil,
			Selector:                nil,
			Template:                corev1.PodTemplateSpec{},
			Strategy:                appsv1.DeploymentStrategy{},
			MinReadySeconds:         0,
			RevisionHistoryLimit:    nil,
			Paused:                  false,
			ProgressDeadlineSeconds: nil,
		},
	}

	d, err := s.Mid.K8sclient.ClientSet.AppsV1().Deployments(util.OwnerToNS(args.OwnerName)).Create(context.TODO(), newDep, metav1.CreateOptions{})
	if err != nil {
		util.Logger.Errorf("exec.CreateDeployment err: %s", err)
		return nil, err
	}

	return d, nil
}
