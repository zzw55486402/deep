package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// 经纬度的设置 查询 范围查询 距离查询 经纬度哈希
func SetGeoSpatial(ctx context.Context) {
	// geoadd k longitude latitude member ... 增加member的经纬度
	rdb.GeoAdd(ctx, "position", &redis.GeoLocation{
		Name:      "singapore",
		Longitude: 55,
		Latitude:  66,
		Dist:      5,
		GeoHash:   33,
	}, &redis.GeoLocation{
		Name:      "beijing",
		Longitude: 35,
		Latitude:  26,
		Dist:      100,
		GeoHash:   3,
	})
}
func GetGeoSpatial(ctx context.Context) {
	// geopos k member 获取坐标
	res, _ := rdb.GeoPos(ctx, "position", "singapore").Result()
	fmt.Println(res[0].Latitude, res[0].Longitude)

	// geodist k m1 m2 flag(单位) 获取两个位置之间对的直线距离
	dist, _ := rdb.GeoDist(ctx, "position", "singapore", "beijing", "km").Result()
	fmt.Println(dist)

	// georadius k longitude latitudde radis flag 找出在半径radis内的元素
	element, _ := rdb.GeoRadius(ctx, "position", 40, 50, &redis.GeoRadiusQuery{
		Radius: 10000,
		Unit:   "km",
	}).Result()
	fmt.Println(element)
}
