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
}

type IngressRule struct {
	Project     string `json:"project"`
	Application string `json:"application"`
	Path        string `json:"path"`
	Port        int32  `json:"port"`
}

type SitIngressRule struct {
	Application string `json:"application"`
	Path        string `json:"path"`
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
	Project       string             `json:"project"`
	Application   string             `json:"application"`
	DeploymentRaw *appsv1.Deployment `json:"deployment_raw"`
	ServiceRaw    *corev1.Service    `json:"service_raw"`
}

type DeploySitApplicationArgs struct {
	Project     string `json:"project"`
	Application string `json:"application"`
	Image       string `json:"image"`
}

type DestroySitApplicationArgs struct {
	Project     string `json:"project"`
	Application string `json:"application"`
}

type GetProjectListQueryArgs struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type DeleteSpecifiedIngressRuleArgs struct {
	Project string `json:"project"`
	Path    string `json:"path"`
}

type PatchReplicaArgs struct {
	Project     string `json:"project"`
	Application string `json:"application"`
	Replica     int32  `json:"replica"`
}

type RestartDeploymentArgs struct {
	Project     string `json:"project"`
	Application string `json:"application"`
}

type SetEnvArgs struct {
	Project     string          `json:"project"`
	Application string          `json:"application"`
	Envs        []corev1.EnvVar `json:"envs"`
}

type CreateConfigmapArgs struct {
	Project     string `json:"project"`
	Application string `json:"application"`
}

type AddConfigToConfigmapArgs struct {
	Project     string   `json:"project"`
	Application string   `json:"application"`
	Configs     []Config `json:"configs"`
}

type Config struct {
	ConfigName string `json:"config_name"`
	Data       string `json:"data"`
}

type DeleteSpecifiedConfigArgs struct {
	Project     string `json:"project"`
	Application string `json:"application"`
	ConfigName  string `json:"config_name"`
}

type UpdateSpecifiedConfigArgs struct {
	Project     string `json:"project"`
	Application string `json:"application"`
	ConfigName  string `json:"config_name"`
	NewVal      string `json:"new_val"`
}
