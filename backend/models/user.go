package models

import "time"

// User 用户表
type User struct {
	Id        int64     `gorm:"primary_key;auto_increment;comment:'主键'" json:"id"`
	Username  string    `gorm:"type:varchar(100);not null;uniqueIndex;comment:'用户名，唯一'" json:"username"`
	Email     string    `gorm:"type:varchar(100);not null;uniqueIndex;comment:'邮箱，唯一'" json:"email"`
	Password  string    `gorm:"type:varchar(100);not null;comment:'密码哈希'" json:"-"`
	CreatedAt time.Time `gorm:"comment:'创建时间'" json:"created_at"`
	UpdatedAt time.Time `gorm:"comment:'更新时间'" json:"updated_at"`
}

// TableName 指定表名
func (User) TableName() string { return "users" }
