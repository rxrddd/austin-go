package logic

import (
	"austin-go/app/austin-web/api/internal/svc"
	"austin-go/app/austin-web/api/internal/types"
	"austin-go/app/austin-web/rpc/austin"
	"austin-go/common/xerr"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) SendLogic {
	return SendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendLogic) Send(req types.SendRequest) (resp *types.Response, err error) {
	send, err := l.svcCtx.SendRpc.Send(l.ctx, &austin.SendRequest{
		Code:              req.Code,
		MessageTemplateId: req.MessageTemplateId,
		MessageParam: &austin.MessageParam{
			Receiver:  req.MessageParam.Receiver,
			Variables: req.MessageParam.Variables,
			Extra:     req.MessageParam.Extra,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("发送消息错误"), "send err:%v", err)
	}
	return &types.Response{Message: send.Code}, err
}
