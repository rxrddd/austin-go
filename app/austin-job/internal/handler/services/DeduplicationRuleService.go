package services

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/svc"
)

type deduplicationRuleService struct {
	svcCtx *svc.ServiceContext
}

func NewDeduplicationRuleService(svcCtx *svc.ServiceContext) *deduplicationRuleService {
	return &deduplicationRuleService{
		svcCtx: svcCtx,
	}
}

func (l deduplicationRuleService) Duplication(taskInfo *types.TaskInfo) {

}
