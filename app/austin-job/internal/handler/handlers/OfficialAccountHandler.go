package handlers

import (
	"austin-go/app/austin-common/dto/account"
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-support/utils/accountUtils"
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/zeromicro/go-zero/core/logx"
)

const colorSep = "|" //以|分割颜色

//公众号订阅消息
type officialAccountHandler struct {
	BaseHandler
}

func NewOfficialAccountHandler() IHandler {
	return &officialAccountHandler{}
}

func (h officialAccountHandler) DoHandler(ctx context.Context, taskInfo types.TaskInfo) (err error) {
	var content content_model.OfficialAccountsContentModel
	getContentModel(taskInfo.ContentModel, &content)
	//拼接消息发送
	var acc account.OfficialAccount

	err = accountUtils.GetAccount(ctx, taskInfo.SendAccount, &acc)
	if err != nil {
		return errors.Wrap(err, "officialAccountHandler get account err")
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

	subscribe := wc.GetOfficialAccount(cfg).GetTemplate()
	templateId := content.TemplateId
	url := content.Url
	params := make(map[string]*message.TemplateDataItem, len(content.Map))

	for key, val := range content.Map {
		color := ""
		value := ""
		arr := strings.Split(val, colorSep)
		if len(arr) == 1 {
			value = arr[0]
		}
		if len(arr) == 2 {
			value = arr[0]
			color = arr[1]
		}
		params[key] = &message.TemplateDataItem{Value: value, Color: color}
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
			logx.Errorw("officialAccountHandler send msg",
				logx.Field("err", err),
				logx.Field("receiver", receiver),
				logx.Field("templateId", templateId))
			continue
		}
		msgIds = append(msgIds, msgID)
	}
	logx.Infow("officialAccountHandler send success", logx.Field("msgIds", msgIds))
	return nil
}
