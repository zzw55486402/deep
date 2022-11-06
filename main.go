package main

import (
	"deep/concurrent/channel"
)

func main() {
	// ctx := context.Background()
	// err := redis.InitClient()
	// if err != nil {
	// 	panic(err)
	// }
	// redis.RedisOperation(ctx)

	// lock.MainThread()
	// lock.UseLock()
	// channel.UseChannel()
	// channel.GetPrime()
	channel.PrintNumberAndCharacter()
}
