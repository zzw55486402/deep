package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// zset 有序集合 每个value都有相应的分数
func SetZSet(ctx context.Context) {
	// zadd k s1 v1 s2 v2 给有序集合k添加元素及其相应score
	rdb.ZAdd(ctx, "zset", &redis.Z{
		Score:  1,
		Member: "11",
	}, &redis.Z{
		Score:  2,
		Member: "22",
	}, &redis.Z{
		Score:  3,
		Member: "33",
	})
	// zincrby k incr v 给元素的score加上增量
	rdb.ZIncrBy(ctx, "zset", 1.1, "11")
	// zrem k v 删除该集合下指定值的元素
	rdb.ZRem(ctx, "zset", "22").Result()
}
func GetZSet(ctx context.Context) {
	// zrange k start stop 返回有序集合k中从start到stop的元素
	elements, _ := rdb.ZRangeWithScores(ctx, "zset", 0, 2).Result()
	fmt.Println(elements)
	// zrangebyscore k min max 返回有序集合中在闭集合[min, max]中的元素 从小到大输出
	ranges, _ := rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{
		Min:    "2",
		Max:    "3",
		Offset: 0,
		Count:  5,
	}).Result()
	fmt.Println(ranges)
	// zrangebyscore k min max 返回有序集合中在闭集合[min, max]中的元素 从大到小输出
	revranges, _ := rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{
		Min:    "2",
		Max:    "3",
		Offset: 0,
		Count:  5,
	}).Result()
	fmt.Println(revranges)
	// zcount k min max 统计该k中从min到max分数的元素个数
	count, _ := rdb.ZCount(ctx, "zset", "0", "5").Result()
	fmt.Println(count)
	// zrank k v 返回该member在集合中的排名
	rank, _ := rdb.ZRank(ctx, "zset", "33").Result()
	fmt.Println(rank)

}
