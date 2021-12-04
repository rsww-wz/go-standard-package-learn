package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	address := `E:\0-code\Go\src\Golang\Go标准库\5.文件复制\test\目标文件.txt`
	file,err := os.OpenFile(address,os.O_CREATE|os.O_RDWR,os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 创建writer对象
	//func NewWriter(w io.Writer) *Writer
	//NewWriter创建一个具有默认大小缓冲、写入w的*Writer

	//func NewWriterSize(w io.Writer, size int) *Writer
	//NewWriterSize创建一个具有最少有size尺寸的缓冲、写入w的*Writer
	//如果参数w已经是一个具有足够大缓冲的*Writer类型值，会返回w
	w1 := bufio.NewWriter(file)

	//WriteString
	n,err := w1.WriteString("hello world\n")
	fmt.Println(err)
	fmt.Println(n)

	//func (b *Writer) Buffered() int
	//Buffered返回缓冲中已使用的字节数
	fmt.Println(w1.Buffered())
}
