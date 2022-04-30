package listen

import (
	"austin-go/app/austin-consumer/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/service"
)

//返回所有消费者
func Mqs(svcCtx *svc.ServiceContext) []service.Service {

	ctx := context.Background()

	var services []service.Service
	services = append(services, RabbitMqs(ctx, svcCtx)...)

	return services
}
