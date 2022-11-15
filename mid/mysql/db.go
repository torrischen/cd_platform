package mysql

import (
	"cd_platform/common"
	"cd_platform/conf"
	"cd_platform/util"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
)

type Client struct {
	Db *gorm.DB
}

type List struct {
	Page       int           `json:"page"`
	PageSize   int           `json:"page_size"`
	TotalPage  int           `json:"total_page"`
	TotalCount int64         `json:"total_count"`
	Total      []interface{} `json:"total,omitempty"`
	Data       interface{}   `json:"list"`
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

func (c *Client) GetProjectList(page int, pagesize int) (*List, error) {
	var ret []*common.Projects
	list, err := c.GetList(&ret, c.Db.Model(&common.Projects{}), page, pagesize, "")
	if err != nil {
		util.Logger.Errorf("mysql.GetProjectList err: %s", err)
		return nil, err
	}

	return list, nil
}

func (c *Client) GetList(data interface{}, cond interface{}, page, pageSize int, order string, fields ...string) (*List, error) {
	var count int64

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if order == "" {
		order = "-id"
	}

	list := &List{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: count,
	}
	//支持自定义查询
	var db *gorm.DB
	switch cond.(type) {
	case *gorm.DB:
		db = cond.(*gorm.DB)
		if db == nil {
			db = c.Db
		}
	case string:
		db = c.Db.Where(cond)
	default:
		db = c.Db
	}
	if len(fields) > 0 {
		for i, s := range fields {
			if strings.Index(s, "`") == -1 {
				if strings.Index(s, ".") != -1 {
					s = strings.Replace(s, ".", "`.`", -1)
				}
				fields[i] = "`" + s + "`"
			}
		}
		db = db.Select(fields)
	}

	if db.Statement.Model == nil {
		db = db.Model(data)
	}
	if len(db.Statement.Selects) == 0 {

	}

	if err := db.Count(&list.TotalCount).Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order(order).
		Find(data).Error; err != nil {
		return list, err
	}
	list.Data = data
	list.TotalPage = int(list.TotalCount) / pageSize
	if int(list.TotalCount)%list.PageSize > 0 {
		list.TotalPage += 1
	}
	return list, nil
}
