package structs

import (
	"austin-go/app/austin-common/types"
	"context"
)

type DeduplicationConfigItem struct {
	Num  int   `json:"num"`
	Time int64 `json:"time"`
}

type DuplicationService interface {
	Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param DeduplicationConfigItem) error
}
