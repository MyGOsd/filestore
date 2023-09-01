package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct {
		User     string
		Password string
		Address  string
		Dbname   string
	}
	Redis struct {
		Address     string
		Password    string
		DB          int
		PoolSize    int
		MinIdleConn int
	}
}
