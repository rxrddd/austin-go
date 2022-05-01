package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/types"
)

type ImContentModel struct {
}

func NewImContentModel() *ImContentModel {
	return &ImContentModel{}
}

func (d ImContentModel) BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{} {
	return d
}
