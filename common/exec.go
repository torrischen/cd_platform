package common

type InitProjectArgs struct {
	Project string `json:"project"`
}

type DeployProjectArgs struct {
	DeploymentRaw []byte        `json:"deployment_raw"`
	ServiceRaw    []byte        `json:"service_raw"`
	IngressRules  []IngressRule `json:"ingress_rules"`
}

type IngressRule struct {
}
