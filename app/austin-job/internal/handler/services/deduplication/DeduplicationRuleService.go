package deduplication

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/services/deduplication/deduplicationService"
	"austin-go/app/austin-job/internal/handler/services/deduplication/structs"
	"austin-go/app/austin-job/internal/svc"
	"context"
	"github.com/spf13/cast"
	"strings"
)

type deduplicationRuleService struct {
	svcCtx *svc.ServiceContext
}

const Content = 10   //N分钟相同内容去重
const Frequency = 20 //一天内N次相同渠道去重
const deduplicationPrefix = "deduplication_"

func NewDeduplicationRuleService(svcCtx *svc.ServiceContext) *deduplicationRuleService {
	return &deduplicationRuleService{
		svcCtx: svcCtx,
	}
}

func (l deduplicationRuleService) Duplication(ctx context.Context, taskInfo *types.TaskInfo) {
	// 配置样例：{"deduplication_10":{"num":1,"time":300},"deduplication_20":{"num":5}}

	var deduplicationConfig = map[string]structs.DeduplicationConfigItem{
		"deduplication_10": {
			Num:  1,
			Time: 300,
		},
		"deduplication_20": {
			Num: 5,
		},
	}

	for key, value := range deduplicationConfig {
		arr := strings.Split(key, "_")
		if len(arr) >= 1 {
			curKey := cast.ToInt(arr[1])
			exec, flag := getExec(curKey, l.svcCtx)
			if flag {
				exec.Deduplication(ctx, taskInfo, value)
			}
		}
	}

}

func getExec(exec int, svcCtx *svc.ServiceContext) (structs.DuplicationService, bool) {
	var duplicationExec = map[int]structs.DuplicationService{
		Content:   deduplicationService.NewContentDeduplicationService(svcCtx),
		Frequency: deduplicationService.NewFrequencyDeduplicationService(svcCtx),
	}
	v, ok := duplicationExec[exec]
	return v, ok
}
