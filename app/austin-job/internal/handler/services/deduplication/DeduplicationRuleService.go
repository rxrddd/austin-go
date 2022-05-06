package deduplication

import (
	"austin-go/app/austin-common/repo"
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/services/deduplication/deduplicationService"
	"austin-go/app/austin-job/internal/handler/services/deduplication/structs"
	"austin-go/app/austin-job/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type deduplicationRuleService struct {
	svcCtx *svc.ServiceContext
}

const Content = "10"   //N分钟相同内容去重
const Frequency = "20" //一天内N次相同渠道去重
const deduplicationPrefix = "deduplication_"

func NewDeduplicationRuleService(svcCtx *svc.ServiceContext) *deduplicationRuleService {
	return &deduplicationRuleService{
		svcCtx: svcCtx,
	}
}

func (l deduplicationRuleService) Duplication(ctx context.Context, taskInfo *types.TaskInfo) {
	// 配置样例：{"deduplication_10":{"num":1,"time":300},"deduplication_20":{"num":5}}
	one, err := repo.NewMessageTemplateRepo(l.svcCtx.Config.CacheRedis).
		One(ctx, taskInfo.MessageTemplateId)
	if err != nil {
		logx.Errorw("deduplicationRuleService 查询模板错误 err", logx.Field("err", err))
		return
	}
	if one.DeduplicationConfig == "" {
		//没有配置去重策略直接不管
		return
	}
	var deduplicationConfig = make(map[string]structs.DeduplicationConfigItem)
	err = jsonx.Unmarshal([]byte(one.DeduplicationConfig), &deduplicationConfig)
	if err != nil {
		logx.Errorw("deduplicationRuleService jsonx.Unmarshal err", logx.Field("err", err))
		return
	}
	if len(deduplicationConfig) <= 0 {
		//没配置限流策略
		return
	}

	for key, value := range deduplicationConfig {
		exec, flag := getExec(key, l.svcCtx)
		//表示没匹配到对于的执行器
		if !flag {
			continue
		}
		err := exec.Deduplication(ctx, taskInfo, value)
		if err != nil {
			logx.Errorw("exec.Deduplication err", logx.Field("err", err))
		}
	}

}

func getExec(exec string, svcCtx *svc.ServiceContext) (structs.DuplicationService, bool) {
	var duplicationExec = map[string]structs.DuplicationService{
		deduplicationPrefix + Content:   deduplicationService.NewContentDeduplicationService(svcCtx),
		deduplicationPrefix + Frequency: deduplicationService.NewFrequencyDeduplicationService(svcCtx),
	}
	v, ok := duplicationExec[exec]
	return v, ok
}
