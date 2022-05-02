package redisUtils

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func MGet(ctx context.Context, rds *redis.Redis, keys []string) (result map[string]string, err error) {
	result = make(map[string]string, len(keys))
	val, err := rds.MgetCtx(ctx, keys...)
	if err != nil {
		return result, nil
	}
	for i, key := range keys {
		result[key] = val[i]
	}
	return result, err
}
func PipelineSetEx(ctx context.Context, rds *redis.Redis, keys map[string]string, seconds int64) (err error) {
	return rds.PipelinedCtx(ctx, func(pipeliner redis.Pipeliner) error {
		return pipeliner.MSet(ctx, keys).Err()
	})
}
