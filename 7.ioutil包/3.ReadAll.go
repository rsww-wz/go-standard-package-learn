/*
ReadAll作用
	不止可以读文件，还可以读对象
	把要被读的数据，放进对象里面即可
 */
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	s1 := "hello world"
	address := `E:\0-code\Go\src\Golang\Go进阶\5.文件复制\test\写入数据文件.txt`
	str1 := readfile(address)
	str2 := readfile(s1)
	fmt.Println(str1)
	fmt.Println(str2)
}

func readfile(address string) string{
	// strings对象也实现了newReader接口
	comtent_obj := strings.NewReader(address)
	target_data,err := ioutil.ReadAll(comtent_obj)
	if err != nil {
		fmt.Println(err)
	}
	return string(target_data)
}