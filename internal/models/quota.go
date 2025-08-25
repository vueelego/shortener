package models

import "shortener/internal/types"

type Quota struct {
	ID           uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID       uint       `gorm:"column:user_id;not null" json:"userId"`
	MonthlyLimit int        `gorm:"column:monthly_limit;comment:'每月限制数量'; not null" json:"monthlyLimit"`
	UsedMonth    int        `gorm:"column:used_month;default:0;comment:'每月已使用数量';not null" json:"usedMonth"`
	LastResetAt  types.Time `gorm:"column:last_reset_at;autoCreateTime;comment:'上次重置时间';not null" json:"lastResetAt"`
}

func (Quota) TableName() string {
	return "quotas"
}
