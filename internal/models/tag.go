package models

type Tag struct {
	ID      uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OwnerID uint   `gorm:"column:owner_id;not null" json:"ownerId"`
	TagName string `gorm:"column:tag_name;type:varchar(32);not null" json:"tagName"`
}

func (Tag) TableName() string {
	return "tags"
}
