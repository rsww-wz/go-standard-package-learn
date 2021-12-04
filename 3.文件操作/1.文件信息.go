  /*
fileInfo文件信息
	fileInfo，err := os.Stat(路径)
	返回的是一个对象
		name() 获取文件名
		size() 获取文件大小
		isDir() 是否成目录
		ModTime() 修改时间
		mode()   权限
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	// 文件信息
	// 反引号：原生字符
	address := `E:\多媒体\文档\小说\婆媳合奸.txt`
	fileInfo,err := os.Stat(address)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\n",fileInfo)

	// 获取文件名
	fmt.Println(fileInfo.Name())
	// 文件大小
	fmt.Println(fileInfo.Size())
	// 是否是目录
	fmt.Println(fileInfo.IsDir())
	// 修改时间
	fmt.Println(fileInfo.ModTime())
	// 获取权限
	fmt.Println(fileInfo.Mode())
}
