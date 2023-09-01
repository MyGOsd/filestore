package models

import (
	"time"

	"gorm.io/gorm"
)

type ShareBasic struct {
	Id                     int
	Identity               string
	UserRepositoryIdentity string
	UserIdentity           string
	RepositoryIdentity     string
	ExpiredTime            int
	ClickNum               int
	Created_at             time.Time
	Updated_at             time.Time
	Deleted_at             gorm.DeletedAt
}

func (*ShareBasic) TableName() string {
	return "share_basic"
}
