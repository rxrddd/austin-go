package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/types"
)

type EmailContentModel struct {
}

func NewEmailContentModel() *EmailContentModel {
	return &EmailContentModel{}
}

func (d EmailContentModel) BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{} {
	return d
}
