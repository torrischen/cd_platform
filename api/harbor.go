package api

type ProjectListArgs struct {
}

type ProjectList struct {
	Name         string `json:"name"`
	OwnerName    string `json:"owner_name"`
	ProjectId    int    `json:"project_id"`
	OwnerId      int    `json:"owner_id"`
	CreationTime string `json:"creation_time"`
	UpdateTime   string `json:"update_time"`
	RepoCount    int    `json:"repo_count"`
}
