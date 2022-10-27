package conf

import "flag"

type Config struct {
	KubecfgDir string
}

var Conf Config

func InitConf() {
	flag.StringVar(&Conf.KubecfgDir, "kubecfgdir", "./conf/kubecfg.yaml", "kubecfg")

	flag.Parse()
}
