package deduplicationService

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/services/deduplication/structs"
	"austin-go/app/austin-job/internal/svc"
)

type frequencyDeduplicationService struct {
	svcCtx *svc.ServiceContext
}

func NewFrequencyDeduplicationService(svcCtx *svc.ServiceContext) structs.DuplicationService {
	return &frequencyDeduplicationService{svcCtx: svcCtx}
}

func (c frequencyDeduplicationService) Deduplication(param structs.DeduplicationConfigItem, taskInfo *types.TaskInfo) {

}
