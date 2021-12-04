package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	address := `E:\0-code\Go\src\Golang\Go进阶\3.文件操作\test\1.txt`
	file,err := os.Open(address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 创建reader对象
	b1 := bufio.NewReader(file)

	// ReadString
	//s1,err := b1.ReadString('\n')
	//fmt.Println(err)
	//fmt.Println(s1)

	// 按行读取文件
	for {
		s1,err := b1.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		fmt.Println(s1)
	}
}
