/*
写出数据


 */
package main

import (
	"fmt"
	"os"
)

func main() {
	address :=  `E:\0-code\Go\src\Golang\Go进阶\3.文件操作\test\2.txt`

	// 打开文件
	// 只读模式
	//file,err := os.Open(address)

	// 指定模式
	// os.O_CREATE:文件不存在可以创建
	// os.O_WRONLY:从头开始写入数据
	// os.O_APPEND:从文件的最后开始写入新数据，追加模式
	file,err := os.OpenFile(address,os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	//关闭文件
	defer file.Close()

	// 定义切片
	//bs := []byte{65,66,67,68,69,70} //ABCDEF
	bs := []byte{97,98,99,100}
	n,err := file.Write(bs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}
