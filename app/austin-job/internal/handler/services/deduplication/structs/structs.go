package structs

import "austin-go/app/austin-common/types"

type DeduplicationConfigItem struct {
	Num  int `json:"num"`
	Time int `json:"time"`
}

type DuplicationService interface {
	Deduplication(param DeduplicationConfigItem, taskInfo *types.TaskInfo)
}
