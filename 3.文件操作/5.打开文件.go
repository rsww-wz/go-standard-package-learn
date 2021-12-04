/*
Open:打开文件，默认模式是只读
fileOpen:打开文件，需要指定模式
	O_RDONLY     // open the file read-only.
	O_WRONLY     // open the file write-only.
	O_RDWR       // open the file read-write.
	O_APPEND     // append data to the file when writing.
	O_CREATE     // create a new file if none exists.
	O_EXCL       // used with O_CREATE, file must not exist.
	O_SYNC       // open for synchronous I/O.
	O_TRUNC      // truncate regular writable file when opened.

关闭文件
	文件对象.close()
 */

package main

import (
	"fmt"
	"os"
)

func main() {

	address := `E:\多媒体\文档\小说\婆媳合奸.txt`

	//打开文件,只读
	file3,err := os.Open(address)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file3)

	//打开文件，可读可写
	//需要指定模式
	file4,err := os.OpenFile(address,os.O_RDONLY|os.O_WRONLY,os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file4)

	// 关闭文件
	err = file4.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件关闭成功")
}
