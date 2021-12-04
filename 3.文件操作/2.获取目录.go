/*
filepath是一个对象
	isAbs:判断是否是绝对路径
	Abs:获取绝对路径
	Base
		获取目录的最后一个元素
		如果路径是空字符，返回空白
		如果路径只有斜线，返回/
 */


package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {

	address := `E:\多媒体\文档\小说\婆媳合奸.txt`

	// 判断绝对路径
	fileName1 := address
	fileName2 := "1.txt"
	fmt.Println(filepath.IsAbs(fileName1))
	fmt.Println(filepath.IsAbs(fileName2))

	//获取绝对路径
	fmt.Println(filepath.Abs(fileName1))

	// 获取父目录
	fmt.Println(path.Join(fileName1,".."))

	//返回路径最后一个元素
	fmt.Println(filepath.Base("./1.txt"))

	//如果路径为空字符串，返回空
	fmt.Println(filepath.Base(""))

	//如果路径只有斜线，返回/
	fmt.Println(filepath.Base("///"))

	//返回路径最后一个元素的目录，路径为空则返回空
	//支持相对路径和绝对路径
	fmt.Println(filepath.Dir("./a/b/c"))
	fmt.Println(filepath.Dir("C:/a/b/c"))

	//返回路径中的扩展名，如果没有点，返回空
	fmt.Println(filepath.Ext("./a/b/c/d.jpg"))

	//分割路径中的目录与文件
	dir, file := filepath.Split("C:/a/b/c/d.jpg")
	fmt.Println(dir, file)

	//返回所有匹配的文件
	match, _ := filepath.Glob("./*.go")
	fmt.Println(match)

	//将路径中的/替换为路径分隔符
	fmt.Println(filepath.FromSlash("./a/b/c"))

	//将路径分隔符使用/替换
	fmt.Println(filepath.ToSlash("C:/a/b"))

	//返回分区名
	fmt.Println(filepath.VolumeName("C:/a/b/c"))

	//遍历指定目录下所有文件
	_ = filepath.Walk(
		"./",
		func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)
			return nil
		},
	)
}
