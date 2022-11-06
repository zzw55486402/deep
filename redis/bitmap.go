package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func SetBitMap(ctx context.Context) {
	// setbit k offset value(0 or 1)
	rdb.SetBit(ctx, "bit", 5, 1).Result()
	rdb.SetBit(ctx, "bit", 6, 1).Result()
	rdb.SetBit(ctx, "bit", 7, 1).Result()
	rdb.SetBit(ctx, "bit2", 0, 1).Result()
	// bittop and/or/not/xor destkey k1 k2 and交集 or并集 not非 xor异或
	rdb.BitOpAnd(ctx, "b1", "bit", "bit2")
	rdb.BitOpOr(ctx, "b2", "bit", "bit2")
	rdb.BitOpNot(ctx, "b3", "bit")
	rdb.BitOpXor(ctx, "b4", "bit", "bit2")

}
func GetBitMap(ctx context.Context) {
	// getbit k offset
	v, _ := rdb.GetBit(ctx, "bit", 5).Result()
	fmt.Println(v)
	v, _ = rdb.GetBit(ctx, "bit", 1).Result()
	fmt.Println(v)
	// bitcount k start end 统计start-end之间为1的个数
	count, _ := rdb.BitCount(ctx, "bit", &redis.BitCount{
		Start: 0,
		End:   10,
	}).Result()
	fmt.Println(count)

	count, _ = rdb.BitCount(ctx, "b1", &redis.BitCount{
		Start: 0,
		End:   10,
	}).Result()
	fmt.Println(count)

	count, _ = rdb.BitCount(ctx, "b2", &redis.BitCount{
		Start: 0,
		End:   10,
	}).Result()
	fmt.Println(count)

	count, _ = rdb.BitCount(ctx, "b3", &redis.BitCount{
		Start: 0,
		End:   10,
	}).Result()
	fmt.Println(count)

	count, _ = rdb.BitCount(ctx, "b4", &redis.BitCount{
		Start: 0,
		End:   10,
	}).Result()
	fmt.Println(count)
}
