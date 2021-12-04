package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	address := `E:\0-code\Go\src\Golang\Go进阶\3.文件操作\test\1.txt`
	file,err := os.Open(address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// scanln 不能读取空格
	//s1 := ""
	//fmt.Scanln(&s1)
	//fmt.Scanln(s1)
	//fmt.Println(s1)

	// 创建对象
	//func NewScanner(r io.Reader) *Scanner
	//NewScanner创建并返回一个从r读取数据的Scanner，默认的分割函数是ScanLines
	b2 := bufio.NewReader(os.Stdin)
	s2,_ := b2.ReadString('\n')
	fmt.Println(s2)


}
