package svc

import (
	"austin-go/app/austin-web/api/internal/config"
	"austin-go/app/austin-web/rpc/austin"
	"austin-go/app/austin-web/rpc/austinclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	SendRpc austin.AustinClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		SendRpc: austinclient.NewAustin(zrpc.MustNewClient(c.SendRpcConf)),
	}
}
