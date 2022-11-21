package aliyun

import (
	"austin-go/app/austin-common/dto/account"
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/script"
	"austin-go/app/austin-support/utils/accountUtils"
	"context"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/zeromicro/go-zero/core/jsonx"
	"strings"

	smsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
)

type AliyunSms struct {
}

func NewAliyunSms() script.SmsScript {
	return &AliyunSms{}
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
	request := &smsapi.SendSmsRequest{}
	request.SetPhoneNumbers(strings.Join(taskInfo.Receiver, ","))
	request.SetSignName(acc.SignName)
	request.SetTemplateCode(taskInfo.TemplateSn)
	str, _ := jsonx.MarshalToString(taskInfo.MessageParam.Variables)

	request.SetTemplateParam(str)
	response, err := cli.SendSms(request)
	if err != nil {
		return fmt.Errorf("Client.Send() error = %v", err)
	}
	fmt.Println(response.Body.String())
	return nil
}
