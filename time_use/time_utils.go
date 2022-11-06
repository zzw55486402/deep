package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	time包的使用
	1年 365天 day
	1天 24小时 hour
	1小时 60分钟 minute
	1分钟 60秒 second
	1秒 1000毫秒 millisecond
	1毫秒 1000微秒 microsecond us
	1微秒 1000纳秒 nanosecond ns
	1纳秒 1000皮秒 picosecond ps
*/

func main() {
	// 1. 获取当前时间
	now := time.Now()
	fmt.Println(now)
	fmt.Printf("%T\n", now)

	// 2.获取指定时间
	date := time.Date(2008, 7, 15, 16, 30, 0, 0, time.Local)
	fmt.Println(date)

	// 3. 格式化时间 time->string类型的转换 这里的2006.01.02 15:04:05是一个格式 go语言的出生时间
	// 我们需要根据这个格式来定义想输出的时间
	strs := now.Format("2006年1月2日 15:04:05")
	fmt.Println(strs)

	// 4. 字符串类型转换为time类型
	t, err := time.Parse("2006年1月2日 15:04:05", strs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

	// 5. 打印time类型的字符串
	fmt.Println(now.String())

	// 6. 获取年月日
	y, m, d := now.Date()
	fmt.Println(y, m, d)

	// 7. 获取时分秒
	h, min, s := now.Clock()
	fmt.Println(h, min, s)

	// 7.单独获取年月日 时分秒
	now.Year()
	now.YearDay() // 今年过了多少天
	now.Month()
	now.Day()
	now.Hour()
	now.Minute()
	now.Second()
	now.Weekday() // 获取当前的星期

	// 8.时间戳 距离1970年1月1日0时0分0秒的差值 秒 纳秒
	now.Unix()        // 秒的差距
	now.UnixNano()    // 纳秒的差距
	time.Now().Unix() // 当前的时间戳
	// 9.时间间隔
	t5 := now.Add(time.Second) // 增加一秒
	fmt.Println(t5)
	t6 := now.AddDate(1, 0, 0) // 增加年月日
	fmt.Println(t6)
	t7 := now.Sub(t5) // 计算两个时间之间的差值
	fmt.Println(t7)

	// 9.当前程序进入睡眠
	time.Sleep(time.Second * 3)

	// 10.随机数的使用
	rand.Seed(time.Now().Unix())
	randNum := int64(rand.Intn(10) + 1)
	fmt.Println(randNum)
	// 需要转换为time.Duration类型来使用
	time.Sleep(time.Second * time.Duration(randNum))
}
