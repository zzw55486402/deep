package main

import (
	"fmt"
	"strconv"
)

/*
	golang泛型
*/
func main() {
	printSlice[string]([]string{"k", "v", "g", "o"})
	printSlice([]string{"k", "v", "g", "o"}) // 函数后的[string]也可以不用加 或推断的
	printSlice[int]([]int{1, 2, 3, 4})
	printSlice[float64]([]float64{1.0, 2.22, 3.34234, 4.42342, 5.42342, 6.42432})

	// 声明一个泛型切片
	type sliceAny[T any] []T
	v := sliceAny[int]{1, 2, 3, 4}
	printSlice[int](v)
	v2 := sliceAny[string]{"1sadas", "2fasfas", "3fasfa", "4fafsaf"}
	printSlice[string](v2)

	// 声明一个泛型map
	type mapAny[K string, V any] map[K]V // map底层不支持K为any所以只能为string
	m1 := mapAny[string, int]{"k": 1}
	m1["s"] = 2
	fmt.Println(m1)

	m2 := mapAny[string, string]{"k": "222ss"}
	m2["sss"] = "dasdasa"
	fmt.Println(m2)
	fmt.Println(m2["sss"])

	// 定义泛型channel
	type C[T any] chan T
	chan1 := make(C[string], 3)
	chan1 <- "string"
	chan1 <- "string2"
	chan1 <- "string3"
	fmt.Println(<-chan1)
	fmt.Println(<-chan1)
	fmt.Println(<-chan1)

	chan2 := make(C[int], 3)
	chan2 <- 1
	chan2 <- 8
	fmt.Println(<-chan2)
	fmt.Println(<-chan2)

	// 泛型约束
	a := add("1", "2xxxx")
	b := add(1, 2)
	// c := add(1, "2") 这样不可取 编译不通过
	fmt.Println(a, b)

	// 通过interface来约束泛型的类型
	ret := ShowPriceList([]Price{1, 2, 3, 4})
	fmt.Println(ret)

	ret2 := ShowPriceList2([]Price1{1, 2, 3, 4})
	// ret4 := ShowPriceList2([]Price1{1.52532, 2, 3, 4}) // 只能是int类型 不能是float会有问题
	ret3 := ShowPriceList2([]Price2{"1xxx", "2ssss"})
	fmt.Println(ret2, ret3)

	// comparable关键字的泛型
	// 支持整数 字符串和自定义的结构体 也可以嵌套在自定义约束中
	fmt.Println(findFunc([]string{"xx", "ss"}, "ss"))
	fmt.Println(findFunc([]int{1, 2}, 1))
}

// 声明泛型函数
func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// 使用类型来约束泛型类型 约束可以用的泛型类型 这样可以提前报错
// 约束NumStr类型为Num或者Str
type NumStr interface {
	Num | Str
}

// 约束Num的类型为整数 浮点数 complex
type Num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~complex64 | ~complex128
}

// 约束Str为string
type Str interface {
	string
}

func add[T NumStr](a, b T) T {
	return a + b
}

// 使用约束方法来约束类型
type Price int

type ShowPrice interface {
	String() string // 通过String来约束传入的值能够调用的方法
}

func (i Price) String() string {
	return strconv.Itoa(int(i))
}

func ShowPriceList[T ShowPrice](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

// 使用约束方法和类型来约束类型
type Price1 int
type Price2 string

type ShowPriceTwo interface {
	String() string // 通过String来约束传入的值能够调用的方法
	~int | ~string  // ~表明这是一个底层类型
}

func (i Price1) String() string {
	return strconv.Itoa(int(i))
}

func (i Price2) String() string {
	return string(i)
}

func ShowPriceList2[T ShowPriceTwo](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

// comparable 泛型关键字

func findFunc[T comparable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			return i
		}
	}
	return -1
}
