package limit

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/services/deduplication/structs"
)

type SlideWindowLimitService struct {
}

func (s SlideWindowLimitService) LimitFilter(duplication structs.DuplicationService, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) {

}
