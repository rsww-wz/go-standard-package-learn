package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 写入数据，如果文件不存在就创建文件，可以指定文件权限
	// 但是如果文件有数据，会清空数据再写入数据
	// 参数：地址，字节切片，模式
	address := `E:\0-code\Go\src\Golang\Go进阶\5.文件复制\test\写入数据文件.txt`
	s1 := "hello world\nGood morning\n"
	err := ioutil.WriteFile(address,[]byte(s1),os.ModePerm)
	fmt.Println(err)
}
