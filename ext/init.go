package ext

import (
	"cd_platform/conf"
	"cd_platform/mid"
	"cd_platform/pkg"
	"cd_platform/util"
)

var MiddleWare *mid.Middle

func InitApp() {
	util.InitLogger()
	conf.InitConf()
	MiddleWare = mid.InitMiddleware(conf.Conf)
	pkg.NewService(MiddleWare)
}
