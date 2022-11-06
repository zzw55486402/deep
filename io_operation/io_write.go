package main

import (
	"fmt"
	"os"
)

func WriteFile() {
	/*
		写入数据
			Writer
				Write(b []byte) int, error
	*/
	// Append 追加写
	file, _ := os.OpenFile("../cc.txt", os.O_RDWR|os.O_APPEND, os.ModePerm)
	defer file.Close()
	str := "大家好"
	n, err := file.Write([]byte(str))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	// 直接写字符串
	file.WriteString("晚上好\n")
}
