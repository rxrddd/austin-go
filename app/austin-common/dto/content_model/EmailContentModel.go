package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/taskUtil"
	"austin-go/app/austin-common/types"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type EmailContentModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewEmailContentModel() *EmailContentModel {
	return &EmailContentModel{}
}

func (d EmailContentModel) BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{} {
	variables := messageParam.Variables
	var content EmailContentModel
	_ = jsonx.Unmarshal([]byte(messageTemplate.MsgContent), &content)
	newVariables := getStringVariables(variables)
	content.Content = taskUtil.ReplaceByMap(content.Content, newVariables)
	content.Title = newVariables["title"]
	return content
}
