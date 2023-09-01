package models

import (
	"time"

	"gorm.io/gorm"
)

type RepositoryPool struct {
	Id         int
	Identity   string
	Hash       string
	Name       string
	Ext        string
	Size       int64
	Path       string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt
}

func (*RepositoryPool) TableName() string {
	return "repository_Pool"
}
