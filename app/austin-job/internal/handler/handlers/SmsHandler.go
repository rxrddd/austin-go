package handlers

import (
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
	"austin-go/common/zutils/dd"
	"fmt"
)

type smsHandler struct {
}

func NewSmsHandler() IHandler {
	return &smsHandler{}
}

func (h smsHandler) DoHandler(taskInfo types.TaskInfo) (err error) {
	var content content_model.SmsContentModel
	getContentModel(taskInfo.ContentModel, &content)
	dd.Print(content)
	//拼接消息发送
	//记录发送记录
	fmt.Println("smsHandler")
	return nil
}
