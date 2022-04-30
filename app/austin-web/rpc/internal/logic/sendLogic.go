package logic

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-web/rpc/austin"
	"austin-go/app/austin-web/rpc/internal/process"
	"austin-go/app/austin-web/rpc/internal/process/action"
	"austin-go/app/austin-web/rpc/internal/process/interfaces"
	"austin-go/app/austin-web/rpc/internal/svc"
	"austin-go/common/xerr"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendLogic) Send(in *austin.SendRequest) (*austin.SendResponse, error) {

	if in.MessageParam == nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("客户端参数错误"), "in:%v", in)
	}

	var sendTaskModel = &types.SendTaskModel{
		MessageTemplateId: in.MessageTemplateId,
		MessageParamList: []types.MessageParam{
			{
				Receiver:  in.MessageParam.Receiver,
				Variables: in.MessageParam.Variables,
				Extra:     in.MessageParam.Extra,
			},
		},
	}
	businessProcess := process.NewBusinessProcess()
	_ = businessProcess.AddProcess([]interfaces.Process{
		action.NewPreParamCheckAction(),           //前置参数校验
		action.NewAssembleAction(),                //拼装参数
		action.NewAfterParamCheckAction(),         //后置参数检查
		action.NewSendMqAction(l.svcCtx.MqClient), //发送到mq
	}...)

	err := businessProcess.Process(l.ctx, sendTaskModel)
	if err != nil {
		return nil, err
	}
	return &austin.SendResponse{Code: in.Code}, nil
}
