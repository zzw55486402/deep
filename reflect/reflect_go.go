package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello,", msg)
}

func (p Person) PrintInfo() {
	fmt.Println("info: ", p)
}

func main() {
	// 反射基于空接口 包括type 和 value两个指针
	// 在程序运行的时候动态的获取接口类型的变量和值
	// Typef 获取类型 ValueOf 获取数值
	var x float64 = 3.14
	t := reflect.TypeOf(x)  // reflect.Type
	v := reflect.ValueOf(x) // reflect.Value
	fmt.Println(t, v)

	// 根据反射的值来获取对应的类型和数值
	// v.Kind 根据value来获取数据类型
	// v.Type 直接获取数据类型
	// v.Float 打印相应类型的数值
	v = reflect.ValueOf(x)
	fmt.Println(v.Kind() == reflect.Float64)
	fmt.Println(v.Type())
	fmt.Println(v.Float())

	// 通过接口得到反射对象
	// 通过反射来得到接口
	// 要操作一个反射变量 他的值必须是可以修改的

	// 知道原有的数据类型
	var num float64 = 1.23
	// 接口类型变量 -> 反射类型对象
	value := reflect.ValueOf(num)
	// value.Interface() 得到该类型的接口然后断言（已知数据类型的情况下）
	// 反射类型对象 -> 接口类型变量
	convertValue := value.Interface().(float64)
	fmt.Println(convertValue)

	// 注意类型 指针要和数据类型区分开
	// 反射类型对象 -> 接口类型变量 理解为强制转换 所以数据类型一定要符合
	pointer := reflect.ValueOf(&num)
	converPointer := pointer.Interface().(*float64)
	fmt.Println(*converPointer)

	// 不知道原有的数据类型

	p1 := Person{
		Name: "zzw",
		Age:  18,
		Sex:  "male",
	}
	GetMessage(p1)

	// 操作结构体的赋值
	// 改变数值
	valuePerson := reflect.ValueOf(&p1)
	fmt.Println(&p1)
	if valuePerson.Kind() == reflect.Ptr {
		newValuePerson := valuePerson.Elem()
		fmt.Println(newValuePerson.CanSet())

		newValuePerson.FieldByName("Age").SetInt(199)            // 根据字段名来获取字段并赋值
		newValuePerson.FieldByName("Name").SetString("shinyruo") // 根据字段名来获取字段并赋值
		fmt.Println(p1)
	}

	// 通过reflect来操作方法的调用和函数的调用
	/*
		接口变量 -> 对象反射对象 value
		获取对应的方法对象 MethodByName
		调用方法对象 Call
	*/
	valueStruct := reflect.ValueOf(p1)
	fmt.Println(valueStruct.Kind(), valueStruct.Type())

	methodValue1 := valueStruct.MethodByName("Say")
	methodValue2 := valueStruct.MethodByName("PrintInfo")
	fmt.Println(methodValue2.Kind(), methodValue2.Type())
	methodValue2.Call([]reflect.Value{}) // 调用传值
	methodValue2.Call(nil)
	// 调用带参的方法
	methodValue1.Call([]reflect.Value{reflect.ValueOf("爆了")})

	// 通过反射调用函数
	// 函数到反射对象Value
	// kind -> func
	// call 调用
	f1 := fun1
	valueFunc1 := reflect.ValueOf(f1)
	fmt.Println(valueFunc1.Kind(), valueFunc1.Type())
	valueFunc1.Call(nil)

	f2 := fun2
	valueFunc2 := reflect.ValueOf(f2)
	valueFunc2.Call([]reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf("gogogo"),
	})
	f3 := fun3
	valueFunc3 := reflect.ValueOf(f3)
	result := valueFunc3.Call([]reflect.Value{
		reflect.ValueOf(2),
		reflect.ValueOf("csgo"),
	})
	fmt.Println(result)
	fmt.Println(result[0].Kind(), result[0].Type())

	s := result[0].Interface().(string)
	fmt.Println(s)
}

func fun1() {
	fmt.Println("hello reflect")
}
func fun2(i int, s string) {
	fmt.Println("func2", i, s)
}
func fun3(i int, s string) string {
	fmt.Println("func3", i, s)
	return "爆了"
}

func GetMessage(input interface{}) {
	getType := reflect.TypeOf(input) // 先获取input的类型
	fmt.Println(getType.Name())      // 类型名称 Person
	fmt.Println(getType.Kind())      // 种类 struct

	getValue := reflect.ValueOf(input) // 获取结构体数值
	fmt.Println("get all fields is: ", getValue)

	// 获取结构体字段名字以及单个数值
	// 先获取Type对象
	/*
		NumField() 字段个数
		Field(index) 获取字段

		通过Field获取每个Field字段

		Interface() 得到每一个的Value
	*/

	for i := 0; i < getType.NumField(); i++ {
		getField := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Println("字段名称: ", getField.Name) // 字段名称
		fmt.Println("字段名称: ", getField.Type) // 字段类型
		fmt.Println("值: ", value)            // 字段类型
	}

	// 获取方法
	for i := 0; i < getType.NumMethod(); i++ {
		getMethod := getType.Method(i)
		fmt.Println("方法名称: ", getMethod.Name) // 方法名称
		fmt.Println("方法名称: ", getMethod.Type) // 方法类型
	}

	// 通过reflect.Value设置实际的变量值
	var number float64 = 1.2345
	fmt.Println(number)

	// 需要操作指针来改变数值
	// 使用reflect.ValueOf() 获取number的value对象
	numberValue := reflect.ValueOf(&number) // 操作指针才可以
	newValue := numberValue.Elem()          // 获取原始值的对象

	fmt.Println("类型: ", newValue.Type())
	fmt.Println("是否可以赋值: ", newValue.CanSet()) // 是否可以赋值

	// 重新赋值
	newValue.SetFloat(3.14)
	fmt.Println(number)

}
