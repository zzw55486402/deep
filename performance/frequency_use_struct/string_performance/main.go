package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

/*
	总结：

	字符串最高效的拼接方式是结合预分配内存方式 Grow 使用 string.Builder
	当使用 + 拼接字符串时，生成新字符串，需要开辟新的空间
	当使用 strings.Builder，bytes.Buffer 或 []byte 的内存是按倍数申请的，在原基础上不断增加
	strings.Builder 比 bytes.Buffer 性能更快，一个重要区别在于 bytes.Buffer 转化为字符串重新申请了一块空间
	存放生成的字符串变量；而 strings.Builder 直接将底层的 []byte 转换成字符串类型返回
*/

func main() {
	// 字符串的五种拼接方式
	// 推荐使用 strings.Builder()
	var builder strings.Builder
	// 直到字符串长度的情况下可以使用Grow 不知道的情况下无需使用Grow
	// builder.Grow(10)
	builder.WriteString("ssss")
	builder.WriteString("xxxx")
	fmt.Println(builder.String())

}

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// 1. + 拼接
func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

// 2. fmt.Sprintf() 拼接
func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}

// 3. strings.Builder() 拼接 （性能最好 推荐使用）
// 可以通过Grow函数来预先声明要拼接的字符串的长度 能够更好的提升性能
func builderConcat(n int, str string) string {
	var builder strings.Builder
	builder.Grow(n)
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

// 4. byte.Buffer() 拼接
func bufferConcat(n int, s string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}

// 5. []byte 拼接
// 已知字符串长度的情况下能够提升性能
func preByteConcat(n int, str string) string {
	buf := make([]byte, 0, n*len(str))
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}
