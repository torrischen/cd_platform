package common

type Labels struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Annotations struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type NewDeploymentArgs struct {
	OwnerName      string        `json:"owner_name"`
	DeploymentName string        `json:"deployment_name"`
	Labels         []Labels      `json:"labels"`
	Annotations    []Annotations `json:"annotations"`
}
