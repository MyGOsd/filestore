package svc

import (
	"cloud_disk/core/internal/config"
	"cloud_disk/core/internal/middleware"
	"cloud_disk/core/models"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	RedisDB *redis.Client
	Auth    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DB:      models.Init(c),
		RedisDB: models.InitRedis(c),
		Auth:    middleware.NewAuthMiddleware().Handle,
	}
}
