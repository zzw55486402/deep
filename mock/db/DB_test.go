package db

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

// go test . -cover -v

/*
	写可测试的代码与写好测试用例是同等重要的，如何写可 mock 的代码呢？
		mock 作用的是接口，因此将依赖抽象为接口，而不是直接依赖具体的类。
		不直接依赖的实例，而是使用依赖注入降低耦合性。

	在软件工程中，依赖注入的意思为，给予调用方它所需要的事物。 “依赖”是指可被方法调用的事物。依赖注入形式下，
	调用方不再直接指使用“依赖”，取而代之是“注入” 。“注入”是指将“依赖”传递给调用方的过程。在“注入”之后，调用方才会调用该“依赖”。
	传递依赖给调用方，而不是让让调用方直接获得依赖，这个是该设计的根本需求。

*/

func TestGetFromDB(t *testing.T) {
	// ctrl 是整个mock的一个控制器  用来判断是否调用了mock方法
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// 通过mock来生成接口对应的对象
	m := NewMockDB(ctrl)
	// 给对应的接口方法设置期待的返回值
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))
	m.EXPECT().Get(gomock.Eq("Jerry")).Return(100, nil)

	/*
		Eq(value) 表示与 value 等价的值。
		Any() 可以用来表示任意的入参。
		Not(value) 用来表示非 value 以外的值。
		Nil() 表示 None 值
	*/
	m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	m.EXPECT().Get(gomock.Any()).Return(630, nil)
	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
	m.EXPECT().Get(gomock.Nil()).Return(0, errors.New("nil"))

	/*
		Return 返回确定的值
		Do Mock 方法被调用时，要执行的操作吗，忽略返回值。
		DoAndReturn 可以动态地控制返回值。
	*/
	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
	m.EXPECT().Get(gomock.Any()).Do(func(key string) {
		t.Log(key)
	})
	m.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string) (int, error) {
		if key == "Sam" {
			return 630, nil
		}
		return 0, errors.New("not exist")
	})

	/*
		Times() 断言 Mock 方法被调用的次数。
		MaxTimes() 最大次数。
		MinTimes() 最小次数。
		AnyTimes() 任意次数（包括 0 次）。
	*/
	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil).Times(2)

	// Inorder() 的调用顺序
	o1 := m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	o2 := m.EXPECT().Get(gomock.Eq("Sam")).Return(630, nil)
	gomock.InOrder(o1, o2)

	type args struct {
		db  DB
		key string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				db:  m,
				key: "Tom",
			},
			want: -1,
		},
		{
			name: "test2",
			args: args{
				db:  m,
				key: "Jerry",
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFromDB(tt.args.db, tt.args.key); got != tt.want {
				t.Errorf("GetFromDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
