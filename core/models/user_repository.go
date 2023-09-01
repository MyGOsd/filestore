package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	Id                 int
	Identity           string
	UserIdentity       string
	ParentId           int64
	RepositoryIdentity string
	Ext                string
	Name               string
	Created_at         time.Time
	Updated_at         time.Time
	Deleted_at         gorm.DeletedAt
}

func (*UserRepository) TableName() string {
	return "user_repository"
}
