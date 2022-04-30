// Code generated by goctl. DO NOT EDIT.
package types

type SendRequest struct {
	Code              string       `json:"code"`
	MessageTemplateId int64        `json:"messageTemplateId"`
	MessageParam      MessageParam `json:"messageParam"`
}

type MessageParam struct {
	Receiver  string            `json:"receiver"`           //接收者 多个用,逗号号分隔开
	Variables map[string]string `json:"variables,optional"` //可选 消息内容中的可变部分(占位符替换)
	Extra     map[string]string `json:"extra,optional"`     //可选 扩展参数
}

type Response struct {
	Message string `json:"message"`
}
