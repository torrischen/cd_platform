package conf

import "flag"

type Config struct {
	KubecfgDir string
	MysqlAddr  string
	MysqlUser  string
	MysqlPass  string
	MysqlDb    string
}

var Conf Config

func InitConf() {
	flag.StringVar(&Conf.KubecfgDir, "kubecfgdir", "./conf/kubecfg.yaml", "kubecfg")
	flag.StringVar(&Conf.MysqlAddr, "mysqladdr", "", "mysqladdr")
	flag.StringVar(&Conf.MysqlUser, "mysqluser", "root", "mysqluser")
	flag.StringVar(&Conf.MysqlPass, "mysqlpass", "", "mysqlpass")
	flag.StringVar(&Conf.MysqlDb, "mysqldb", "", "mysqldb")

	flag.Parse()
}
