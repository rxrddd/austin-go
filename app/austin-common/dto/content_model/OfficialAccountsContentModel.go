package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/taskUtil"
	"austin-go/app/austin-common/types"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type OfficialAccountsContentModel struct {
	Map map[string]string `json:"map"`
	Url string            `json:"url"`
}

func NewOfficialAccountsContentModel() *OfficialAccountsContentModel {
	return &OfficialAccountsContentModel{}
}

func (d OfficialAccountsContentModel) BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{} {
	variables := messageParam.Variables
	var content OfficialAccountsContentModel
	_ = jsonx.Unmarshal([]byte(messageTemplate.MsgContent), &content)
	newVariables := getStringVariables(variables)
	if v, ok := newVariables["url"]; ok && v != "" {
		content.Url = taskUtil.GenerateUrl(v, messageTemplate.ID, messageTemplate.TemplateType)
	}
	if v, ok := variables["map"].(map[string]string); ok {
		content.Map = v
	}
	return d
}
