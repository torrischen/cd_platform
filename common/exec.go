package common

type InitProjectArgs struct {
	Name string `json:"name"`
}

type CreateProjectArgs struct {
	Project       string      `json:"project"`
	DeploymentRaw []byte      `json:"deployment_raw"`
	ServiceRaw    []byte      `json:"service_raw"`
	IngressRule   IngressRule `json:"ingress_rule"`
}

type IngressRule struct {
	Application string `json:"application"`
	Port        int32  `json:"port"`
}

type DestroyProjectArgs struct {
	Project     string `json:"project"`
	Application string `json:"application"`
}

type DeployApplicationArgs struct {
	Project     string `json:"project"`
	Application string `json:"application"`
	Image       string `json:"image"`
}
