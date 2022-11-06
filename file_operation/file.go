package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	/*
		文件信息 fileInfo
	*/
	// 获取文件信息
	fileInfo, err := os.Stat("/Users/zzw/go/src/deep/aa.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", fileInfo)
	// 文件名
	fmt.Println(fileInfo.Name())
	// 文件大小 字节为单位
	fmt.Println(fileInfo.Size())
	// 是否为目录
	fmt.Println(fileInfo.IsDir())
	// 修改时间
	fmt.Println(fileInfo.ModTime())
	// 文件权限
	fmt.Println(fileInfo.Mode())

	/*
		文件操作
			相对路径和绝对路径
	*/

	// 路径
	fileName1 := "/Users/zzw/go/src/deep/aa.txt"
	fileName2 := "ab.txt"
	// 判断是否是绝对路径
	fmt.Println(filepath.IsAbs(fileName1))
	fmt.Println(filepath.IsAbs(fileName2))
	// 获取绝对路径
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2))
	// 获取父目录
	fmt.Println(path.Join(fileName1, ".."))

	// 创建一个目录
	err = os.Mkdir("/Users/zzw/go/src/deep/bb", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	// 创建循环目录
	err = os.MkdirAll("/Users/zzw/go/src/deep/bb/cc/dd", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	/*
		创建文件
	*/
	// Create 文件存在会覆盖成空文件 文件不存在则直接创建 默认模式为0666 可读写 不可执行
	// 相对路径以当前工程路径为参照物
	file, err := os.Create("/Users/zzw/go/src/deep/cc.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)

	/*
		打开文件
	*/
	// Open打开文件 相对路径和绝对路径都可以写 打开的权限是只读的
	file, err = os.Open("../aa.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)

	// OpenFile 打开文件 文件名 文件的权限 文件不存在则进行创建 创建得到的文件的权限
	file, err = os.OpenFile("../aabb.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)

	// 关闭文件
	file.Close()

	// 删除文件或者文件夹 只能删除单个文件和单个文件夹
	err = os.Remove("../aabb.txt")
	if err != nil {
		fmt.Println(err)
	}
	// 循环删除文件以及文件目录
	err = os.RemoveAll("../bb")
	if err != nil {
		fmt.Println(err)
	}
}
