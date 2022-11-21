package aliyun

import (
	"austin-go/app/austin-common/dto/account"
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/script"
	"austin-go/app/austin-support/utils"
	"austin-go/app/austin-support/utils/accountUtils"
	"austin-go/common/idgen"
	"austin-go/common/zutils/timex"
	"context"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	smsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/jsonx"
	"time"
)

type AliyunSms struct {
	SendSuccess func(msg []byte)
}

func NewAliyunSms(sendSuccess func(msg []byte)) script.SmsScript {
	return &AliyunSms{
		SendSuccess: sendSuccess,
	}
}
func (t *AliyunSms) Send(ctx context.Context, taskInfo types.TaskInfo) (err error) {
	var acc account.AliyunSmsAccount
	err = accountUtils.GetAccount(ctx, taskInfo.SendAccount, &acc)
	if err != nil {
		return err
	}

	config := &openapi.Config{
		AccessKeyId:     &acc.AccessKeyId,
		AccessKeySecret: &acc.AccessSecret,
	}
	// 访问的域名
	config.Endpoint = tea.String(acc.GatewayURL)
	cli, err := smsapi.NewClient(config)
	if err != nil {
		return fmt.Errorf("smsapi.NewClient error = %v", err)
	}
	var content content_model.SmsContentModel
	utils.GetContentModel(taskInfo.ContentModel, &content)

	for _, v := range taskInfo.Receiver {
		request := &smsapi.SendSmsRequest{}
		request.SetPhoneNumbers(v)
		request.SetSignName(acc.SignName)
		request.SetTemplateCode(taskInfo.TemplateSn)
		str, _ := jsonx.MarshalToString(taskInfo.MessageParam.Variables)

		request.SetTemplateParam(str)
		response, err := cli.SendSms(request)
		if err != nil {
			return fmt.Errorf("Client.Send() error = %v", err)
		}
		if *response.Body.Code == "OK" {
			t.smsRecord(response, taskInfo.MessageTemplateId, v, content)
		}
	}
	return nil
}
func (t *AliyunSms) smsRecord(response *smsapi.SendSmsResponse, messageTemplateId int64, phoneNumber string, content content_model.SmsContentModel) {
	requestId := *response.Body.RequestId
	var insert = model.SmsRecord{
		ID:                idgen.NextID(),
		MessageTemplateID: messageTemplateId,
		Phone:             cast.ToInt64(phoneNumber),
		MsgContent:        content.ReplaceContent,
		Status:            10,
		SendDate:          cast.ToInt32(time.Now().Format(timex.DateLayout)),
		Created:           cast.ToInt32(time.Now().Unix()),
		RequestId:         requestId,
		BizId:             *response.Body.BizId,
		SendChannel:       script.ALIYUN,
	}
	marshal, _ := jsonx.Marshal([]model.SmsRecord{insert})
	t.SendSuccess(marshal)
}
