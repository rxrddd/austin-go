package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/task_util"
	"austin-go/app/austin-common/types"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type SmsContentModel struct {
}

func NewSmsContentModel() *SmsContentModel {
	return &SmsContentModel{}
}

/**
messageParam 入参
messageTemplate 模板数据库配置
*/
func (s SmsContentModel) BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{} {
	variables := messageParam.Variables

	var originParams map[string]string
	_ = jsonx.Unmarshal([]byte(messageTemplate.MsgContent), &originParams)

	var newMap = make(map[string]string)
	//todo:: 这里写的有问题 逻辑理不清楚 需要找个java朋友学习一下
	for _, key := range s.GetVariables() {
		if v, ok := originParams[key]; ok {
			newMap[key] = task_util.ReplaceByMap(v, variables)
		}
		if key == "url" {
			newMap[key] = task_util.GenerateUrl(newMap["url"], messageTemplate.ID, messageTemplate.TemplateType)
		}
	}
	return newMap
}
func (s SmsContentModel) GetVariables() []string {
	return []string{
		"content",
		"url",
	}
}
