package ratelimitutil

import (
	"austin-go/app/austin-common/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/time/rate"
	"time"
)

const RequestRateLimit = 10     //根据真实请求数限流
const SendUserNumRateLimit = 20 //根据发送用户数限流

type RateLimitUtil struct {
	limiter           *rate.Limiter
	rateLimitStrategy int64
}

func NewRateLimitUtil(limiter *rate.Limiter, rateLimitStrategy int64) *RateLimitUtil {
	return &RateLimitUtil{limiter: limiter, rateLimitStrategy: rateLimitStrategy}
}

func (r RateLimitUtil) Limit(ctx context.Context, taskInfo types.TaskInfo) (err error) {
	now := time.Now()
	if r.rateLimitStrategy == RequestRateLimit {
		err = r.limiter.WaitN(ctx, 1)
	}
	if r.rateLimitStrategy == SendUserNumRateLimit {
		err = r.limiter.WaitN(ctx, len(taskInfo.Receiver))
	}
	logx.Info("costTime:" + time.Now().Sub(now).String())
	return nil
}
