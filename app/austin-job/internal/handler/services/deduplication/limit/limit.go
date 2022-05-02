package limit

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/services/deduplication/structs"
)

type LimitService interface {
	LimitFilter(duplication structs.DuplicationService, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem)
}
type DeduplicationService interface {
	Deduplication(param structs.DeduplicationConfigItem)
}

//模板方法
