package content_model

import (
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/taskUtil"
	"austin-go/app/austin-common/types"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type OfficialAccountsContentModel struct {
	Map        map[string]string `json:"map"`         //消息数据
	Url        string            `json:"url"`         // 消息的URL地址
	TemplateId string            `json:"template_sn"` // 发送消息的模版ID
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
	content.Map = cast.ToStringMapString(variables["map"])
	content.TemplateId = messageTemplate.TemplateSn
	return content
}
