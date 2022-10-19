package limit

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/services/deduplication/structs"
	"austin-go/app/austin-job/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

const slideWindowLimitServiceTag = "SW_"
const slideWindowLimitServicePrefix = "slideWindowLimitServicePrefix_"

//滑动窗口
type slideWindowLimitService struct {
	svcCtx *svc.ServiceContext
}

func (s slideWindowLimitService) LimitFilter(_ context.Context, duplication structs.DeduplicationService, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) (filterReceiver []string, err error) {
	filterReceiver = make([]string, 0)
	for _, receiver := range taskInfo.Receiver {
		key := slideWindowLimitServiceTag + deduplicationSingleKey(duplication, taskInfo, receiver)
		todayEnd, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
		todayEndUnix := todayEnd.AddDate(0, 0, 1).Unix()
		period := todayEndUnix - time.Now().Unix()
		l := limit.NewPeriodLimit(int(period), param.Num, s.svcCtx.RedisClient, slideWindowLimitServicePrefix)
		code, err := l.Take(key)
		if err != nil {
			logx.Errorf("slideWindowLimitService Take err:%v", err)
			continue
		}
		//表示到了上线 直接过滤掉
		if code == limit.OverQuota {
			filterReceiver = append(filterReceiver, receiver)
		}

	}
	return filterReceiver, nil
}

func NewSlideWindowLimitService(svcCtx *svc.ServiceContext) structs.LimitService {
	return &slideWindowLimitService{svcCtx: svcCtx}
}
