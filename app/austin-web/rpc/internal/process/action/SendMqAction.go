package action

import (
	"austin-go/app/austin-common/enums/channelType"
	"austin-go/app/austin-common/enums/messageType"
	"austin-go/app/austin-common/types"
	"austin-go/common/mq"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type SendMqAction struct {
	mqClient mq.IMessagingClient
}

func NewSendMqAction(mqClient mq.IMessagingClient) *SendMqAction {
	return &SendMqAction{mqClient: mqClient}
}

func (p SendMqAction) Process(_ context.Context, sendTaskModel *types.SendTaskModel) error {
	marshal, err := jsonx.Marshal(sendTaskModel)
	if err != nil {
		return err
	}
	channel := channelType.TypeCodeEn[sendTaskModel.TaskInfo[0].SendChannel]
	msgType := messageType.TypeCodeEn[sendTaskModel.TaskInfo[0].MsgType]
	return p.mqClient.Publish(marshal, fmt.Sprintf("austin.biz.%s.%s", channel, msgType))
}
