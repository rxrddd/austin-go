package send_account

import (
	"austin-go/common/zcopier"
	"context"

	"austin-go/app/austin-admin/api/internal/svc"
	"austin-go/app/austin-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendAccountListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) SendAccountListLogic {
	return SendAccountListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendAccountListLogic) SendAccountList(req types.SendAccountListReq) (resp *types.SendAccountListResp, err error) {

	resp = new(types.SendAccountListResp)
	all, err := l.svcCtx.SendAccountRepo.FindAll(l.ctx, req)
	if err != nil {
		return nil, err
	}
	resp.Items = make([]types.SendAccountItem, 0)
	zcopier.Copy(&resp.Items, &all)
	return resp, nil
}
