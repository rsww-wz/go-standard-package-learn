/*
TempDir:临时目录
TempFile：临时文件
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	address := `E:\0-code\Go\src\Golang\Go进阶\7.ioutil包`

	// 临时目录
	dir,err := ioutil.TempDir(address,"test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(dir)
	fmt.Println(dir)

	// 临时文件
	file,err := ioutil.TempFile(dir,"Test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(file.Name())
	fmt.Println(file.Name())
}
