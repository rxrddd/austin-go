package handlers

import (
	"austin-go/app/austin-common/dto/account"
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
	"context"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

//公众号订阅消息
type officialAccountHandler struct {
}

func NewOfficialAccountHandler() IHandler {
	return &officialAccountHandler{}
}

func (h officialAccountHandler) DoHandler(ctx context.Context, taskInfo types.TaskInfo) (err error) {
	var content content_model.OfficialAccountsContentModel
	getContentModel(taskInfo.ContentModel, &content)
	//拼接消息发送
	var acc account.OfficialAccount

	err = getAccount(ctx, taskInfo.SendAccount, &acc)
	if err != nil {
		logx.Errorf(" officialAccountHandler 解析账号错误  获取账号错误:%s err:%v", taskInfo, err)
	}
	wc := wechat.NewWechat()
	cacheImpl := cache.NewMemory()

	cfg := &offConfig.Config{
		AppID:          acc.AppID,
		AppSecret:      acc.AppSecret,
		Token:          acc.Token,
		EncodingAESKey: acc.EncodingAESKey,
		Cache:          cacheImpl,
	}

	messageTemplateId := taskInfo.MessageTemplateId
	subscribe := wc.GetOfficialAccount(cfg).GetTemplate()
	templateId := h.getRealWxMpTemplateId(messageTemplateId)
	url := content.Url
	params := make(map[string]*message.TemplateDataItem, len(content.Map))

	for key, val := range content.Map {
		params[key] = &message.TemplateDataItem{Value: val}
	}
	var msgIds []int64
	//如果需要实现跳转小程序 需要在getRealWxMpTemplateId里面返回对应的数据进行操作
	for _, receiver := range taskInfo.Receiver {
		msgID, err := subscribe.Send(&message.TemplateMessage{
			ToUser:     receiver,
			TemplateID: templateId,
			URL:        url,
			Data:       params,
		})
		if err != nil {
			logx.Errorf("officialAccountHandler send msg err:%v receiver:%s templateId:%s", err, receiver, templateId)
			continue
		}
		msgIds = append(msgIds, msgID)
	}
	logx.Info("officialAccountHandler send success msgIds:%v", msgIds)
	return nil
}
func (h officialAccountHandler) getRealWxMpTemplateId(messageTemplateId int64) string {
	//查询数据库得到真实的模板id
	return cast.ToString(messageTemplateId)
}
