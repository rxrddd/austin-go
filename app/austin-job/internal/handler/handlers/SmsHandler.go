package handlers

import (
	"austin-go/app/austin-common/types"
	"fmt"
)

type smsHandler struct {
}

func NewSmsHandler() IHandler {
	return &smsHandler{}
}

func (h smsHandler) DoHandler(taskInfo types.TaskInfo) (err error) {

	//拼接消息发送
	//记录发送记录
	fmt.Println("smsHandler")
	return nil
}
