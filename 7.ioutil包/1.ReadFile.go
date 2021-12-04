package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 读取文件中的所有数据
	// 但是只会读取一次，如果文件很大，不适合用该方法
	// 不需要手动打开和关闭文件，只需要传地址即可
	// 如果需要读取的文件不是超大，这个方法还是很好用的
	address := `E:\Document\0-code\Golang\src\Golang\Go标准库\5.文件复制\test\目标文件.txt`
	data,err := ioutil.ReadFile(address)
	if err != nil{
		fmt.Println(err)
	}
	// data是字节类型
	//fmt.Println(data)
	// 转换成字符类类型
	fmt.Println(string(data))
}
