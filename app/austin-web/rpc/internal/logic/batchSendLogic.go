package logic

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-web/rpc/internal/process"
	"austin-go/app/austin-web/rpc/internal/process/action"
	"austin-go/app/austin-web/rpc/internal/process/interfaces"
	"austin-go/common/xerr"
	"context"
	"github.com/pkg/errors"

	"austin-go/app/austin-web/rpc/austin"
	"austin-go/app/austin-web/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchSendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchSendLogic {
	return &BatchSendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchSendLogic) BatchSend(in *austin.BatchSendRequest) (*austin.SendResponse, error) {
	if in.MessageParam == nil || len(in.MessageParam) <= 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("客户端参数错误"), "in:%v", in)
	}
	var messageParamList = make([]types.MessageParam, 0)
	for _, item := range in.MessageParam {
		messageParamList = append(messageParamList, types.MessageParam{
			Receiver:  item.Receiver,
			Variables: item.Variables,
			Extra:     item.Extra,
		})
	}
	var sendTaskModel = &types.SendTaskModel{
		MessageTemplateId: in.MessageTemplateId,
		MessageParamList:  messageParamList,
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
