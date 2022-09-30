package svc

import (
	"austin-go/app/austin-job/internal/config"
	"austin-go/common/mq"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	MqClient    mq.IMessagingClient
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, err := mq.NewMessagingClientURL(c.Rabbit.URL)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:   c,
		MqClient: client,
		RedisClient: redis.New(c.CacheRedis[0].Host, func(r *redis.Redis) {
			r.Type = c.CacheRedis[0].Type
			r.Pass = c.CacheRedis[0].Pass
		}),
	}
}
