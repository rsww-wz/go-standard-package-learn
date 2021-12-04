package main

/*
unsafe包作用
	提供了一些跳过go语言类型安全限制的操作

unsafe包只有两个结构，三个函数
两个结构
	type ArbitraryType int
		ArbitraryType在本文档里表示任意一种类型，但并非一个实际存在与unsafe包的类型

	type Pointer *ArbitraryType
		Pointer类型用于表示任意类型的指针。有4个特殊的只能用于Pointer类型的操作：
		1) 任意类型的指针可以转换为一个Pointer类型值
		2) 一个Pointer类型值可以转换为任意类型的指针
		3) 一个uintptr类型值可以转换为一个Pointer类型值
		4) 一个Pointer类型值可以转换为一个uintptr类型值

三个函数
	func Sizeof(x ArbitraryType) uintptr
		Sizeof返回类型v本身数据所占用的字节数。返回值是“顶层”的数据占有的字节数。
		例如，若v是一个切片，它会返回该切片描述符的大小，而非该切片底层引用的内存的大小

	func Offsetof(x ArbitraryType) uintptr
		Alignof返回类型v的对齐方式（即类型v在内存中占用的字节数）
		若是结构体类型的字段的形式，它会返回字段f在该结构体中的对齐方式

	func Alignof(x ArbitraryType) uintptr
		Offsetof返回类型v所代表的结构体字段在结构体中的偏移量，它必须为结构体类型的字段的形式
		换句话说，它返回该结构起始处与该字段起始处之间的字节数

Sizeof   变量大小
Offsetof 偏移量
Alignof  对齐
Pointer  指针类型
*/

func main() {

}