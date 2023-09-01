package models

import (
	"cloud_disk/core/internal/config"
	"fmt"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(c config.Config) *gorm.DB {
	user := c.Mysql.User
	address := c.Mysql.Address
	password := c.Mysql.Password
	dbname := c.Mysql.Dbname
	dns := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, address, dbname)
	DB, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return DB
}

func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Address,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})

}
