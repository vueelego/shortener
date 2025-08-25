package models

import (
	"shortener/internal/types"
)

type Entry struct {
	Model
	UserID      uint       `gorm:"column:user_id;not null" json:"userId"`
	ShortCode   string     `gorm:"column:short_code;type:varchar(8);uniqueIndex;not null" json:"shortCode"`
	OriginalUrl string     `gorm:"column:original_url;type:varchar(2048);not null" json:"originalUrl"`
	Title       string     `gorm:"column:title;type:varchar(32);not null" json:"title"`
	ExpiresAt   types.Time `gorm:"column:expires_at;type:datetime;not null" json:"expiresAt"`
}

func (Entry) TableName() string {
	return "entries"
}
