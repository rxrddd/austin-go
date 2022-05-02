package listen

import (
	"austin-go/app/austin-job/internal/mqs"
	"austin-go/app/austin-job/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/service"
)

//asynq
func RabbitMqs(ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {

	return []service.Service{
		mqs.NewRabbitTask(ctx, svcCtx),
	}

}
