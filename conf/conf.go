package conf

import "flag"

type Config struct {
	KubecfgDir string
	MysqlAddr  string
	MysqlUser  string
	MysqlPass  string
	MysqlDb    string
	HarborAddr string
	HarborUser string
	HarborPass string
}

var Conf Config

func InitConf() {
	flag.StringVar(&Conf.KubecfgDir, "kubeConfigDir", "./conf/kubecfg.yaml", "kubecfg")
	flag.StringVar(&Conf.MysqlAddr, "mysqlAddr", "192.168.3.51:3306", "mysqladdr")
	flag.StringVar(&Conf.MysqlUser, "mysqlUser", "root", "mysqluser")
	flag.StringVar(&Conf.MysqlPass, "mysqlPass", "123456", "mysqlpass")
	flag.StringVar(&Conf.MysqlDb, "mysqlDbName", "cd", "mysqldb")
	flag.StringVar(&Conf.HarborAddr, "harborAddr", "http://192.168.3.250/api/v2.0", "harboraddr")
	flag.StringVar(&Conf.HarborUser, "harborUser", "admin", "harboruser")
	//flag.StringVar(&Conf.HarborUser, "harborUser", "robot$cd", "harboruser")
	flag.StringVar(&Conf.HarborPass, "harborPass", "123..COM", "harborpass")
	//flag.StringVar(&Conf.HarborPass, "harborPass", "6FlWSYIITCh0OdoMjRxB3Rig9cM4WpTG", "harborpass")

	flag.Parse()
}
