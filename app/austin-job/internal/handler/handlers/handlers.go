package handlers

import (
	"austin-go/app/austin-common/enums/channelType"
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/svc"
	"context"
	"sync"
)

var (
	once          sync.Once
	handlerHolder map[int]IHandler
)

const flowControlEmail = "flow_control_email"

// SetUp 初始化所有handler
func SetUp(svcCtx *svc.ServiceContext) {
	once.Do(func() {
		handlerHolder = map[int]IHandler{
			channelType.Sms:                NewSmsHandler(),
			channelType.Email:              NewEmailHandler(svcCtx),
			channelType.OfficialAccounts:   NewOfficialAccountHandler(),
			channelType.EnterpriseWeChat:   NewEnterpriseWeChatHandler(),
			channelType.DingDing:           NewDingDingRobotHandler(),
			channelType.DingDingWorkNotice: NewDingDingWorkNoticeHandler(),
		}
	})

}

func GetHandler(sendChannel int) IHandler {
	return handlerHolder[sendChannel]
}

type BaseHandler struct {
}

func (b BaseHandler) Limit(ctx context.Context, taskInfo types.TaskInfo) bool {
	return true
}
