package redis

import (
	"context"
	"fmt"
)

// 链表
func PushList(ctx context.Context) {
	// lpush or rpush insert some values from left or right
	rdb.LPush(ctx, "list", "1", "2", "3").Result()
	rdb.RPush(ctx, "list", "4", "5", "6").Result()
	// rpoplpush k1 k2 从k1右边吐出一个值 插到k2左边
	rdb.RPopLPush(ctx, "list", "list2").Result()
	// linsert k before v newv 在k的v之前插入新v
	rdb.LInsertBefore(ctx, "list", "1", "p").Result()
	// linsert k fter v newv 在k的v之后插入新v
	rdb.LInsertAfter(ctx, "list", "1", "q").Result()
	// lrem k n v 从边做删除n个value 从做到右
	rdb.LRem(ctx, "list", 1, "0").Result()
	// lset k index v 将列表key下标为index的值替换为value
	rdb.LSet(ctx, "list", 1, "n").Result()
}

func PopList(ctx context.Context) {
	// lrange 按照索引下标获取元素
	value, _ := rdb.LRange(ctx, "list", 0, -1).Result()
	fmt.Println(value)
	// lpop rpop 从左边或者右边吐出一个值 值吐干净了 key也就不存在了
	left, _ := rdb.LPop(ctx, "list").Result()
	fmt.Println(left)
	right, _ := rdb.RPop(ctx, "list").Result()
	fmt.Println(right)
	// lrange 按照索引下标获取元素
	v, _ := rdb.LRange(ctx, "list2", 0, -1).Result()
	fmt.Println(v)
	// lindex k index 根据index获取value
	i, _ := rdb.LIndex(ctx, "list", 1).Result()
	fmt.Println(i)
	// llen 获取key的value长度
	length, _ := rdb.LLen(ctx, "list").Result()
	fmt.Println(length)
}
