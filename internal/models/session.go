package models

import "shortener/internal/types"

type Session struct {
	ID        uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID    uint       `gorm:"column:user_id;not null" json:"user_id"`
	IPAddress string     `gorm:"column:ip_address;type:varchar(45);not null" json:"ipAddress"`
	UserAgent string     `gorm:"column:user_agent;type:varchar(512);not null" json:"userAgent"`
	ExpiresAt types.Time `gorm:"column:expires_at;type:datetime;not null" json:"expiresAt"`
	CreatedAt types.Time `gorm:"column:created_at;type:datetime;autoCreateTime;not null" json:"createdAt"`
}

func (Session) TableName() string {
	return "sessions"
}
