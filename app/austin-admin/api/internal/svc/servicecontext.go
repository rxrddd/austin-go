package svc

import (
	"austin-go/app/austin-admin/api/internal/config"
	"austin-go/app/austin-admin/api/internal/middleware"
	"austin-go/app/austin-admin/api/internal/repo"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware

	AccountRepo     *repo.AccountRepo
	SendAccountRepo *repo.SendAccountRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		Auth:            middleware.NewAuthMiddleware().Handle,
		AccountRepo:     repo.NewAccountRepo(),
		SendAccountRepo: repo.NewSendAccountRepo(),
	}
}
