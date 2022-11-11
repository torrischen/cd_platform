package harbor

import (
	"cd_platform/common"
	"cd_platform/conf"
	"cd_platform/util"
	"errors"
	"github.com/google/uuid"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	HarborClient *resty.Client
	HarborAddr   string
}

func Init(conf conf.Config) *Client {
	return &Client{
		HarborClient: resty.New().SetBasicAuth(conf.HarborUser, conf.HarborPass),
		HarborAddr:   conf.HarborAddr,
	}
}

func (c *Client) ListRepo(project string) ([]*common.RepoItem, error) {
	var repolist []*common.RepoItem
	_, err := c.HarborClient.R().
		SetHeader("X-Request-Id", uuid.New().String()).
		SetResult(&repolist).
		Get(c.HarborAddr + "/projects" + "/" + project + "/repositories")
	if err != nil {
		util.Logger.Errorf("harbor.ListRepo err: %s", err)
		return nil, err
	}

	return repolist, nil
}

func (c *Client) GetRepoTag(project string, repo string) (*common.ImageList, error) {
	var taglist []*common.TagList
	_, err := c.HarborClient.R().
		SetHeader("X-Request-Id", uuid.New().String()).
		SetHeader("X-Accept-Vulnerabilities", "application/vnd.scanner.adapter.vuln.report.harbor+json; version=1.0").
		SetQueryParam("with_tag", "true").
		SetResult(&taglist).
		Get(c.HarborAddr + "/projects" + "/" + project + "/repositories" + "/" + repo + "/artifacts")
	if err != nil {
		util.Logger.Errorf("harbor.GetRepoTag err: %s", err)
		return nil, err
	}

	var imglist common.ImageList
	for i := 0; i < len(taglist); i++ {
		for j := 0; j < len(taglist[i].Tags); j++ {
			tmp := "harbor.devops" + "/" + project + "/" + repo + ":" + taglist[i].Tags[j].Name
			imglist.Data = append(imglist.Data, tmp)
		}
	}

	return &imglist, nil
}

func (c *Client) CreateProject(project string) error {
	args := &common.CreateHarborProjectArgs{
		ProjectName: project,
		Metadata: &common.CreateHarborProjectMetadata{
			Public: "true",
		},
		StorageLimit: 10737418240,
		RegistryId:   nil,
	}

	resp, err := c.HarborClient.R().
		SetHeader("X-Request-Id", uuid.New().String()).
		SetBody(args).
		Post(c.HarborAddr + "/projects")
	if err != nil {
		util.Logger.Errorf("harbor.CreateProject err: %s", err)
		return err
	}
	if resp.StatusCode() != 200 {
		util.Logger.Errorf("harbor.CreateProject err: %s", string(resp.Body()))
		return errors.New(string(resp.Body()))
	}

	return nil
}
