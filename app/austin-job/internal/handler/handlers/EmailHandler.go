package handlers

import (
	"austin-go/app/austin-common/dto/account"
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/svc"
	"austin-go/app/austin-support/utils/accountUtils"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/limit"
	"gopkg.in/gomail.v2"
)

type emailHandler struct {
	limit *limit.TokenLimiter
}

func NewEmailHandler(svcCtx *svc.ServiceContext) IHandler {
	return emailHandler{
		limit: limit.NewTokenLimiter(3, 10, svcCtx.RedisClient, flowControlEmail),
	}
}

func (h emailHandler) Limit(ctx context.Context, taskInfo types.TaskInfo) bool {
	return h.limit.Allow()
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
