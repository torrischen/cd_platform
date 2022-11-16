package common

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type InitProjectArgs struct {
	Name string `json:"name"`
}

type CreateApplicationArgs struct {
	Project       string             `json:"project"`
	DeploymentRaw *appsv1.Deployment `json:"deployment_raw"`
	ServiceRaw    *corev1.Service    `json:"service_raw"`
	IngressRule   IngressRule        `json:"ingress_rule"`
}

type IngressRule struct {
	Application string `json:"application"`
	Port        int32  `json:"port"`
}

type DestroyApplicationArgs struct {
	Project     string `json:"project"`
	Application string `json:"application"`
}

type DeployApplicationArgs struct {
	Project     string `json:"project"`
	Application string `json:"application"`
	Image       string `json:"image"`
}

type CreateSitApplicationArgs struct {
	Application   string             `json:"application"`
	DeploymentRaw *appsv1.Deployment `json:"deployment_raw"`
	ServiceRaw    *corev1.Service    `json:"service_raw"`
	IngressRule   IngressRule        `json:"ingress_rule"`
}

type DeploySitApplicationArgs struct {
	Application string `json:"application"`
	Image       string `json:"image"`
}

type DestroySitApplicationArgs struct {
	Application string `json:"application"`
}

type GetProjectListQueryArgs struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}
