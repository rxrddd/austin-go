package script

import (
	"austin-go/app/austin-common/dto/account"
	"austin-go/app/austin-common/model"
	"austin-go/app/austin-support/utils/accountUtils"
	"austin-go/common/idgen"
	"context"
	"fmt"
	errors2 "github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"time"
)

type TencentSms struct {
}

func NewTencentSms() SmsScript {
	return &TencentSms{}
}

const httpProfileEndpoint = "sms.tencentcloudapi.com"

//腾讯云短信发送实现  未测试
func (t TencentSms) Send(ctx context.Context, smsParams SmsParams) (smsRecord []*model.SmsRecord, err error) {
	var acc account.TencentSmsAccount
	err = accountUtils.GetAccount(ctx, smsParams.SendAccount, &acc)
	smsRecord = make([]*model.SmsRecord, 0)
	if err != nil {
		return smsRecord, err
	}

	credential := common.NewCredential(
		acc.SecretId,
		acc.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = httpProfileEndpoint
	client, _ := sms.NewClient(credential, "", cpf)

	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs(smsParams.Phones)
	request.SmsSdkAppId = common.StringPtr(acc.SmsSdkAppId)
	request.SignName = common.StringPtr(acc.SignName)
	request.TemplateId = common.StringPtr(acc.TemplateId)
	request.TemplateParamSet = common.StringPtrs([]string{smsParams.Content})

	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return smsRecord, errors2.Wrap(err, "An API error has returned")
	}
	if err != nil {
		return smsRecord, err
	}
	//腾讯云返回结果 根据业务进行处理
	fmt.Printf("%s", response.ToJsonString())

	for _, status := range response.Response.SendStatusSet {
		smsRecord = append(smsRecord, &model.SmsRecord{
			ID:                idgen.NextID(),
			MessageTemplateID: smsParams.MessageTemplateId,
			Phone:             string([]byte(*status.PhoneNumber)[3:]),
			SupplierID:        acc.SupplierId,
			SupplierName:      acc.SupplierName,
			MsgContent:        smsParams.Content,
			SeriesID:          acc.SecretId,
			ChargingNum:       int(*status.Fee),
			ReportContent:     *status.Code,
			Status:            10,
			SendDate:          time.Now().Unix(),
			Created:           time.Now().Unix(),
			Updated:           time.Now().Unix(),
		})
	}

	return smsRecord, nil
}
