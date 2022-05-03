package handlers

import (
	"austin-go/app/austin-common/dto/account"
	"austin-go/app/austin-common/dto/content_model"
	"austin-go/app/austin-common/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type dingDingRobotHandler struct {
}

func NewDingDingRobotHandler() IHandler {
	return dingDingRobotHandler{}
}
func (h dingDingRobotHandler) DoHandler(ctx context.Context, taskInfo types.TaskInfo) (err error) {
	var content content_model.DingDingContentModel
	getContentModel(taskInfo.ContentModel, &content)

	var acc account.DingDingRobotAccount
	err = getAccount(ctx, taskInfo.SendAccount, &acc)
	if err != nil {
		logx.Errorf(" dingDingRobotHandler 解析账号错误  获取账号错误:%s err:%v", taskInfo, err)
	}
	return nil
}
