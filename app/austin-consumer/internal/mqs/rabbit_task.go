package mqs

import (
	"austin-go/app/austin-consumer/internal/svc"
	"austin-go/common/mq"
	"context"
	"fmt"
	"github.com/streadway/amqp"
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

	l.client.Subscribe("test-queue", func(delivery amqp.Delivery) {
		fmt.Println(string(delivery.Body))
	})
	//go func() {
	//	for {
	//		time.Sleep(time.Second)
	//		client.Publish([]byte("test"), "test-queue")
	//	}
	//}()
	select {}
}

func (l *RabbitTask) Stop() {
	l.client.Close()
}
