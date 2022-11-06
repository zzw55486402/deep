package main

/*
	package 中引入多个包 包的路径是相对于src文件夹下目录的绝对路径 （gomod）
	包可以嵌套 例如timeutils
	同包下的函数可以直接使用
	只有main函数下的包的才可以用main包
	包可以起别名
	一个目录下的统计文件归属一个包 package的声明要一致
	. 包名 表示可以省略前缀的包名 直接使用该包下的函数
	_ 包名 表示仅引入该包 但是不使用 例如mysql 或者其他包的init的初始化函数等执行
	别名 包名 给一个包起别名
*/

/*
	1. 结构体的大小写也关乎到结构体的私有和公有
*/

/*
	1. 同一个文件夹下的所有文件的包名要相同
	2. 函数名首字母大写代表公共函数，其他的包可以调用，如果是小写则只能在该包内使用
*/

/*
	init函数不能有参数不能有返回值
	引入的该包的时候会先执行init函数再执行其他的函数
	go程序自动的调用
	同一个go文件中可以定义多个init 执行顺序由上到下
	同一个包下不同文件中的init函数会按照文件名从小到大排序 然后顺序的调用各个文件中的init函数
	对于不同包下的函数则是按照引入包的顺序来进行调用init函数
	如果存在依赖-> main-A-B-C 则调用顺序是 C-B-A-main main是最后一个被初始化的 因为他总是依赖于别人
	避免出现循环引用A-B-C-A
	一个包被其它多个包import 但是他只会被初始化一次
*/

/*
	引入外部包 使用 go get 包名 来引入外部包 会存入到GOPATH对应的bin目录中
*/

import (
	"deep/high_level/l_package/people"
	"deep/high_level/l_package/utils"
	now "deep/high_level/l_package/utils/timeutils"
	_ "deep/mock"
	. "fmt"
)

func main() {
	now.GetTime()
	_ = utils.Count(1, 2)
	Println("it is ok")
	p := people.People{}
	Println(p)
}
