package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	address := `E:\0-code\Go\src\Golang\Go进阶\5.文件复制\test\目标文件.txt`
	file,err := os.OpenFile(address,os.O_CREATE|os.O_RDWR,os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 创建writer对象
	w1 := bufio.NewWriter(file)

	// 要写入的数据大于缓存区
	// 会把缓存区满了的数据写入文件

	for i:=0;i<=1000;i++ {
		w1.WriteString(fmt.Sprintf("%d:hello\n",i))
	}

	// 缓存区没有的数据不会写入文件
	// 需要手动flush写入
	// 刷新缓存区,即把缓存区的数据写入文件
	//func (b *Writer) Flush() error
	//Flush方法将缓冲中的数据写入下层的io.Writer接口
	w1.Flush()
}
