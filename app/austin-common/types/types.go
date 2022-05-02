package types

import "github.com/zeromicro/go-zero/core/jsonx"

type TaskInfo struct {
	MessageTemplateId int64       `json:"messageTemplateId"`
	BusinessId        int64       `json:"businessId"`
	Receiver          []string    `json:"receiver"` //先去重
	IdType            int         `json:"idType"`
	SendChannel       int         `json:"sendChannel"`
	TemplateType      int         `json:"templateType"`
	MsgType           int         `json:"msgType"`
	ShieldType        int         `json:"shieldType"`
	ContentModel      interface{} `json:"contentModel"`
	SendAccount       int         `json:"sendAccount"`
}

func (t TaskInfo) String() string {
	marshal, _ := jsonx.Marshal(t)
	return string(marshal)
}

type ContentModel struct {
}

type SendTaskModel struct {
	MessageTemplateId int64          `json:"messageTemplateId"`
	MessageParamList  []MessageParam `json:"messageParamList"`
	TaskInfo          []TaskInfo     `json:"taskInfo"`
}

type MessageParam struct {
	Receiver  string            `json:"receiver"`           //接收者 多个用,逗号号分隔开
	Variables map[string]string `json:"variables,optional"` //可选 消息内容中的可变部分(占位符替换)
	Extra     map[string]string `json:"extra,optional"`     //可选 扩展参数
}
