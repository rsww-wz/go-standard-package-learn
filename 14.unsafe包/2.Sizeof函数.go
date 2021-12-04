package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// unsafe.Sizeof 作用:查看变量类型所占字节大小
	// 和平台相关
	fmt.Println("基本数据类型")

	a := 4
	b := 3.14
	c := 'a'
	d := "hello"
	e := "good morning"

	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(d))
	fmt.Println(unsafe.Sizeof(e))
	fmt.Println()

	fmt.Println("引用数据类型")

	arr1 := [3]int{1, 2, 4}
	arr2 := [3]byte{'a', 'b', 'c'}

	var lst1 = []int{1, 2, 3}
	var lst2 = []byte{'a', 'b', 'c'}

	fmt.Println(unsafe.Sizeof(lst1))
	fmt.Println(unsafe.Sizeof(lst2))

	fmt.Println(unsafe.Sizeof(arr1))
	fmt.Println(unsafe.Sizeof(arr2))
	fmt.Println()

	fmt.Println("复合数据类型")

	type person struct {
		age  int
		name string
	}
	var p = person{20, "小明"}
	var p1 = person{21, "小王"}
	var pa = [2]person{}
	var pb = []person{p, p1, p}

	fmt.Println(unsafe.Sizeof(p))
	fmt.Println(unsafe.Sizeof(pa))
	fmt.Println(unsafe.Sizeof(pb))
}
