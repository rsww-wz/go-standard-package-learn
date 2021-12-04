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

	//func (b *Reader) ReadByte() (c byte, err error)
	//ReadByte读取并返回一个字节。如果没有可用的数据，会返回错误
	data,err := b1.ReadBytes('\n')
	fmt.Println(err)
	fmt.Println(string(data))
}
