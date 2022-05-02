package limit

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/services/deduplication/structs"
	"austin-go/app/austin-job/internal/svc"
	"context"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
)

const slideWindowLimitServiceTag = "SW_"

//滑动窗口
type slideWindowLimitService struct {
	svcCtx *svc.ServiceContext
}

func (s slideWindowLimitService) LimitFilter(_ context.Context, duplication structs.DeduplicationService, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) (filterReceiver []string, err error) {
	filterReceiver = make([]string, 0)
	for _, receiver := range taskInfo.Receiver {
		key := slideWindowLimitServiceTag + deduplicationSingleKey(duplication, taskInfo, receiver)
		l := limit.NewPeriodLimit(cast.ToInt(param.Time), param.Num, s.svcCtx.RedisClient, "slideWindowLimitService")
		code, err := l.Take(key)
		if err != nil {
			logx.Errorf("slideWindowLimitService Take err:%v", err)
			continue
		}
		if code != limit.Allowed {
			filterReceiver = append(filterReceiver, receiver)
		}

	}
	return filterReceiver, nil
}

func NewSlideWindowLimitService(svcCtx *svc.ServiceContext) structs.LimitService {
	return &slideWindowLimitService{svcCtx: svcCtx}
}
