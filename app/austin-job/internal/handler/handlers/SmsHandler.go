package handlers

import (
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/script"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type smsHandler struct {
}

func NewSmsHandler() IHandler {
	return &smsHandler{}
}

func (h smsHandler) DoHandler(ctx context.Context, taskInfo types.TaskInfo) (err error) {
	err = script.NewTencentSms().Send(ctx, script.SmsParams{
		MessageTemplateId: taskInfo.MessageTemplateId,
		Phones:            taskInfo.Receiver,
		Content:           getContent(taskInfo),
		SendAccount:       taskInfo.SendAccount,
	})
	if err != nil {
		logx.Errorf("smsHandler 发送消息错误:%s err:%v", taskInfo, err)
		return err
	}
	return nil
}

func getContent(taskInfo types.TaskInfo) string {
	var content content_model.SmsContentModel
	getContentModel(taskInfo.ContentModel, &content)
	if content.Url != "" {
		return content.Content + " " + content.Url
	}
	return content.Content
}
