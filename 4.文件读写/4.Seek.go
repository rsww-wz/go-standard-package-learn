/*
功能：设置指针光标的位置
Seek(offset int 64，whence int) (int64,error)
	第一个参数：偏移量
	第二个参数：如何设置

	SeekStart    // seek relative to the origin of the file
	SeekCurrent  // seek relative to the current offset
	SeekEnd      // seek relative to the end
 */

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	address := `E:\0-code\Go\src\Golang\Go进阶\5.文件复制\test\源文件.txt`

	file,err := os.OpenFile(address,os.O_RDWR,os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	//读写
	bs := []byte{0}
	_,_ = file.Read(bs)
	fmt.Println(string(bs))

	// 设置光标位置：开头位置
	_,_ = file.Seek(4,io.SeekStart)
	_,_ = file.Read(bs)
	fmt.Println(string(bs))

	// 设置光标位置：光标所在位置
	_,_ = file.Seek(3,io.SeekCurrent)
	_,_ = file.Read(bs)
	fmt.Println(string(bs))

	//设置光标位置：末尾位置
	_,_ = file.Seek(0,io.SeekEnd)
	_,_ = file.WriteString("ABC")

}
