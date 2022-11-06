package redis

import (
	"context"
	"fmt"
	"time"
)

// 字符串
func StringGet(ctx context.Context) {
	// get
	value, _ := rdb.Get(ctx, "key2").Result()
	fmt.Println(value)
	// strlen s获取value的长度
	strlen, _ := rdb.StrLen(ctx, "key1").Result()
	fmt.Println(strlen)
	// 查看所有的key
	keys, _ := rdb.Keys(ctx, "*").Result()
	fmt.Println(keys)
	number, _ := rdb.Get(ctx, "number").Result()
	fmt.Println(number)
	// 多个get操作
	values, _ := rdb.MGet(ctx, "mkey1", "mkey2").Result()
	fmt.Println(values...)
	// getrange 获取key 一定范围内的value
	r, _ := rdb.GetRange(ctx, "key1", 5, 10).Result()
	fmt.Println(r)
	value, _ = rdb.Get(ctx, "key1").Result()
	fmt.Println(value)
}

func StringSet(ctx context.Context) {
	// k v expire time
	// set
	rdb.Set(ctx, "key1", "value1", -1).Result()
	// append 将给定的值追加到之前value的末尾 key不存在则创建
	rdb.Append(ctx, "key1", "value2").Result()
	// setnx 只有key不存在才设置
	rdb.SetNX(ctx, "key3", "value3", -1).Result()
	// setex 设置相应的过期时间
	rdb.SetEX(ctx, "key4", "value4", 1000*time.Second).Result()
	rdb.Set(ctx, "number", 1, -1).Result()
	// incr 值自增1 非int类型值不行 空值则自增为1
	rdb.Incr(ctx, "number").Result()
	rdb.Incr(ctx, "number").Result()
	rdb.Incr(ctx, "number").Result()
	// decr 值减1 非int类型值不行 空值则为-1
	rdb.Decr(ctx, "number").Result()

	// incrby 按照给定的值进行自增+2
	rdb.IncrBy(ctx, "number", 2).Result()
	// decrby 按照给定的值进行自减-2
	rdb.DecrBy(ctx, "number", 2).Result()

	// mset 多个set操作
	rdb.MSet(ctx, "mkey1", "mvalue1", "mkey2", "mvalue2").Result()

	// msetnx 多个setnx操作 必须是设置的所有key都不存在才行
	rdb.MSetNX(ctx, "mkey3", "mvalue3", "mkey4", "mvalue4").Result()

	// setrange 覆盖从起始位置开始到后面的的value值 value1 -> vvava
	rdb.SetRange(ctx, "key1", 1, "vava").Result()

	// getset 以新换旧 得到旧值
	rdb.GetSet(ctx, "key1", "valuevaluevalue").Result()
}
