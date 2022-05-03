package handlers

import (
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
)

type dingDingWorkNoticeHandler struct {
}

func NewDingDingWorkNoticeHandler() IHandler {
	return dingDingWorkNoticeHandler{}
}
func (h dingDingWorkNoticeHandler) DoHandler(taskInfo types.TaskInfo) (err error) {
	var content content_model.DingDingContentModel
	getContentModel(taskInfo.ContentModel, &content)

	return nil
}
