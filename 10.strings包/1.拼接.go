package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// 方法1:最简单的加号方式，效率低
	str1 := "数据结构"
	str2 := "与算法"
	r1 := str1 + str2
	fmt.Println(r1)

	// 方法2：用切片的方式，可以自定义拼接的符号
	//func Join(a []string, sep string) string
	dataList := []string{"数据结构","与算法"}
	r2 := strings.Join(dataList,"")
	r3 := strings.Join(dataList,"|")
	fmt.Println(r2)
	fmt.Println(r3)

	// 方法3：创建一个bytes.Buffer对象
	// 把数据写入的内存中，然后统一拼接
	var buffer bytes.Buffer
	buffer.WriteString("数据结构")
	buffer.WriteString("与算法")
	r4 := buffer.String()
	fmt.Println(r4)

	// 方法4：创建一个Strings.Builder对象
	var builder strings.Builder
	builder.WriteString("数据结构")
	builder.WriteString("与算法")
	r5 := builder.String()
	fmt.Println(r5)

}
