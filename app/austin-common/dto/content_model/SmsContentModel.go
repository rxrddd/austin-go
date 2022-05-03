package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/taskUtil"
	"austin-go/app/austin-common/types"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type SmsContentModel struct {
	Content string `json:"content"`
	Url     string `json:"url"`
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
	var content SmsContentModel
	_ = jsonx.Unmarshal([]byte(messageTemplate.MsgContent), &content)
	newVariables := getStringVariables(variables)
	content.Content = taskUtil.ReplaceByMap(content.Content, newVariables)
	if v, ok := newVariables["url"]; ok && v != "" {
		content.Url = taskUtil.GenerateUrl(v, messageTemplate.ID, messageTemplate.TemplateType)
	}
	return content
}
