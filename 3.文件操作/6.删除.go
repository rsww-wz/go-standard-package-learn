/*
前提:如果文件或者目录不存在会报错

Remove():删除一级目录或者文件
RemoveAll():删除多级目录以及目录下的文件

注意：用代码删除的文件夹或者文件都是不经过回收站，要慎用
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	// 删除一级文件
	err := os.Remove(`E:\0-code\Go\src\Golang\Go进阶\3.文件操作\test\ab.txt`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件删除成功")

	// 删除一级文件夹
	err = os.Remove(`E:\0-code\Go\src\Golang\Go进阶\3.文件操作\test\test1`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件夹删除成功")

	// 删除多重文件夹
	// 如果文件夹内有文件也会被删除
	err = os.RemoveAll(`E:\0-code\Go\src\Golang\Go进阶\3.文件操作\test\test1`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("多级文件夹删除成功")
}
