package handlers

import (
	"austin-go/app/austin-common/types"
	"fmt"
)

type officialAccountHandler struct {
}

func NewOfficialAccountHandler() IHandler {
	return &officialAccountHandler{}
}

func (h officialAccountHandler) DoHandler(taskInfo types.TaskInfo) (err error) {

	//拼接消息发送

	//记录发送记录
	fmt.Println("officialAccountHandler")

	return nil
}
