package handlers

import (
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
)

type dingDingRobotHandler struct {
}

func NewDingDingRobotHandler() IHandler {
	return dingDingRobotHandler{}
}
func (h dingDingRobotHandler) DoHandler(taskInfo types.TaskInfo) (err error) {
	var content content_model.DingDingContentModel
	getContentModel(taskInfo.ContentModel, &content)

	return nil
}
