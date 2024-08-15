package domain

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	Id int `gorm:"primary_key;column:id;autoIncrement"`
	Title string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	UserId int64		`gorm:"column:user_id"`
	User User 			`gorm:"foreignKey:user_id;references:id"`
}

func (n *Note) TableName()string {
	return "notes"
}