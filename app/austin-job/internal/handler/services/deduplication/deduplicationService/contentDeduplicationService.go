package deduplicationService

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/services/deduplication/structs"
	"austin-go/app/austin-job/internal/svc"
)

type contentDeduplicationService struct {
	svcCtx *svc.ServiceContext
}

func NewContentDeduplicationService(svcCtx *svc.ServiceContext) structs.DuplicationService {
	return &contentDeduplicationService{svcCtx: svcCtx}
}

func (c contentDeduplicationService) Deduplication(param structs.DeduplicationConfigItem, taskInfo *types.TaskInfo) {

}
