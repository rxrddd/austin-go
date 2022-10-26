package config

import (
	"austin-go/common/dbx"
	"austin-go/common/mq"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type Config struct {
	Rabbit     mq.RabbitConf
	CacheRedis cache.CacheConf // redis缓存
	Mysql      dbx.DbConfig
}
