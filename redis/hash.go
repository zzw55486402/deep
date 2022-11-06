package redis

import (
	"context"
	"fmt"
)

// 哈希表
func SetHash(ctx context.Context) {
	// hset k field value 给哈希表k中插入键值对 field value
	rdb.HSet(ctx, "hash", "k", "v").Result()
	rdb.HSet(ctx, "hash", "num", 1).Result()
	// hmset k f1 v1 f2 v2 给哈希表k中插入多个键值对 f v
	rdb.HMSet(ctx, "hash", "k1", "v1", "k2", "v2").Result()
	// hincrby k f incr/decr 为哈希表中对的f对应的值+1 or -1
	rdb.HIncrBy(ctx, "hash", "num", -1).Result()
	// hsetnx k f v 哈希表k中的f设为v当且仅当f不存在
	rdb.HSetNX(ctx, "hash", "hh", "vv").Result()
	// hdel k f 删除该f
	rdb.HDel(ctx, "hash", "hh").Result()
}
func GetHash(ctx context.Context) {
	// hget k f 从哈希表k中取出f的值
	v, _ := rdb.HGet(ctx, "hash", "k1").Result()
	fmt.Println(v)
	n, _ := rdb.HGet(ctx, "hash", "num").Result()
	fmt.Println(n)
	// hmget k f1, f2 从哈希表k中取出f1, f2 ...s的值
	values, _ := rdb.HMGet(ctx, "hash", "k", "k1", "k2").Result()
	fmt.Println(values...)
	// hexists k f 查看哈希表k中是否存在属性f
	b, _ := rdb.HExists(ctx, "hash", "k1").Result()
	fmt.Println(b)
	// hkeys k 列出k这个哈希表的所有field
	keys, _ := rdb.HKeys(ctx, "hash").Result()
	fmt.Println(keys)
	// hvals k 列出这个哈希表集合的所有value
	vals, _ := rdb.HVals(ctx, "hash").Result()
	fmt.Println(vals)
}
