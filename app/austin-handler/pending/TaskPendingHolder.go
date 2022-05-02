package pending

import (
	"austin-go/app/austin-common/taskUtil"
	"context"
	"github.com/panjf2000/ants/v2"
)

type TaskPendingHolder struct {
	pool map[string]*ants.Pool
}

const defaultPoolLimit = 5000

var defaultTaskPendingHolder = NewTaskPendingHolder()

func NewTaskPendingHolder() *TaskPendingHolder {
	//初始化所有的链接池
	groupIds := taskUtil.GetAllGroupIds()
	pool := make(map[string]*ants.Pool)
	for _, value := range groupIds {
		newPool, _ := ants.NewPool(defaultPoolLimit)
		pool[value] = newPool
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
