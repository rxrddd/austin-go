package pending

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-handler/handlers"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

type Task struct {
	TaskInfo types.TaskInfo
}

func (t Task) Run(ctx context.Context) {
	threading.GoSafe(func() {
		// 0. 丢弃消息
		// 1. 屏蔽消息
		// 2.平台通用去重
		// 3. 真正发送消息
		err := handlers.GetHandler(t.TaskInfo.SendChannel).DoHandler(t.TaskInfo)
		if err != nil {
			logx.Error(err)
		}
	})
}

type TaskRun interface {
	Run(ctx context.Context)
}
