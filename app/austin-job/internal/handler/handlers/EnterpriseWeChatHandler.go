package handlers

import (
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
	"fmt"
)

type enterpriseWeChatHandler struct {
}

func NewEnterpriseWeChatHandler() IHandler {
	return &enterpriseWeChatHandler{}
}

func (h enterpriseWeChatHandler) DoHandler(taskInfo types.TaskInfo) (err error) {
	var content content_model.EnterpriseWeChatContentModel
	getContentModel(taskInfo.ContentModel, &content)
	//拼接消息发送

	//记录发送记录
	fmt.Println(taskInfo)

	return nil
}
