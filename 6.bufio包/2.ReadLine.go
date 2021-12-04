package main

import (
	"bufio"
	"fmt"
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

	//func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
	//ReadLine是一个低水平的行数据读取原语。大多数调用者应使用ReadBytes('\n')或ReadString('\n')代替，或者使用Scanner
	//ReadLine尝试返回一行数据，不包括行尾标志的字节
	//如果行太长超过了缓冲，返回值isPrefix会被设为true，并返回行的前面一部分。该行剩下的部分将在之后的调用中返回。
	data,flag,err := b1.ReadLine()
	fmt.Println(flag)
	fmt.Println(err)
	fmt.Println(data)
	fmt.Println(string(data))
}
