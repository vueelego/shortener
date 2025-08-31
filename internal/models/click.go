package models

import (
	"shortener/internal/types"
)

type Click struct {
	ID         uint       `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	EntryID    uint       `gorm:"column:entry_id;not null" json:"entryId"`
	IPAddress  string     `gorm:"column:ip_address;type:varchar(45);not null" json:"ipAddress"`
	UserAgent  string     `gorm:"column:user_agent;type:varchar(512);not null" json:"userAgent"`
	DeviceType string     `gorm:"column:device_type;type:varchar(32);not null" json:"deviceType"`
	CreatedAt  types.Time `gorm:"column:created_at;not null" json:"createdAt"`
}

func (Click) TableName() string {
	return "clicks"
}
