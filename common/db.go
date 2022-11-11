package common

import "gorm.io/gorm"

type Projects struct {
	gorm.Model
	Name string `gorm:"unique" json:"name"`
}
