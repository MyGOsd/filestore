package models

import (
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	Id         int
	Identity   string
	Name       string
	Password   string
	Email      string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
