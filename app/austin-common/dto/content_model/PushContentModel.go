package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/types"
)

type PushContentModel struct {
}

func NewPushContentModel() *PushContentModel {
	return &PushContentModel{}
}

func (d PushContentModel) BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{} {
	return d
}
