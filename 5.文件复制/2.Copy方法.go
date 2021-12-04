/*
函数原型
	func copy(dst Writer,src Reader) (written int64,err error)

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
	total,err := MyCopy(srcFile,destFile)
	fmt.Println(total,err)
}

func MyCopy(srcFile,destFile string) (int64,error) {
	file1,err:=os.Open(srcFile)
	if err!=nil{
		return 0, err
	}

	file2,err:=os.OpenFile(destFile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err != nil{
		return 0, err
	}

	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2,file1)
}
