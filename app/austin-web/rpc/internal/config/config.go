package config

import (
	"austin-go/common/dbx"
	"austin-go/common/mq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql  dbx.DbConfig
	Rabbit mq.RabbitConf
}
