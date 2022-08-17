package script

import (
	"austin-go/app/austin-common/model"
	"context"
)

type SmsScript interface {
	Send(ctx context.Context, sms SmsParams) (smsRecord []*model.SmsRecord, err error)
}

type SmsParams struct {
	MessageTemplateId int64    `json:"messageTemplateId"`
	Phones            []string `json:"phones"`
	Content           string   `json:"content"`
	SendAccount       int      `json:"sendAccount"`
}
