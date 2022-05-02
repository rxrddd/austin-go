package listen

import (
	"austin-go/app/austin-job/internal/cron"
	"austin-go/app/austin-job/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/service"
)

//返回所有消费者
func Mqs(svcCtx *svc.ServiceContext) []service.Service {

	ctx := context.Background()

	var services []service.Service
	services = append(services, RabbitMqs(ctx, svcCtx)...)
	services = append(services, []service.Service{
		cron.NewCronTask(ctx, svcCtx),
	}...)

	return services
}
