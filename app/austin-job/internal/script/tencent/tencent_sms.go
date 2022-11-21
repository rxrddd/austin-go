package tencent

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
	errors2 "github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"github.com/zeromicro/go-zero/core/jsonx"
	"regexp"
	"strings"
	"time"
)

const httpProfileEndpoint = "sms.tencentcloudapi.com"

type TencentSms struct {
	SendSuccess func(msg []byte)
}

func NewTencentSms(sendSuccess func(msg []byte)) script.SmsScript {
	return &TencentSms{
		SendSuccess: sendSuccess,
	}
}

func (t *TencentSms) Send(ctx context.Context, taskInfo types.TaskInfo) (err error) {
	var acc account.TencentSmsAccount
	err = accountUtils.GetAccount(ctx, taskInfo.SendAccount, &acc)
	if err != nil {
		return err
	}

	var content content_model.SmsContentModel
	utils.GetContentModel(taskInfo.ContentModel, &content)

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
	request.TemplateParamSet = common.StringPtrs(templateParamSet(taskInfo, content))

	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return errors2.Wrap(err, "An API error has returned")
	}
	if err != nil {
		return err
	}
	//腾讯云返回结果 根据业务进行处理
	t.smsRecord(response, taskInfo, acc, content)
	return nil
}

func (t *TencentSms) smsRecord(response *sms.SendSmsResponse, taskInfo types.TaskInfo, acc account.TencentSmsAccount, content content_model.SmsContentModel) {
	requestId := *response.Response.RequestId
	insert := make([]model.SmsRecord, 0)
	for _, v := range response.Response.SendStatusSet {
		insert = append(insert, model.SmsRecord{
			ID:                idgen.NextID(),
			MessageTemplateID: taskInfo.MessageTemplateId,
			Phone:             cast.ToInt64(strings.ReplaceAll(*v.PhoneNumber, "+86", "")),
			SupplierID:        cast.ToInt8(acc.SupplierId),
			SupplierName:      acc.SupplierName,
			MsgContent:        content.ReplaceContent,
			SeriesID:          *v.SerialNo,
			ChargingNum:       cast.ToInt8(v.Fee),
			Status:            20,
			SendDate:          cast.ToInt32(time.Now().Format(timex.DateLayout)),
			Created:           cast.ToInt32(time.Now().Unix()),
			RequestId:         requestId,
			SendChannel:       script.TENCENT,
		})
	}

	//处理发现消息记录
	marshal, _ := jsonx.Marshal(insert)
	t.SendSuccess(marshal)
}

func templateParamSet(taskInfo types.TaskInfo, content content_model.SmsContentModel) []string {
	//特殊处理一下顺序,根据模板内容顺序处理,因为map会乱序,腾讯云需要一个字符串数组
	res, _ := regexp.Compile("\\{\\$([a-zA-Z]+)\\}")
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
