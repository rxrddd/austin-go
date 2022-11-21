package handlers

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/script"
	"austin-go/app/austin-job/internal/script/aliyun"
	"austin-go/app/austin-job/internal/script/tencent"
	"context"
	"fmt"
	"github.com/pkg/errors"
)

type smsHandler struct {
	BaseHandler
}

func NewSmsHandler() IHandler {
	return &smsHandler{}
}

var sender = map[string]script.SmsScript{
	script.TENCENT: tencent.NewTencentSms(),
	script.ALIYUN:  aliyun.NewAliyunSms(),
}

func (h smsHandler) DoHandler(ctx context.Context, taskInfo types.TaskInfo) (err error) {
	var curSender script.SmsScript
	var ok bool
	if curSender, ok = sender[taskInfo.SmsChannel]; !ok {
		return fmt.Errorf("[%s] 匹配短信发送渠道异常", taskInfo.SmsChannel)
	}
	err = curSender.Send(ctx, taskInfo)
	if err != nil {
		return errors.Wrap(err, "smsHandler send err")
	}
	return nil
}
