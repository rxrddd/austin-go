package handlers

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/script"
	"austin-go/app/austin-job/internal/script/aliyun"
	"austin-go/app/austin-job/internal/script/tencent"
	"austin-go/app/austin-job/internal/svc"
	"context"
	"fmt"
	"github.com/pkg/errors"
)

type smsHandler struct {
	svcCtx *svc.ServiceContext
	BaseHandler
}

func NewSmsHandler(svcCtx *svc.ServiceContext) IHandler {
	return &smsHandler{
		svcCtx: svcCtx,
	}
}

func (h smsHandler) DoHandler(ctx context.Context, taskInfo types.TaskInfo) (err error) {
	sender := map[string]script.SmsScript{
		script.TENCENT: tencent.NewTencentSms(h.SendSuccess),
		script.ALIYUN:  aliyun.NewAliyunSms(h.SendSuccess),
	}
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
func (h smsHandler) SendSuccess(msg []byte) {
	_ = h.svcCtx.MqClient.Publish(msg, "sms-record")
}
