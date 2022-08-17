package handlers

import (
	"austin-go/app/austin-common/consts"
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/script"
	"austin-go/app/austin-job/internal/svc"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type smsHandler struct {
	svcCtx *svc.ServiceContext
}

func NewSmsHandler(svcCtx *svc.ServiceContext) IHandler {
	return &smsHandler{
		svcCtx: svcCtx,
	}
}

func (h smsHandler) DoHandler(ctx context.Context, taskInfo types.TaskInfo) (err error) {
	smsRecord, err := script.NewTencentSms().Send(ctx, script.SmsParams{
		MessageTemplateId: taskInfo.MessageTemplateId,
		Phones:            taskInfo.Receiver,
		Content:           getContent(taskInfo),
		SendAccount:       taskInfo.SendAccount,
	})
	if err != nil {
		return errors.Wrap(err, "smsHandler send err")
	}

	marshal, err := jsonx.Marshal(smsRecord)
	if err != nil {
		return errors.Wrap(err, "smsHandler jsonx.Marshal err")
	}
	err = h.svcCtx.MqClient.Publish(marshal, consts.QueueSmsRecord)
	if err != nil {
		return errors.Wrap(err, "smsHandler Publish err")
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
