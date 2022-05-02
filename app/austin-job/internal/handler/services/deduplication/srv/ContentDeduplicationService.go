package srv

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/services/deduplication/structs"
	"austin-go/app/austin-job/internal/svc"
	"austin-go/common/zutils/encrypt"
	"context"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type contentDeduplicationService struct {
	svcCtx *svc.ServiceContext
	limit  structs.LimitService

	deduplicationService
}

func NewContentDeduplicationService(svcCtx *svc.ServiceContext, limit structs.LimitService) structs.DeduplicationService {
	return &contentDeduplicationService{svcCtx: svcCtx, limit: limit}
}

func (c contentDeduplicationService) Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) error {
	return c.deduplicationService.Deduplication(ctx, c.limit, c, taskInfo, param)
}

func (c contentDeduplicationService) DeduplicationSingleKey(taskInfo *types.TaskInfo, receiver string) string {
	str, _ := jsonx.Marshal(taskInfo.ContentModel)
	return encrypt.MD5(cast.ToString(taskInfo.MessageTemplateId) + receiver + string(str))
}
