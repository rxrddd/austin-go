package send_account

import (
	"austin-go/app/austin-common/model"
	"context"
	"github.com/spf13/cast"

	"austin-go/app/austin-admin/api/internal/svc"
	"austin-go/app/austin-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendAccountEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendAccountEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) SendAccountEditLogic {
	return SendAccountEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendAccountEditLogic) SendAccountEdit(req types.SendAccountEditReq) error {
	return l.svcCtx.SendAccountRepo.Edit(l.ctx, &model.SendAccount{
		ID:          cast.ToInt64(req.ID),
		SendChannel: req.SendChannel,
		Config:      req.Config,
		Title:       req.Title,
	})
}
