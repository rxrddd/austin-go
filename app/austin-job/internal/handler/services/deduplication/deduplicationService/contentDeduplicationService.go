package deduplicationService

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/services/deduplication/limit"
	"austin-go/app/austin-job/internal/handler/services/deduplication/srv"
	"austin-go/app/austin-job/internal/handler/services/deduplication/structs"
	"austin-go/app/austin-job/internal/svc"
	"context"
)

type contentDeduplicationService struct {
	svcCtx *svc.ServiceContext
}

func NewContentDeduplicationService(svcCtx *svc.ServiceContext) structs.DuplicationService {
	return &contentDeduplicationService{svcCtx: svcCtx}
}

func (c contentDeduplicationService) Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) error {
	return srv.NewContentDeduplicationService(c.svcCtx, limit.NewSimpleLimitService(c.svcCtx)).
		Deduplication(ctx, taskInfo, param)
}
