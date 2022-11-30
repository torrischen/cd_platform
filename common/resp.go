package common

import (
	corev1 "k8s.io/api/core/v1"
)

type ResponseBody struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ApplicationInfo struct {
	Name       string   `json:"name"`
	Ip         []string `json:"ip"`
	Status     string   `json:"status"`
	Replica    int      `json:"replica"`
	Path       string   `json:"path"`
	CreateTime string   `json:"create_time"`
}

type PodByApplication struct {
	Application string       `json:"application"`
	Pods        []*PodDetail `json:"pods"`
}

type PodDetail struct {
	Name       string                 `json:"name"`
	Namespace  string                 `json:"namespace"`
	Image      string                 `json:"image"`
	CreateTime string                 `json:"create_time"`
	HostIp     string                 `json:"host_ip"`
	PodIp      string                 `json:"pod_ip"`
	Status     *corev1.ContainerState `json:"status"`
}

type ConfigList struct {
	ConfigNameList []string `json:"config_name_list"`
}

type ConfigDetail struct {
	ConfigName  string `json:"config_name"`
	ConfigValue string `json:"config_value"`
}

type DeploymentItem struct {
	Name         string `json:"name"`
	Replica      int32  `json:"replica"`
	CreationTime string `json:"creation_time"`
	Generation   int64  `json:"generation"`
}
