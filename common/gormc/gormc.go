package gormc

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
	"gorm.io/gorm"
	"time"
)

const cacheSafeGapBetweenIndexAndPrimary = time.Second * 5

var (
	// ErrNotFound is an alias of sqlx.ErrNotFound.
	ErrNotFound = gorm.ErrRecordNotFound

	// can't use one SingleFlight per conn, because multiple conns may share the same cache key.
	singleFlights = syncx.NewSingleFlight()
	stats         = cache.NewStat("gormc")
)

type (
	// ExecCtxFn defines the sql exec method.
	ExecCtxFn func(ctx context.Context) (sql.Result, error)
	// QueryCtxFn defines the query method.
	QueryCtxFn func(ctx context.Context, v interface{}) error

	// A CachedConn is a DB connection with cache capability.
	CachedConn struct {
		cache cache.Cache
	}
)

// NewConn returns a CachedConn with a redis cluster cache.
func NewConn(c cache.CacheConf, opts ...cache.Option) CachedConn {
	cc := cache.New(c, singleFlights, stats, gorm.ErrRecordNotFound, opts...)
	return NewConnWithCache(cc)
}

// NewConnWithCache returns a CachedConn with a custom cache.
func NewConnWithCache(c cache.Cache) CachedConn {
	return CachedConn{
		cache: c,
	}
}

// NewNodeConn returns a CachedConn with a redis node cache.
func NewNodeConn(rds *redis.Redis, opts ...cache.Option) CachedConn {
	c := cache.NewNode(rds, singleFlights, stats, gorm.ErrRecordNotFound, opts...)
	return NewConnWithCache(c)
}

// DelCache deletes cache with keys.
func (cc CachedConn) DelCache(keys ...string) error {
	return cc.DelCacheCtx(context.Background(), keys...)
}

// DelCacheCtx deletes cache with keys.
func (cc CachedConn) DelCacheCtx(ctx context.Context, keys ...string) error {
	return cc.cache.DelCtx(ctx, keys...)
}

// GetCache unmarshals cache with given key into v.
func (cc CachedConn) GetCache(key string, v interface{}) error {
	return cc.GetCacheCtx(context.Background(), key, v)
}

// GetCacheCtx unmarshals cache with given key into v.
func (cc CachedConn) GetCacheCtx(ctx context.Context, key string, v interface{}) error {
	return cc.cache.GetCtx(ctx, key, v)
}

// ExecCtx runs given exec on given keys, and returns execution result.
func (cc CachedConn) ExecCtx(ctx context.Context, exec ExecCtxFn, keys ...string) (sql.Result, error) {
	res, err := exec(ctx)
	if err != nil {
		return nil, err
	}

	if err := cc.DelCacheCtx(ctx, keys...); err != nil {
		return nil, err
	}

	return res, nil
}

// QueryRowCtx unmarshals into v with given key and query func.
func (cc CachedConn) QueryRowCtx(ctx context.Context, v interface{}, key string, query QueryCtxFn) error {
	return cc.cache.TakeCtx(ctx, v, key, func(v interface{}) error {
		return query(ctx, v)
	})
}

// SetCache sets v into cache with given key.
func (cc CachedConn) SetCache(key string, val interface{}) error {
	return cc.SetCacheCtx(context.Background(), key, val)
}

// SetCacheCtx sets v into cache with given key.
func (cc CachedConn) SetCacheCtx(ctx context.Context, key string, val interface{}) error {
	return cc.cache.SetCtx(ctx, key, val)
}
