package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/types"
)

type DingDingContentModel struct {
	SendType string `json:"sendType"`
	Content  string `json:"content"`
	MediaId  string `json:"mediaId"`
}

func NewDingDingContentModel() *DingDingContentModel {
	return &DingDingContentModel{}
}

func (d DingDingContentModel) BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{} {
	return d
}
