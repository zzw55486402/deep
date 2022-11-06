package redis

import (
	"context"
	"fmt"
)

// 集合 集合中的v各不相同
func SetSet(ctx context.Context) {
	// sadd k v1 v2 在集合k中加入多个v
	rdb.SAdd(ctx, "s", 1, 2, 3, 4, 5).Result()
	// srem k v1 v2 删除集合中的元素
	rdb.SRem(ctx, "s", 1, 2).Result()
	// smove source des v 把集合中的一个值移动到另一个集合
	rdb.SMove(ctx, "s", "e", 4).Result()
}
func GetSet(ctx context.Context) {
	// sinter k1 k2 返回两个集合的交集
	inter, _ := rdb.SInter(ctx, "s", "e").Result()
	fmt.Println(inter)
	// sunion k1 k2 返回两个集合的并集
	union, _ := rdb.SUnion(ctx, "s", "e").Result()
	fmt.Println(union)
	// sdiff k1 k2 返回两个集合的差集
	diff, _ := rdb.SDiff(ctx, "s", "e").Result()
	fmt.Println(diff)
	// smembers k 取出该集合的所有值
	m, _ := rdb.SMembers(ctx, "s").Result()
	fmt.Println(m)
	// smembers k 取出该集合的所有值
	e, _ := rdb.SMembers(ctx, "e").Result()
	fmt.Println(e)
	// sismember k v 判断该v是否在集合k中
	b, _ := rdb.SIsMember(ctx, "s", 2).Result()
	fmt.Println(b)
	// scard k 返回该集合的元素个数
	length, _ := rdb.SCard(ctx, "s").Result()
	fmt.Println(length)
	// spop k 随机的从集合k中吐出一个值
	p, _ := rdb.SPop(ctx, "s").Result()
	fmt.Println(p)
	// srandmember k n 随机的从集合中取出n个值 不够就全部取出来
	n, _ := rdb.SRandMemberN(ctx, "s", 2).Result()
	fmt.Println(n)
}
