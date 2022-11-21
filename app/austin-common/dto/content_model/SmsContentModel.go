package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/taskUtil"
	"austin-go/app/austin-common/types"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type SmsContentModel struct {
	Content        string `json:"content"` //原始模板 您的验证码是{$code}，{$min}分钟内有效。请勿向他人泄露。如果非本人操作，可忽略本消息。
	ReplaceContent string `json:"content"` //替换后的模板 您的验证码是1011，15分钟内有效。请勿向他人泄露。如果非本人操作，可忽略本消息。
	Url            string `json:"url"`
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
	content.ReplaceContent = taskUtil.ReplaceByMap(content.Content, newVariables)
	if v, ok := newVariables["url"]; ok && v != "" {
		content.Url = taskUtil.GenerateUrl(v, messageTemplate.ID, messageTemplate.TemplateType)
	}
	return content
}
