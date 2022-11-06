package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

// 数据库的连接
func InitClient() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 地址
		Password: "",               // 密码
		DB:       0,                // 默认DB
		PoolSize: 100,              // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	return err
}

func RedisOperation(ctx context.Context) {
	// StringSet(ctx)
	// StringGet(ctx)

	// PushList(ctx)
	// PopList(ctx)

	// SetSet(ctx)
	// GetSet(ctx)

	// SetHash(ctx)
	// GetHash(ctx)

	// SetZSet(ctx)
	// GetZSet(ctx)

	// PublishMessage(ctx)
	// ReceiveMessage(ctx)

	// SetBitMap(ctx)
	// GetBitMap(ctx)

	// SetHyperLogLog(ctx)
	// GetHyperLogLog(ctx)

	SetGeoSpatial(ctx)
	GetGeoSpatial(ctx)
}
