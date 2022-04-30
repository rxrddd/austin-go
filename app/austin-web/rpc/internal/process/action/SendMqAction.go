package action

import (
	"austin-go/common/mq"
	"austin-go/common/zutils/dd"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type SendMqAction struct {
	mqClient mq.IMessagingClient
}

func NewSendMqAction(mqClient mq.IMessagingClient) *SendMqAction {
	return &SendMqAction{mqClient: mqClient}
}

func (p SendMqAction) Process(ctx context.Context, data interface{}) error {
	dd.Print(data)
	//sendTaskModel, ok := data.(*types.SendTaskModel)
	//if !ok {
	//	return errors.Wrapf(sendErr, "AssembleAction 类型错误 err:%v", data)
	//}
	marshal, err := jsonx.Marshal(data)
	if err != nil {
		return err
	}
	return p.mqClient.Publish(marshal, "austin.biz")
}
