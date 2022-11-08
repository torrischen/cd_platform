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
	flag.StringVar(&Conf.MysqlAddr, "mysqlAddr", "localhost:3306", "mysqladdr")
	flag.StringVar(&Conf.MysqlUser, "mysqlUser", "root", "mysqluser")
	flag.StringVar(&Conf.MysqlPass, "mysqlPass", "123456", "mysqlpass")
	flag.StringVar(&Conf.MysqlDb, "mysqlDbName", "cd", "mysqldb")
	flag.StringVar(&Conf.HarborAddr, "harborAddr", "", "harboraddr")

	flag.Parse()
}
