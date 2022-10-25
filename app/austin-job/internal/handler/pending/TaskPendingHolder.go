package pending

import (
	"austin-go/app/austin-common/taskUtil"
	"context"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"runtime"
)

type TaskPendingHolder struct {
	pool map[string]*ants.Pool
}

var defaultTaskPendingHolder = NewTaskPendingHolder()

func NewTaskPendingHolder() *TaskPendingHolder {
	//初始化所有的链接池
	groupIds := taskUtil.GetAllGroupIds()
	pool := make(map[string]*ants.Pool)
	size := runtime.NumCPU() * 2
	for _, value := range groupIds {
		var pushWorkerPool *ants.Pool
		if wp, err := ants.NewPool(size); err != nil {
			panic(fmt.Errorf("error occurred when creating push worker: %w", err))
		} else {
			pushWorkerPool = wp
		}
		pool[value] = pushWorkerPool
	}
	return &TaskPendingHolder{pool: pool}
}

//把任务提交到对应的池子内
func (t TaskPendingHolder) Submit(ctx context.Context, groupId string, run TaskRun) error {
	return t.pool[groupId].Submit(func() {
		run.Run(ctx)
	})
}

func Submit(ctx context.Context, groupId string, run TaskRun) error {
	return defaultTaskPendingHolder.Submit(ctx, groupId, run)
}
