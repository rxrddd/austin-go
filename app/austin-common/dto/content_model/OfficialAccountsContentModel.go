package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/types"
)

type OfficialAccountsContentModel struct {
}

func NewOfficialAccountsContentModel() *OfficialAccountsContentModel {
	return &OfficialAccountsContentModel{}
}

func (d OfficialAccountsContentModel) BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{} {
	return d
}
