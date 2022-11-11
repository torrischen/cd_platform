package common

type InitProjectArgs struct {
	Project string `json:"project"`
}

type DeployProjectArgs struct {
	Project       string      `json:"project"`
	DeploymentRaw []byte      `json:"deployment_raw"`
	ServiceRaw    []byte      `json:"service_raw"`
	IngressRule   IngressRule `json:"ingress_rule"`
}

type IngressRule struct {
	Application string `json:"application"`
	Port        int32  `json:"port"`
}
