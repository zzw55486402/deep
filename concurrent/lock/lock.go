package lock

import (
	"fmt"
	"sync"
	"time"
)

/*
	goroutine 1s 输出一句话
	主线程 2s 输出一句话 10次后退出程序
*/

func MainThread() {
	go Routine(10)
	for i := 1; i <= 10; i++ {
		fmt.Println("main routine print: ", i)
		time.Sleep(time.Second * 2)
	}
}

func Routine(times int) {
	for i := 1; i <= times; i++ {
		fmt.Println("goroutine print: ", i)
		time.Sleep(time.Second)
	}
}

/*
	加锁解决并发竞争资源的问题
*/

var (
	lock    sync.Mutex
	testMap = make(map[int]int, 10)
)

func testNum(num int) {
	res := 1
	for i := 1; i <= num; i++ {
		res *= i
	}
	lock.Lock()
	defer lock.Unlock()
	testMap[num] = res
}

func UseLock() {
	start := time.Now()
	for i := 1; i < 20; i++ {
		go testNum(i)
	}
	time.Sleep(time.Second * 5)
	for k, v := range testMap {
		fmt.Println("数字：", k, "值：", v)
	}
	end := time.Since(start)
	fmt.Println("time: ", end)
}
