package tencent

import (
	"austin-go/app/austin-common/dto/account"
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/script"
	"austin-go/app/austin-support/utils"
	"austin-go/app/austin-support/utils/accountUtils"
	"context"
	"fmt"
	errors2 "github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"regexp"
)

const httpProfileEndpoint = "sms.tencentcloudapi.com"

type TencentSms struct {
}

func NewTencentSms() script.SmsScript {
	return &TencentSms{}
}

func (t *TencentSms) Send(ctx context.Context, taskInfo types.TaskInfo) (err error) {
	var acc account.TencentSmsAccount
	err = accountUtils.GetAccount(ctx, taskInfo.SendAccount, &acc)
	if err != nil {
		return err
	}

	credential := common.NewCredential(
		acc.SecretId,
		acc.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = httpProfileEndpoint
	client, _ := sms.NewClient(credential, "", cpf)

	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs(taskInfo.Receiver)
	request.SmsSdkAppId = common.StringPtr(acc.SmsSdkAppId)
	request.SignName = common.StringPtr(acc.SignName)
	request.TemplateId = common.StringPtr(acc.TemplateId)
	request.TemplateParamSet = common.StringPtrs(templateParamSet(taskInfo))

	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return errors2.Wrap(err, "An API error has returned")
	}
	if err != nil {
		return err
	}
	//腾讯云返回结果 根据业务进行处理
	fmt.Printf("%s", response.ToJsonString())

	return nil
}

func templateParamSet(taskInfo types.TaskInfo) []string {
	//特殊处理一下顺序,根据模板内容顺序处理,因为map会乱序,腾讯云需要一个字符串数组
	res, _ := regexp.Compile("\\{\\$([a-zA-Z]+)\\}")

	var content content_model.SmsContentModel
	utils.GetContentModel(taskInfo.ContentModel, &content)

	arr := res.FindAllString(content.Content, -1)
	newMap := make(map[string]string, len(taskInfo.MessageParam.Variables))
	for k, v := range taskInfo.MessageParam.Variables {
		newMap["{$"+k+"}"] = cast.ToString(v)
	}
	templateParamSet := make([]string, 0)
	for _, s := range arr {
		templateParamSet = append(templateParamSet, newMap[s])
	}
	return templateParamSet
}
