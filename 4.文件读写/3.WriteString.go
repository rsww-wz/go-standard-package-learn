package main

import (
	"fmt"
	"os"
)

func main() {
	address :=  `E:\0-code\Go\src\Golang\Go进阶\3.文件操作\test\2.txt`

	file,err := os.OpenFile(address,os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	//关闭文件
	defer file.Close()

	// 写出字符串
	n,err := file.WriteString("hello world")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}
