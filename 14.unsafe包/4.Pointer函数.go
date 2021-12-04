package main

import (
	"fmt"
	"unsafe"
)

type ww struct {
	a byte  // 1
	b int32 // 4
	c int64 // 8
}

// Pointer函数作用:修改指针类型
// 任意类型指针都可以转换为pointer指针，所以pointer指针是通用指针类型
// 然后再把通用指针改成你需要的类型，最后再通过指针修改值
// 不同类型的变量不能直接赋值，但是值是一样的，可以采用这种方式
// byte -- pointer -- int

func main() {
	//a := ww{0, 0, 0}
	//
	//// 将结构体指针转换为通用指针
	//s := unsafe.Pointer(&a)
	//pb := (*byte)(s)
	//*pb = 10
	//fmt.Println(a)
	//
	//pc := (*string)(s)
	//*pc = "h"
	//fmt.Println(a)

	var v1 byte
	v1 = 10
	v2 := 20

	// 类型不匹配
	//v1 = v2

	// 把v1转成通用指针
	v := unsafe.Pointer(&v1)
	v3 := (*int)(v)
	fmt.Println(v3)
	v2 = *v3
	fmt.Println(v2)

}
