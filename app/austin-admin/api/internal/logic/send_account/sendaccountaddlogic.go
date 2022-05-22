package send_account

import (
	"austin-go/app/austin-common/model"
	"context"

	"austin-go/app/austin-admin/api/internal/svc"
	"austin-go/app/austin-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendAccountAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendAccountAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) SendAccountAddLogic {
	return SendAccountAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendAccountAddLogic) SendAccountAdd(req types.SendAccountAddReq) error {
	return l.svcCtx.SendAccountRepo.Create(l.ctx, &model.SendAccount{
		SendChannel: req.SendChannel,
		Config:      req.Config,
		Title:       req.Title,
	})
}
