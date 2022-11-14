package mysql

import (
	"cd_platform/common"
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
	if err := db.AutoMigrate(&common.Projects{}); err != nil {
		util.Logger.Fatalf("Fail to init mysql: %s", err)
		return nil
	}

	c.Db = db

	return c
}

func (c *Client) CreateProject(project string) error {
	tmp := &common.Projects{}
	tmp.Name = project
	if err := c.Db.Create(tmp).Error; err != nil {
		util.Logger.Errorf("mysql.CreateProject err: %s", err)
		return err
	}

	return nil
}

func (c *Client) GetProjectList() ([]*common.Projects, error) {
	var ret []*common.Projects
	if err := c.Db.Model(&common.Projects{}).Find(&ret).Error; err != nil {
		util.Logger.Errorf("mysql.GetProjectList err: %s", err)
		return nil, err
	}

	return ret, nil
}
