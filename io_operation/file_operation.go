package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	/*
		copy文件
	*/
	// file1, _ := os.OpenFile("../cc.txt", os.O_RDWR|os.O_APPEND, os.ModePerm)
	// defer file1.Close()

	// file2, _ := os.OpenFile("../abc.txt", os.O_RDWR|os.O_APPEND, os.ModePerm)
	// defer file2.Close()

	// bytes := make([]byte, 10)
	// for {
	// 	n, err := file1.Read(bytes)
	// 	if err == io.EOF || n == 0 {
	// 		break
	// 	} else if err != nil {
	// 		fmt.Println("error", err)
	// 	}
	// 	file2.Write(bytes[:n])
	// }

	// 简单的方法 io.Copy
	// written, _ := io.Copy(file2, file1)
	// fmt.Println(written)

	// ioutil读取整个文件 一次性读写
	// bs, _ := ioutil.ReadFile(file1.Name())
	// ioutil.WriteFile(file2.Name(), bs, os.ModePerm)

	ioutil.ReadDir("../") // 读取本地文件夹下的目录

	ioutil.TempDir("../", "xx") // 临时目录

	ioutil.TempFile("../", "yy.txt") // 临时文件

	/*
		断点续传功能 Seek表示给当前的file文件句柄一个具体的位置读或者写
		Seek(offset int64, whence int) (int64, error)
		offset表示偏移量 whence表示位置 Seek表示相对于whence位置偏移多少来读取数据
		io.SeekStart
		io.SeekCurrent
		io.SeekEnd
		所谓的断点续传就是要记录临时文件 然后根据临时文件的大小 根据seek去找到位置然后读写
		边复制 边记录文件的总量
	*/
	// file, _ := os.OpenFile("../cc.txt", os.O_RDWR|os.O_APPEND, os.ModePerm)
	// defer file.Close()
	// bytes := make([]byte, 9)
	// _, _ = file.Seek(3, io.SeekStart)
	// file.Read(bytes)
	// fmt.Println(string(bytes))
	// // write也是可以的
	// file.Write(bytes)

	/*
		bufio的操作
	*/
	file, _ := os.OpenFile("../cc.txt", os.O_RDWR|os.O_APPEND, os.ModePerm)
	defer file.Close()
	reader := bufio.NewReader(file)
	strs, _ := reader.ReadBytes(100)
	fmt.Println(string(strs))
	reader.ReadLine()       // 读取一行
	reader.ReadString('\n') // 读取String 根据标志位来停止
	reader.Read(strs)

	writer := bufio.NewWriter(file)
	writer.Write(strs)
	writer.WriteString("ss")
	writer.WriteByte('s')
	writer.WriteRune('我')

	writer.Flush() // 要记得将缓存区的内容刷到磁盘中

}
