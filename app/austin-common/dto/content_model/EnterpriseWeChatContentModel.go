package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/types"
)

type EnterpriseWeChatContentModel struct {
}

func NewEnterpriseWeChatContentModel() *EnterpriseWeChatContentModel {
	return &EnterpriseWeChatContentModel{}
}

func (d EnterpriseWeChatContentModel) BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{} {
	return d
}
