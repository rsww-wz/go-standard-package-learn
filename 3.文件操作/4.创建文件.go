package main

import (
	"fmt"
	"os"
)

func main() {
	//创建文件
	//采用模式0666（任何人都可读可写，不可执行），如果文件已存在会截断它（为空文件）
	//支持绝对路径和相对路径
	file1,err1 := os.Create(`E:\0-code\Go\src\Golang\Go进阶\3.文件操作\test\ab.txt`)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(file1)
}
