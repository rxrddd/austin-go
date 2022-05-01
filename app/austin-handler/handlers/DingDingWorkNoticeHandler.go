package handlers

import (
	"austin-go/app/austin-common/types"
	"fmt"
)

type dingDingWorkNoticeHandler struct {
}

func NewDingDingWorkNoticeHandler() IHandler {
	return dingDingWorkNoticeHandler{}
}
func (h dingDingWorkNoticeHandler) DoHandler(taskInfo types.TaskInfo) (err error) {
	fmt.Println(taskInfo)
	return nil
}
