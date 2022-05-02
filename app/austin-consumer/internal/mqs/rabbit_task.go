package mqs

import (
	"austin-go/app/austin-common/enums/channelType"
	"austin-go/app/austin-common/enums/messageType"
	"austin-go/app/austin-common/taskUtil"
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-consumer/internal/svc"
	"austin-go/app/austin-handler/pending"
	"austin-go/common/mq"
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type RabbitTask struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	client mq.IMessagingClient
}

func NewRabbitTask(ctx context.Context, svcCtx *svc.ServiceContext) *RabbitTask {
	client, err := mq.NewMessagingClientURL(svcCtx.Config.Rabbit.URL)
	if err != nil {
		panic(err)
	}
	return &RabbitTask{
		ctx:    ctx,
		svcCtx: svcCtx,
		client: client,
	}
}

func (l *RabbitTask) Start() {

	fmt.Println("RabbitTask start ")
	for _, groupId := range taskUtil.GetAllGroupIds() {
		_ = l.client.Subscribe(fmt.Sprintf("austin.biz.%s", groupId), onMassage)
	}
	select {}
}

func (l *RabbitTask) Stop() {
	fmt.Println("RabbitTask stop ")
	l.client.Close()
}

func onMassage(delivery amqp.Delivery) {
	ctx := context.Background()
	var SendTaskModel types.SendTaskModel
	_ = jsonx.Unmarshal(delivery.Body, &SendTaskModel)
	for _, taskInfo := range SendTaskModel.TaskInfo {
		logx.WithContext(ctx).Infof("消息接收成功,开始消费,内容: %s", string(delivery.Body))
		channel := channelType.TypeCodeEn[taskInfo.SendChannel]
		msgType := messageType.TypeCodeEn[taskInfo.MsgType]
		err := pending.Submit(ctx, fmt.Sprintf("%s.%s", channel, msgType), pending.Task{TaskInfo: taskInfo})
		if err != nil {
			logx.WithContext(ctx).Errorf("submit task err:%v ,内容: %s", err, string(delivery.Body))
		}
	}
}
