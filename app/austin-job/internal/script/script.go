package script

import (
	"context"
)

type SmsScript interface {
	Send(ctx context.Context, sms SmsParams) (err error)
}

type SmsParams struct {
	MessageTemplateId int64    `json:"messageTemplateId"`
	Phones            []string `json:"phones"`
	Content           string   `json:"content"`
	SendAccount       int      `json:"sendAccount"`
}
