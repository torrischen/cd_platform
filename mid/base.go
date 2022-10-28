package mid

import (
	"cd_platform/conf"
	"cd_platform/mid/harbor"
	"cd_platform/mid/k8sclient"
	"cd_platform/mid/mysql"
)

type Middle struct {
	K8sclient    *k8sclient.Client
	MysqlClient  *mysql.Client
	HarborClient *harbor.Client
}

func InitMiddleware(conf conf.Config) *Middle {
	mid := &Middle{}
	mid.K8sclient = k8sclient.Init(conf)
	mid.MysqlClient = mysql.Init(conf)
	mid.HarborClient = harbor.Init(conf)

	return mid
}
