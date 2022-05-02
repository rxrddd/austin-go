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
	//c.AddFunc("*/5 * * * * ?", l.nightShieldHandler)
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
			logx.Errorf("nightShieldHandler LlenCtx err : %v", err)
			break
		}
		if length <= 0 {
			break
		}

		pop, err := l.svcCtx.RedisClient.LpopCtx(ctx, services.NightShieldButNextDaySendKey)
		if err != nil {
			logx.Errorf("nightShieldHandler LpopCtx err : %v", err)
			continue
		}
		var taskInfo types.TaskInfo
		err = jsonx.Unmarshal([]byte(pop), &taskInfo)
		if err != nil {
			logx.Errorf("nightShieldHandler jsonx.Unmarshal err : %v", err)
			continue
		}
		//todo:: 有个bug就是消息扔不回去了
		channel := channelType.TypeCodeEn[taskInfo.SendChannel]
		msgType := messageType.TypeCodeEn[taskInfo.MsgType]
		str, _ := jsonx.Marshal(taskInfo)
		err = l.svcCtx.MqClient.Publish(str, taskUtil.GetMqKey(channel, msgType))
		if err != nil {
			logx.Errorf("nightShieldHandler Publish err:%v,taskInfo:%s", err, taskInfo)
		}
	}
}

func (l *cronTask) Stop() {
}
