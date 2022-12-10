package main

import "fmt"

func main() {
	/*
		首先golang的空结构体不占用内存，因为空结构体不占据内存空间，因此被广泛作为各种场景下的占位符使用。
		一是节省资源，二是空结构体本身就具备很强的语义，即这里不需要任何值，仅作为占位符。
	*/

	// 1.用map实现set，只要key 不要value value的类型设置为空结构体作为占位符 不占用空间
	s := make(Set)
	s.Add("zzw")
	fmt.Println(s)
	// 2. 不发送数据的channel 只用来通知
	ch := make(chan struct{})
	go worker(ch)
	ch <- struct{}{}

	// 3. 在部分场景下，结构体只包含方法，不包含任何的字段。可以用空结构体替代。
}

type Set map[string]struct{}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}

func worker(ch chan struct{}) {
	<-ch
	fmt.Println("go")
}
