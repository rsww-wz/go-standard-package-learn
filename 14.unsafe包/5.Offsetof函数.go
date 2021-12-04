package main

import (
	"fmt"
	"unsafe"
)

type v struct {
	a byte  // 1
	b int32 // 4
	c int64 // 8
}

// offsetof函数作用：偏移
// 结构体内存是连续分布的，可以用offset函数偏移到结构体的下一个字段
// 对于没有可见性的结构体可以用这种方法修改结构体的私有属性

func main() {
	v1 := v{0, 0, 0}
	s := unsafe.Pointer(&v1)

	// uintptr指针类型，专门存放变量地址的类型
	field := uintptr(s)
	fieldNext := field + unsafe.Offsetof(v1.b)
	p1 := unsafe.Pointer(fieldNext)
	p2 := (*byte)(p1)
	*p2 = 20
	fmt.Println(v1)
}
