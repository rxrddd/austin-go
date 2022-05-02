package services

import (
	"austin-go/app/austin-common/types"
	"austin-go/common/zutils/arrayUtils"
)

type discardMessageService struct {
}

func NewDiscardMessageService() *discardMessageService {
	return &discardMessageService{}
}

func (l discardMessageService) IsDiscard(taskInfo *types.TaskInfo) bool {
	//根据动态配置的模板id来直接丢弃,使用redis或者配置中心
	var discardMessageTemplateIds = []int64{0}

	if arrayUtils.ArrayInt64In(discardMessageTemplateIds, taskInfo.MessageTemplateId) {
		return true
	}
	return false
}
