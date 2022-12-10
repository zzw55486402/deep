package unit_test

import (
	"fmt"
	"testing"
)

// go test -v -run TestXXX 打印出详细的信息
// t.Helper() 打印出详细的调用失败者

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "one_true",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "two_false",
			args: args{
				a: 2,
				b: 2,
			},
			want: 4,
		},
		{
			name: "three_true",
			args: args{
				a: 3,
				b: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		// t.Run(测试名, 测试函数) 子测试
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "one_true",
			args: args{
				a: 1,
				b: 2,
			},
			want: 2,
		},
		{
			name: "two_false",
			args: args{
				a: 2,
				b: 2,
			},
			want: 3,
		},
		{
			name: "three_true",
			args: args{
				a: 3,
				b: 2,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 基准测试 BenchmarkXXX XXX 为函数名
/*
	函数名必须以 Benchmark 开头，后面一般跟待测试的函数名
	参数为 b *testing.B。
	执行基准测试时，需要添加 -bench 参数。
*/

/*
	基准测试中各个字段的含义
	type BenchmarkResult struct {
		N         int           // 迭代次数
		T         time.Duration // 基准测试花费的时间
		Bytes     int64         // 一次迭代处理的字节数
		MemAllocs uint64        // 总的分配内存的次数
		MemBytes  uint64        // 总的分配内存的字节数
	}
*/

// go test -benchmem -bench .
func BenchmarkAdd(b *testing.B) {

	// 如果在运行前基准测试需要一些耗时的配置，则可以使用 b.ResetTimer() 先重置定时器
	b.ResetTimer()
	// b.N 是2的10次方 一个常量
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
