package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/types"
)

type MiniProgramContentModel struct {
}

func NewMiniProgramContentModel() *MiniProgramContentModel {
	return &MiniProgramContentModel{}
}

func (m MiniProgramContentModel) BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{} {
	return m
}
