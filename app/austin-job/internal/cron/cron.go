package cron

import (
	"austin-go/app/austin-common/enums/channelType"
	"austin-go/app/austin-common/enums/messageType"
	"austin-go/app/austin-common/taskUtil"
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/services"
	"austin-go/app/austin-job/internal/svc"
	"context"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type cronTask struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronTask(ctx context.Context, svcCtx *svc.ServiceContext) *cronTask {
	return &cronTask{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *cronTask) Start() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0 8 * * ?", l.nightShieldHandler)
	c.Start()
	defer c.Stop()
	select {}
}

//夜间屏蔽凌晨8点开始发送
func (l *cronTask) nightShieldHandler() {
	ctx := context.Background()
	for {
		length, err := l.svcCtx.RedisClient.LlenCtx(ctx, services.NightShieldButNextDaySendKey)
		if err != nil {
			logx.Errorw("nightShieldHandler LlenCtx", logx.Field("err", err))
			break
		}
		if length <= 0 {
			break
		}

		pop, err := l.svcCtx.RedisClient.LpopCtx(ctx, services.NightShieldButNextDaySendKey)
		if err != nil {
			logx.Errorw("nightShieldHandler LpopCtx", logx.Field("err", err))
			continue
		}
		var taskInfo types.TaskInfo
		err = jsonx.Unmarshal([]byte(pop), &taskInfo)
		if err != nil {
			logx.Errorw("nightShieldHandler jsonx.Unmarshal", logx.Field("err", err))
			continue
		}
		channel := channelType.TypeCodeEn[taskInfo.SendChannel]
		msgType := messageType.TypeCodeEn[taskInfo.MsgType]
		str, _ := jsonx.Marshal([]types.TaskInfo{taskInfo})
		err = l.svcCtx.MqClient.Publish(str, taskUtil.GetMqKey(channel, msgType))
		if err != nil {
			logx.Errorw("nightShieldHandler Publish",
				logx.Field("taskInfo", taskInfo),
				logx.Field("err", err))
		}

	}
}

func (l *cronTask) Stop() {
}
