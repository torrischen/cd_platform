package common

type RepoItem struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	CreationTime string `json:"creation_time"`
	UpdateTime   string `json:"update_time"`
}

type TagItem struct {
	Name string `json:"name"`
}

type TagList struct {
	Tags []*TagItem `json:"tags"`
}

type ImageList struct {
	Data []string `json:"data"`
}

type CreateHarborProjectArgs struct {
	ProjectName  string                       `json:"project_name"`
	Metadata     *CreateHarborProjectMetadata `json:"metadata"`
	StorageLimit int64                        `json:"storage_limit"`
	RegistryId   *int                         `json:"registry_id"`
}

type CreateHarborProjectMetadata struct {
	Public string `json:"public"`
}
