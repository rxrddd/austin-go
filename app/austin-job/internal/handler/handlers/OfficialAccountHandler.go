package handlers

import (
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
	"fmt"
)

type officialAccountHandler struct {
}

func NewOfficialAccountHandler() IHandler {
	return &officialAccountHandler{}
}

func (h officialAccountHandler) DoHandler(taskInfo types.TaskInfo) (err error) {
	var content content_model.OfficialAccountsContentModel
	getContentModel(taskInfo.ContentModel, &content)
	//拼接消息发送

	//记录发送记录
	fmt.Println("officialAccountHandler")

	return nil
}
