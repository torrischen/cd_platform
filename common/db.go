package common

import (
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id"  gorm:"primarykey"`
	CreatedAt *DateTime      `json:"created_at,omitempty" gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP()"`
	UpdatedAt *DateTime      `json:"updated_at,omitempty" gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Projects struct {
	Model
	Name string `gorm:"unique" json:"name"`
}
