package channel

import (
	"fmt"
	"sync"
	"time"
)

/*
	Use channel to handle concurrent
*/

func testNum(num int, ch chan [2]int) {
	res := 1
	for i := 1; i <= num; i++ {
		res *= i
	}
	ch <- [2]int{num, res}
}

func UseChannel() {
	ch := make(chan [2]int, 20)
	start := time.Now()
	for i := 1; i < 20; i++ {
		go testNum(i, ch)
	}
	// for k := range ch { // 使用显示的for循环需要close 否则会出错 关闭的管道不能写入
	// 	fmt.Println("数字：", k[0], "值：", k[1])
	// }
	time.Sleep(time.Second * 5)
	close(ch)
forEach:
	for {
		/*
			select 用来遍历channel 如果channel无响应则看有没有default
			如果有default 那么就去走default 如果没有就一直等待 直到某个channel响应
			select是一次性的 如果想监听channel多次需要配合for循环
		*/
		select {
		case val, ok := <-ch:
			if !ok {
				break forEach
			}
			fmt.Println("数字：", val[0], "值：", val[1])
		default:
			fmt.Println("default")
		}
	}
	end := time.Since(start)
	fmt.Println("time: ", end)
}

func isPrime(ch, prime chan int, exitChan chan bool) {
	flag := false
label:
	for {
		select {
		case num, ok := <-ch:
			if !ok {
				break label
			}
			flag = true
			for j := 2; j < num; j++ {
				if num%j == 0 {
					flag = false
					continue
				}
			}
			if flag {
				prime <- num
			}
		default:
			break label
		}
	}
	exitChan <- true
}

var (
	ch = make(chan int, 100)
)

func GetPrime() {
	prime := make(chan int, 100)
	exitChan := make(chan bool, 8)
	go initChan(100)
	for i := 0; i < 8; i++ {
		go isPrime(ch, prime, exitChan)
	}
	go func() {
		for i := 0; i < 7; i++ {
			<-exitChan
		}
	}()
label:
	for {
		select {
		case res := <-prime:
			fmt.Println("素数是：", res)
		default:
			break label
		}
	}
}

func initChan(num int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func PrintOddAndNormal() {
	odd := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(2)
	defer close(odd)
	go func() {
		defer wg.Done()
		for i := 0; i <= 100; i++ {
			if i%2 == 0 {
				fmt.Println("偶数为：", i)
			}
			odd <- true
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i <= 100; i++ {
			<-odd
			if i%2 == 1 {
				fmt.Println("奇数为：", i)
			}
		}
	}()
	wg.Wait()
}

func PrintNumberAndCharacter() {
	num := make(chan bool)
	char := make(chan bool)
	exitChan := make(chan bool)
	go func() {
		number := 1
		for {
			<-num
			fmt.Println(number)
			number++
			fmt.Println(number)
			number++
			char <- true
		}

	}()
	go func() {
		c := 'A'
		for {
			<-char
			if c >= 'Z' {
				exitChan <- true
				return
			}
			fmt.Println(string(c))
			c++
			fmt.Println(string(c))
			c++
			num <- true
		}
	}()
	num <- true
	<-exitChan
	close(exitChan)
	close(char)
	close(num)
}
