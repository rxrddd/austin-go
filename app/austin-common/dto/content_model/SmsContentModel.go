package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/taskUtil"
	"austin-go/app/austin-common/types"
	"github.com/spf13/cast"
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
	jsonx.Unmarshal([]byte(messageTemplate.MsgContent), &content)

	var newVariables = make(map[string]string, len(variables))
	for key, variable := range variables {
		if v, ok := variable.(string); ok {
			newVariables[key] = v
		}
	}

	//首先需要把前端数据 json转化到content model 然后处理url内容
	content.Content = taskUtil.ReplaceByMap(content.Content, newVariables)
	if v, ok := variables["url"]; ok && v != "" {
		content.Url = taskUtil.GenerateUrl(cast.ToString(v), messageTemplate.ID, messageTemplate.TemplateType)
	}
	return content
}
