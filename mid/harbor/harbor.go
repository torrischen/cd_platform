package harbor

import (
	"cd_platform/api"
	"cd_platform/conf"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	HarborClient *resty.Client
	HarborAddr   string
}

func Init(conf conf.Config) *Client {
	return &Client{
		HarborClient: resty.New(),
		HarborAddr:   conf.HarborAddr,
	}
}

func (c *Client) ListProject() *api.ProjectList {
	return nil
}
