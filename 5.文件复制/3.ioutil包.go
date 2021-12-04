/*
ioutil包中的ioutil.WriteFile()和ioutil.ReadFile()
但是由于使用一次性读取文件，再一次性写入文件的方式，所以该方法不适用于大文件，容易内存溢出。
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	srcFile := 	`E:\0-code\Go\src\Golang\Go进阶\5.文件复制\test\源文件.txt`
	destFile := `E:\0-code\Go\src\Golang\Go进阶\5.文件复制\test\目标文件.txt`
	total,err := IOCopy(srcFile,destFile)
	fmt.Println(total,err)
}

func IOCopy(srcFile,destFile string) (int,error) {
	bs,err := ioutil.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}

	err = ioutil.WriteFile(destFile,bs,os.ModePerm)
	if err != nil {
		return 0, err
	}

	return len(bs),nil
}
