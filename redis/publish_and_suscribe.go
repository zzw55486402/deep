package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// 发布和订阅
func Subcribe(ctx context.Context) *redis.PubSub {
	// 订阅
	return rdb.Subscribe(ctx, "channel1")
}
func PublishMessage(ctx context.Context) {
	// 发布
	rdb.Publish(ctx, "channel1", "gogogo").Result()
}

func ReceiveMessage(ctx context.Context) {
	// 获取
	client := Subcribe(ctx)
	message, _ := client.ReceiveMessage(ctx)
	fmt.Println(message)
}
