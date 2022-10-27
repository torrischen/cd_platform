package mysql

import (
	"cd_platform/conf"
	"cd_platform/util"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Client struct {
	Db *gorm.DB
}

func Init(conf conf.Config) *Client {
	c := &Client{}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.MysqlUser, conf.MysqlPass, conf.MysqlAddr, conf.MysqlDb)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		util.Logger.Fatalf("Fail to init mysql: %s", err)
		return nil
	}

	c.Db = db

	return c
}
