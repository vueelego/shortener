package models

type User struct {
	Model
	Username     string `gorm:"column:username;type:varchar(32);not null" json:"username"`
	Email        string `gorm:"column:email;type:varchar(100);uniqueIndex;not null" json:"email"`
	PasswordHash string `gorm:"column:password_hash;type:varchar(64);not null" json:"-"`
	Avatar       string `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`
	Role         bool   `gorm:"column:role;default:0;comment:'默认0,管理员1';not null" json:"role"`
}

func (User) TableName() string {
	return "users"
}
