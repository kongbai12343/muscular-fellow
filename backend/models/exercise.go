package models

import "time"

type Exercise struct {
	Id          int64     `json:"id" gorm:"primary_key;auto_increment;comment:'主键'"`
	UserId      int64     `json:"user_id" gorm:"type:bigint;not null;index;uniqueIndex:idx_exercises_user_name;comment:'用户id'"`
	Name        string    `json:"name" gorm:"type:varchar(50);not null;uniqueIndex:idx_exercises_user_name;comment:'名称'"`
	MuscleGroup int16     `json:"muscle_group" gorm:"type:smallint;default:0;comment:'肌群'"`
	Category    int16     `json:"category" gorm:"type:smallint;default:0;comment:'类别'"`
	Note        string    `json:"note" gorm:"type:varchar(300);comment:'备注'"`
	CreatedAt   time.Time `json:"created_at" gorm:"comment:'创建时间'"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"comment:'更新时间'"`
}

func (Exercise) TableName() string {
	return "exercises"
}
