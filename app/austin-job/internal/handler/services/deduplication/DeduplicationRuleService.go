package deduplication

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/services/deduplication/deduplicationService"
	"austin-go/app/austin-job/internal/handler/services/deduplication/structs"
	"austin-go/app/austin-job/internal/svc"
	"austin-go/repo"
	"context"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
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
	//应该查询数据库的配置进行限流操作
	one, err := repo.NewMessageTemplateRepo(l.svcCtx.Config.CacheRedis).
		One(ctx, taskInfo.MessageTemplateId)
	if err != nil {
		logx.Errorf("deduplicationRuleService 查询模板错误 err:%v", err)
		return
	}

	var deduplicationConfig = make(map[string]structs.DeduplicationConfigItem)
	err = jsonx.Unmarshal([]byte(one.DeduplicationConfig), &deduplicationConfig)
	if err != nil {
		logx.Errorf("deduplicationRuleService jsonx.Unmarshal err:%v", err)
		return
	}
	if len(deduplicationConfig) <= 0 {
		//没配置限流策略
		return
	}

	for key, value := range deduplicationConfig {
		arr := strings.Split(key, "_")
		if len(arr) < 1 {
			continue
		}
		curKey := cast.ToInt(arr[1])
		exec, flag := getExec(curKey, l.svcCtx)
		//表示没匹配到对于的执行器
		if !flag {
			continue
		}
		err := exec.Deduplication(ctx, taskInfo, value)
		if err != nil {
			logx.Error("exec.Deduplication err:%v", err)
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
