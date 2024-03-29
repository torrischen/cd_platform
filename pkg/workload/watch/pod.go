package watch

import (
	"cd_platform/common"
	"cd_platform/util"
	"context"
	"io"
	"sort"
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func (s *Service) GetPodByName(ctx context.Context, ns string, name string) (*corev1.Pod, error) {
	ret, err := s.Mid.K8sclient.PodLister.Pods(ns).Get(name)
	if err != nil {
		util.Logger.Errorf("watch.GetPodByName err: %s", err)
		return nil, err
	}
	return ret, nil
}

func (s *Service) GetPodByLabel(ctx context.Context, cond *common.SelectorCondList) ([]*corev1.Pod, error) {
	m := make(map[string]string)
	for i := 0; i < len(cond.Cond); i++ {
		m[cond.Cond[i].Key] = cond.Cond[i].Value
	}
	slt := labels.SelectorFromSet(m)

	ret, err := s.Mid.K8sclient.PodLister.List(slt)
	if err != nil {
		util.Logger.Errorf("watch.GetPodByLabel err: %s", err)
		return nil, err
	}

	return ret, nil
}

func (s *Service) GetPodByApplication(ctx context.Context, project string, application string) ([]*common.PodDetail, error) {
	podlist, err := s.GetPodByLabel(ctx, &common.SelectorCondList{Cond: []common.SelectorCond{{Key: "app", Value: application}}})
	if err != nil {
		util.Logger.Errorf("watch.GetPodByApplication err: %s", err)
		return nil, err
	}

	podDetails := make([]*common.PodDetail, 0)
	for i := 0; i < len(podlist); i++ {
		if podlist[i].Namespace != util.ProjectToNS(project) {
			continue
		}
		pd := &common.PodDetail{
			Name:       podlist[i].Name,
			Namespace:  podlist[i].Namespace,
			Image:      podlist[i].Spec.Containers[0].Image,
			CreateTime: podlist[i].CreationTimestamp.In(common.TimeZone).Format(common.DateTimeLayout),
			HostIp:     podlist[i].Status.HostIP,
			PodIp:      podlist[i].Status.PodIP,
			Status:     &podlist[i].Status.ContainerStatuses[0].State,
		}
		podDetails = append(podDetails, pd)
	}

	sort.Slice(podDetails, func(i, j int) bool {
		return podDetails[i].Name < podDetails[j].Name
	})

	return podDetails, nil
}

func (s *Service) GetPodByProject(ctx context.Context, project string) ([]*common.PodByApplication, error) {
	deplist, err := s.Mid.K8sclient.DeploymentLister.Deployments(util.ProjectToNS(project)).List(labels.NewSelector())
	if err != nil {
		util.Logger.Errorf("watch.GetPodByProject err: %s", err)
		return nil, err
	}

	ret := make([]*common.PodByApplication, 0)
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(deplist))
	for i := 0; i < len(deplist); i++ {
		application := deplist[i].Name
		go func() {
			defer wg.Done()

			podlist, err := s.GetPodByLabel(ctx, &common.SelectorCondList{Cond: []common.SelectorCond{{Key: "app", Value: application}}})
			if err != nil {
				util.Logger.Errorf("watch.GetPodByProject err: %s", err)
				return
			}

			podDetails := make([]*common.PodDetail, 0)
			for i := 0; i < len(podlist); i++ {
				if podlist[i].Namespace != util.ProjectToNS(project) {
					continue
				}
				pd := &common.PodDetail{
					Name:       podlist[i].Name,
					Namespace:  podlist[i].Namespace,
					Image:      podlist[i].Spec.Containers[0].Image,
					CreateTime: podlist[i].CreationTimestamp.In(common.TimeZone).Format(common.DateTimeLayout),
					HostIp:     podlist[i].Status.HostIP,
					PodIp:      podlist[i].Status.PodIP,
					Status:     &podlist[i].Status.ContainerStatuses[0].State,
				}
				podDetails = append(podDetails, pd)
			}

			tmp := &common.PodByApplication{
				Application: application,
				Pods:        podDetails,
			}

			m.Lock()
			ret = append(ret, tmp)
			m.Unlock()
		}()
	}
	wg.Wait()

	return ret, nil
}

func (s *Service) GetPodLog(ctx context.Context, project string, podname string) (io.ReadCloser, error) {
	line := int64(3000)
	log := s.Mid.K8sclient.ClientSet.CoreV1().Pods(util.ProjectToNS(project)).GetLogs(
		podname,
		&corev1.PodLogOptions{
			Follow:    true,
			TailLines: &line,
		})
	podlog, err := log.Stream(ctx)

	if err != nil {
		util.Logger.Errorf("watch.GetPodLog err: %s", err)
		return nil, err
	}

	return podlog, nil
}
