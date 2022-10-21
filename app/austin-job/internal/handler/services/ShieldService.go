package services

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
)

type shieldService struct {
	svcCtx *svc.ServiceContext
}

const (
	NightNoShield                = 10                  //夜间不屏蔽
	NightShield                  = 20                  //夜间屏蔽
	NightShieldButNextDaySend    = 30                  //夜间屏蔽(次日早上9点发送)
	NightShieldButNextDaySendKey = "night_shield_send" //夜间屏蔽redis key
)

//屏蔽服务
func NewShieldService(svcCtx *svc.ServiceContext) *shieldService {
	return &shieldService{
		svcCtx: svcCtx,
	}
}

func (l shieldService) Shield(ctx context.Context, taskInfo *types.TaskInfo) {
	if taskInfo.ShieldType == NightNoShield {
		return
	}

	if isNight() {
		if taskInfo.ShieldType == NightShield {
			//夜间屏蔽
			//发送到mq
			taskInfo.Receiver = []string{} //置空发送人
		}
		if taskInfo.ShieldType == NightShieldButNextDaySend {
			//夜间屏蔽,次日9点发送 扔到redis list里面 定时任务消费
			//发送到mq
			marshal, _ := jsonx.Marshal(taskInfo)
			err := l.svcCtx.RedisClient.PipelinedCtx(ctx, func(pipeliner redis.Pipeliner) (err error) {
				_, err = l.svcCtx.RedisClient.Lpush(NightShieldButNextDaySendKey, marshal)
				if err != nil {
					return err
				}

				expire := int(time.Now().AddDate(0, 0, 1).Unix() - time.Now().Unix())
				err = l.svcCtx.RedisClient.ExpireCtx(ctx, NightShieldButNextDaySendKey, expire)
				if err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				logx.WithContext(ctx).Errorw("夜间屏蔽(次日早上9点发送)模式 写入redis错误", logx.Field("task_info", taskInfo), logx.Field("err", err))
			}
			taskInfo.Receiver = []string{} //置空发送人
		}
	}
}

func isNight() bool {
	//return true
	return time.Now().Hour() < 8
}
