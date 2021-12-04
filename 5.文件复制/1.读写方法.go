/*
要求
	把一个文件的内容复制到另一个文件上

方法1
	通过io包下的read方法和write方法，边读边写，就能够实现文件的复制
 */
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile := 	`E:\0-code\Go\src\Golang\Go进阶\5.文件复制\test\源文件.txt`
	destFile := `E:\0-code\Go\src\Golang\Go进阶\5.文件复制\test\目标文件.txt`
	total,err := CopyFile(srcFile,destFile)
	fmt.Println(total,err)
}

// 该函数用于通过IO操作实现文件的拷贝，返回值是拷贝的总数量（字节），错误
func CopyFile(srcFile,destFile string) (int,error) {
	file1,err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2,err := os.OpenFile(destFile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err != nil {
		return 0, err
	}

	defer file1.Close()
	defer file2.Close()

	// 读写
	bs := make([]byte,1024,1024)
	n := -1
	total := 0
	for {
		n,err = file1.Read(bs)
		if err == io.EOF || n ==0 {
			fmt.Println("拷贝完毕")
			break
		} else if err != nil {
			fmt.Println("报错了")
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total,nil
}
