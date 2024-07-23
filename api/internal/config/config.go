package config

import (
	"UBC/api/library/database"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	ORM database.MysqlConfig

	Address    []string
	Invoice    []string
	PythonPath string
}
