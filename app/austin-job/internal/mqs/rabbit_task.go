package mqs

import (
	"austin-go/app/austin-common/enums/channelType"
	"austin-go/app/austin-common/enums/messageType"
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/taskUtil"
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/pending"
	"austin-go/app/austin-job/internal/svc"
	"austin-go/common/dbx"
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type RabbitTask struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRabbitTask(ctx context.Context, svcCtx *svc.ServiceContext) *RabbitTask {

	return &RabbitTask{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RabbitTask) Start() {

	fmt.Println("RabbitTask start ")
	for _, groupId := range taskUtil.GetAllGroupIds() {
		_ = l.svcCtx.MqClient.Subscribe(fmt.Sprintf("austin.biz.%s", groupId), l.onMassage)
	}

	_ = l.svcCtx.MqClient.Subscribe("sms.record", l.onMassage)

	select {}
}

func (l *RabbitTask) Stop() {
	fmt.Println("RabbitTask stop ")
	l.svcCtx.MqClient.Close()
}

func (l *RabbitTask) onMassage(delivery amqp.Delivery) {
	ctx := context.Background()
	var taskList []types.TaskInfo
	_ = jsonx.Unmarshal(delivery.Body, &taskList)
	for _, taskInfo := range taskList {
		channel := channelType.TypeCodeEn[taskInfo.SendChannel]
		msgType := messageType.TypeCodeEn[taskInfo.MsgType]
		err := pending.Submit(ctx, fmt.Sprintf("%s.%s", channel, msgType), pending.NewTask(taskInfo, l.svcCtx))
		if err != nil {
			logx.WithContext(ctx).Errorw("submit task err",
				logx.Field("内容", string(delivery.Body)),
				logx.Field("err", err))
		}
	}
}

func (l *RabbitTask) handlerSmsRecord(delivery amqp.Delivery) {
	ctx := context.Background()
	var msg []model.SmsRecord
	_ = jsonx.Unmarshal(delivery.Body, &msg)
	dbx.GetDb(ctx).Create(msg)
}
