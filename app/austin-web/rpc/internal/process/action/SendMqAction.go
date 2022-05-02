package action

import (
	"austin-go/app/austin-common/enums/channelType"
	"austin-go/app/austin-common/enums/messageType"
	"austin-go/app/austin-common/taskUtil"
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-web/rpc/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type SendMqAction struct {
	svcCtx *svc.ServiceContext
}

func NewSendMqAction(svcCtx *svc.ServiceContext) *SendMqAction {
	return &SendMqAction{svcCtx: svcCtx}
}

func (p SendMqAction) Process(_ context.Context, sendTaskModel *types.SendTaskModel) error {
	marshal, err := jsonx.Marshal(sendTaskModel.TaskInfo)
	if err != nil {
		return err
	}
	channel := channelType.TypeCodeEn[sendTaskModel.TaskInfo[0].SendChannel]
	msgType := messageType.TypeCodeEn[sendTaskModel.TaskInfo[0].MsgType]
	return p.svcCtx.MqClient.Publish(marshal, taskUtil.GetMqKey(channel, msgType))
}
