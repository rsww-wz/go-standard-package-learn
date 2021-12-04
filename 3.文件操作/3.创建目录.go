/*
创建目录的前提
	目录不存在，否则会报错

创建目录用到的是os包
	Mkdir：只能创建一级目录
	MkdirAll:创建多级目录

参数
	第一个参数都是地址
	第二个参数是模式
		ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket
					| ModeDevice | ModeCharDevice | ModeIrregular
路径补充
	相对路径
		相对于当前工作目录
	绝对路径
		从盘符开始的目录

	.当前目录
	..上一层目录
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	//创建目录，只能创建一级目录
	//如果文件夹存在会报错
	err := os.Mkdir(`E:\0-code\Go\src\Golang\Go进阶\3.文件操作\test\test1`,os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件夹创建成功")

	//创建多级目录
	err = os.MkdirAll(`E:\0-code\Go\src\Golang\Go进阶\3.文件操作\test\aa\bb`,os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("多层文件夹创建成功")
}
