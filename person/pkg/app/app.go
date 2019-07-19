package app

import (
	"git.sdkeji.top/share/sdlib/log"
	"git.sdkeji.top/share/sdlib/mysql"
	"github.com/go-xorm/xorm"
	"sdkeji/person/pkg/resource"
)

var (
	Conf   Config
	DB     *xorm.Engine
	Logger log.Logger
	Redis  *RedisClient
)

func Init() {
	var err error
	Conf, err = NewConfig()
	if err != nil {
		panic(err)
	}

	{
		Logger, err = log.New(Conf.Debug, System)
		if err != nil {
			panic(err)
		}

		Logger.Info("logger ready.", "config", Conf)
	}

	{
		resource.Load()
	}

	{
		DB, err = mysql.OpenDB(Conf.Mysql, resource.MigrationBox, Logger)
		if err != nil {
			panic(err)
		}
		Logger.Info("db ready.")
	}

	{
		Redis, err = OpenRedis(Conf.Redis, 10, Logger)
		if err != nil {
			panic(err)
		}
		Logger.Info("redis ready.")
	}

}
