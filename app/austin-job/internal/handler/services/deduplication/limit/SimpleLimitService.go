package limit

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/services/deduplication/structs"
	"austin-go/app/austin-job/internal/svc"
	"austin-go/app/austin-support/utils/redisUtils"
	"context"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

const simpleLimitServiceTag = "SP_"

//简单去重器（目前承载着 N分钟相同内容去重）
type simpleLimitService struct {
	svcCtx *svc.ServiceContext
}

func NewSimpleLimitService(svcCtx *svc.ServiceContext) structs.LimitService {
	return &simpleLimitService{svcCtx: svcCtx}
}

func (s simpleLimitService) LimitFilter(ctx context.Context, duplication structs.DeduplicationService, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) (filterReceiver []string, err error) {
	filterReceiver = make([]string, 0)
	readyPutRedisReceiver := make(map[string]string, len(taskInfo.Receiver))
	keys := each(deduplicationAllKey(duplication, taskInfo), simpleLimitServiceTag)
	inRedisValue, err := redisUtils.MGet(ctx, s.svcCtx.RedisClient, keys)
	if err != nil {
		logx.Errorw("simpleLimitService inRedisValue MGet err", logx.Field("err", err))
		return filterReceiver, nil
	}
	for _, receiver := range taskInfo.Receiver {
		key := simpleLimitServiceTag + deduplicationSingleKey(duplication, taskInfo, receiver)
		if v, ok := inRedisValue[key]; ok {
			if cast.ToInt(v) > param.Num {
				filterReceiver = append(filterReceiver, receiver)
			} else {
				readyPutRedisReceiver[receiver] = key
			}
		}
	}
	err = s.putInRedis(ctx, readyPutRedisReceiver, inRedisValue, param.Time)
	if err != nil {
		logx.Errorw("simpleLimitService putInRedis err", logx.Field("err", err))
		return filterReceiver, nil
	}
	return filterReceiver, nil
}
func each(keys []string, tag string) []string {
	newRows := make([]string, len(keys))
	for i, key := range keys {
		newRows[i] = tag + key
	}
	return newRows
}

func (s simpleLimitService) putInRedis(ctx context.Context, readyPutRedisReceiver, inRedisValue map[string]string, deduplicationTime int64) error {
	keyValues := make(map[string]string, len(readyPutRedisReceiver))
	for _, value := range readyPutRedisReceiver {
		if val, ok := inRedisValue[value]; ok {
			keyValues[value] = cast.ToString(cast.ToInt(val) + 1)
		} else {
			keyValues[value] = "1"
		}
	}

	if len(keyValues) > 0 {
		return redisUtils.PipelineSetEx(ctx, s.svcCtx.RedisClient, keyValues, deduplicationTime)
	}
	return nil
}
