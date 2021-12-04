package main

import (
	"fmt"
	"unsafe"
)


type w struct {
	a byte   // 1
	b int32  // 4
	c int64  // 8
}

// Alignof函数作用:对齐内存
// 在结构体中以字段长度最大的对齐内存

func main() {
	var w1 *w
	fmt.Println(unsafe.Sizeof(*w1))
	fmt.Println(unsafe.Alignof(w1))
	fmt.Println(unsafe.Alignof(w1.a))
	fmt.Println(unsafe.Alignof(w1.b))
	fmt.Println(unsafe.Alignof(w1.c))
}
