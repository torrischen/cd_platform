package k8sclient

import (
	"cd_platform/conf"
	"cd_platform/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	listerappsv1 "k8s.io/client-go/listers/apps/v1"
	listercorev1 "k8s.io/client-go/listers/core/v1"
	listernetworkv1 "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	"time"
)

type Client struct {
	ClientSet         *kubernetes.Clientset
	DynamicClient     dynamic.Interface
	DeploymentLister  listerappsv1.DeploymentLister
	StatefulSetLister listerappsv1.StatefulSetLister
	PodLister         listercorev1.PodLister
	ServiceLister     listercorev1.ServiceLister
	IngressLister     listernetworkv1.IngressLister
}

func Init(conf conf.Config) *Client {
	c := &Client{}
	config, err := clientcmd.BuildConfigFromFlags("", conf.KubecfgDir)
	if err != nil {
		util.Logger.Fatalf("mid.Init clientSet failed: %s", err)
		return nil
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		util.Logger.Fatalf("mid.Init clientSet failed: %s", err)
		return nil
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		util.Logger.Fatalf("mid.Init dynamicClient failed: %s", err)
		return nil
	}

	sharedIM := informers.NewSharedInformerFactory(clientset, 30*time.Second)

	sharedIM.Apps().V1().Deployments().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			util.Logger.Infof("New deployemnt added: %s", obj.(metav1.Object).GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oobj := oldObj.(metav1.Object)
			nobj := newObj.(metav1.Object)
			util.Logger.Infof("deployment update, %s has been updated to %s", oobj.GetName(), nobj.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			util.Logger.Infof("deployment deleted: %s", obj.(metav1.Object).GetName())
		},
	})

	sharedIM.Apps().V1().StatefulSets().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			util.Logger.Infof("New statefulset added: %s", obj.(metav1.Object).GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oobj := oldObj.(metav1.Object)
			nobj := newObj.(metav1.Object)
			util.Logger.Infof("statefulset update, %s has been updated to %s", oobj.GetName(), nobj.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			util.Logger.Infof("statefulset deleted: %s", obj.(metav1.Object).GetName())
		},
	})

	sharedIM.Core().V1().Pods().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			util.Logger.Infof("New pod added: %s", obj.(metav1.Object).GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oobj := oldObj.(metav1.Object)
			nobj := newObj.(metav1.Object)
			util.Logger.Infof("pod updated, %s has been updated to %s", oobj.GetName(), nobj.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			util.Logger.Infof("pod deleted: %s", obj.(metav1.Object).GetName())
		},
	})

	sharedIM.Core().V1().Services().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			util.Logger.Infof("New service added: %s", obj.(metav1.Object).GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oobj := oldObj.(metav1.Object)
			nobj := newObj.(metav1.Object)
			util.Logger.Infof("service updated, %s has been updated to %s", oobj.GetName(), nobj.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			util.Logger.Infof("service deleted: %s", obj.(metav1.Object).GetName())
		},
	})

	sharedIM.Networking().V1().Ingresses().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			util.Logger.Infof("New ingress added: %s", obj.(metav1.Object).GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oobj := oldObj.(metav1.Object)
			nobj := newObj.(metav1.Object)
			util.Logger.Infof("ingress updated, %s has been updated to %s", oobj.GetName(), nobj.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			util.Logger.Infof("ingress deleted: %s", obj.(metav1.Object).GetName())
		},
	})

	depLister := sharedIM.Apps().V1().Deployments().Lister()
	stsLister := sharedIM.Apps().V1().StatefulSets().Lister()
	podLister := sharedIM.Core().V1().Pods().Lister()
	serviceLister := sharedIM.Core().V1().Services().Lister()
	ingressLister := sharedIM.Networking().V1().Ingresses().Lister()

	go sharedIM.Start(make(chan struct{}))

	c.ClientSet = clientset
	c.DynamicClient = dynamicClient
	c.DeploymentLister = depLister
	c.StatefulSetLister = stsLister
	c.PodLister = podLister
	c.ServiceLister = serviceLister
	c.IngressLister = ingressLister

	return c
}
