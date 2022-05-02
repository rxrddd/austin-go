package redisUtils

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
)

func MGet(ctx context.Context, rds *redis.Redis, keys []string) (result map[string]string, err error) {
	result = make(map[string]string, 0)
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
		err = pipeliner.MSet(ctx, keys).Err()
		if err != nil {
			return err
		}
		for key := range keys {
			err = pipeliner.Expire(ctx, key, time.Duration(seconds)*time.Second).Err()
			if err != nil {
				return err
			}
		}
		return nil
	})
}
