package config

import (
	"austin-go/common/mq"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Rabbit mq.RabbitConf
}
