package main

import (
	"fmt"
	"io"
	"os"
)

func ReadFile() {
	/*
		读取数据
			Reader接口
				Read(p []byte) (n int, err error)
	*/
	// 读取本地aa.txt文件中的数据
	file, _ := os.Open("../aa.txt")
	defer file.Close()

	for {
		bytes := make([]byte, 5)
		n, err := file.Read(bytes)
		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		fmt.Println(n)
		fmt.Println(string(bytes))
	}
	// ReadAt 从起始位置读
	// ReadFrom 从某个io.Reader中读取

}
