package mid

import (
	"cd_platform/conf"
	"cd_platform/mid/k8sclient"
)

type Middle struct {
	K8sclient *k8sclient.Client
}

func InitMiddleware(conf conf.Config) *Middle {
	mid := &Middle{}
	mid.K8sclient = k8sclient.Init(conf)

	return mid
}
