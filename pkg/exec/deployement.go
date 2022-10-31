package exec

import (
	"cd_platform/common"
	"cd_platform/util"

	"context"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) CreateDeployment(ctx context.Context, args *common.NewDeploymentArgs) (*appsv1.Deployment, error) {
	newDep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "",
			Namespace:   "",
			Labels:      nil,
			Annotations: nil,
		},
	}

	d, err := s.Mid.K8sclient.ClientSet.AppsV1().Deployments(util.OwnerToNS(args.OwnerName)).Create(context.TODO(), newDep, metav1.CreateOptions{})
	if err != nil {
		util.Logger.Errorf("exec.CreateDeployment err: %s", err)
		return nil, err
	}

	return d, nil
}
