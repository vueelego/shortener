package models

import (
	"shortener/internal/types"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	CreatedAt types.Time     `gorm:"autoCreateTime;not null" json:"createdAt"`
	UpdatedAt types.Time     `gorm:"autoUpdateTime;not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
