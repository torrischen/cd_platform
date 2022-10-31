package conf

import "flag"

type Config struct {
	KubecfgDir string
	MysqlAddr  string
	MysqlUser  string
	MysqlPass  string
	MysqlDb    string
	HarborAddr string
}

var Conf Config

func InitConf() {
	flag.StringVar(&Conf.KubecfgDir, "kubeConfigDir", "./conf/kubecfg.yaml", "kubecfg")
	flag.StringVar(&Conf.MysqlAddr, "mysqlAddr", "", "mysqladdr")
	flag.StringVar(&Conf.MysqlUser, "mysqlUser", "root", "mysqluser")
	flag.StringVar(&Conf.MysqlPass, "mysqlPass", "", "mysqlpass")
	flag.StringVar(&Conf.MysqlDb, "mysqlDbName", "", "mysqldb")
	flag.StringVar(&Conf.HarborAddr, "harborAddr", "", "harboraddr")

	flag.Parse()
}
