package svc

import (
	"austin-go/app/austin-web/rpc/internal/config"
	"austin-go/common/mq"
)

type ServiceContext struct {
	Config   config.Config
	MqClient mq.IMessagingClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, err := mq.NewMessagingClientURL(c.Rabbit.URL)
	handleErr(err)
	return &ServiceContext{
		Config:   c,
		MqClient: client,
	}
}
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
