package limit

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/services/deduplication/structs"
)

//简单去重器（目前承载着 N分钟相同内容去重）
type SimpleLimitService struct {
}

func (s SimpleLimitService) LimitFilter(duplication structs.DuplicationService, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) {

}
