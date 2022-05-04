package handlers

import (
	"austin-go/app/austin-common/dto/account"
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
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
	err = getAccount(ctx, taskInfo.SendAccount, &acc)
	if err != nil {
		logx.Errorf(" emailHandler 解析账号错误  获取账号错误:%s err:%v", taskInfo, err)
		return
	}

	m.SetHeader("From", m.FormatAddress(acc.Username, "官方"))

	m.SetHeader("To", taskInfo.Receiver...) //主送

	m.SetHeader("Subject", content.Title)
	//发送html格式邮件。
	m.SetBody("text/html", content.Content)

	d := gomail.NewDialer(acc.Host, acc.Port, acc.Username, acc.Password)
	if err := d.DialAndSend(m); err != nil {
		logx.Errorf(" emailHandler DialAndSend  taskInfo:%s err:%v", taskInfo, err)
	}
	return nil
}
