package ext

import (
	"cd_platform/conf"
	"cd_platform/mid"
	"cd_platform/pkg/workload"
	"cd_platform/pkg/workload/watch"
	"cd_platform/util"
)

var MiddleWare *mid.Middle

func InitApp() {
	util.InitLogger()
	conf.InitConf()
	MiddleWare = mid.InitMiddleware(conf.Conf)
	workload.Init(MiddleWare)
	watch.Init(MiddleWare)
}
