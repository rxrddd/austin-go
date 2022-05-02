package handlers

import (
	"austin-go/app/austin-common/types"
)

type dingDingRobotHandler struct {
}

func NewDingDingRobotHandler() IHandler {
	return dingDingRobotHandler{}
}
func (h dingDingRobotHandler) DoHandler(taskInfo types.TaskInfo) (err error) {
	return nil
}
