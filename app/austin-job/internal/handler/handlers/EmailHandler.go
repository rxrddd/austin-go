package handlers

import (
	"austin-go/app/austin-common/dto/account"
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-support/utils/accountUtils"
	"context"
	"github.com/pkg/errors"
	"gopkg.in/gomail.v2"
)

type emailHandler struct {
}

func NewEmailHandler() IHandler {
	return emailHandler{}
}
func (h emailHandler) DoHandler(ctx context.Context, taskInfo types.TaskInfo) (err error) {
	var content content_model.EmailContentModel
	getContentModel(taskInfo.ContentModel, &content)
	m := gomail.NewMessage()

	var acc account.EmailAccount
	err = accountUtils.GetAccount(ctx, taskInfo.SendAccount, &acc)
	if err != nil {
		return errors.Wrap(err, "emailHandler get account err")
	}

	m.SetHeader("From", m.FormatAddress(acc.Username, "官方"))

	m.SetHeader("To", taskInfo.Receiver...) //主送

	m.SetHeader("Subject", content.Title)
	//发送html格式邮件。
	m.SetBody("text/html", content.Content)

	d := gomail.NewDialer(acc.Host, acc.Port, acc.Username, acc.Password)
	if err := d.DialAndSend(m); err != nil {
		return errors.Wrap(err, "emailHandler DialAndSend err")
	}
	return nil
}
