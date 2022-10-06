package handlers

import (
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
	"context"
)

type dingDingWorkNoticeHandler struct {
	BaseHandler
}

func NewDingDingWorkNoticeHandler() IHandler {
	return dingDingWorkNoticeHandler{}
}
func (h dingDingWorkNoticeHandler) DoHandler(ctx context.Context, taskInfo types.TaskInfo) (err error) {
	var content content_model.DingDingContentModel
	getContentModel(taskInfo.ContentModel, &content)

	return nil
}
