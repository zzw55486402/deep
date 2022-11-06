package redis

import (
	"context"
	"fmt"
)

// hyperloglog 基数 用来统计uv和不重复数的个数
func SetHyperLogLog(ctx context.Context) {
	// pfadd k e1 e2 添加指定元素到hyperloglog中
	rdb.PFAdd(ctx, "h1", "1", 2, 4, 6, 7, 8).Result()
	rdb.PFAdd(ctx, "h2", "11", 12, 14, 16, 17, 18).Result()
	// pfmerge dest source1 source2 合并计算多个key 可以通过每一天来计算一个星期或者一年
	rdb.PFMerge(ctx, "h3", "h1", "h2")
}
func GetHyperLogLog(ctx context.Context) {
	// pfcount k1 k2 计算出近似基数 可以计算多个近似基数 每个存储每天的UV 然后合并计算即可
	count, _ := rdb.PFCount(ctx, "h").Result()
	fmt.Println(count)
	count, _ = rdb.PFCount(ctx, "h3").Result()
	fmt.Println(count)
}
