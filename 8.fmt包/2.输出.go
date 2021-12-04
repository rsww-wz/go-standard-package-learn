package main

import (
	"fmt"
)

func main() {
	//func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	//Fprintf根据format参数生成格式化的字符串并写入w。返回写入的字节数和遇到的任何错误


	//func Sprintf(format string, a ...interface{}) string
	//Sprintf根据format参数生成格式化的字符串并返回该字符串
	s1 := fmt.Sprintf("%d-%d-%d %d:%d:%d\n",2021,2,25,10,30,18)
	fmt.Println(s1)
}
