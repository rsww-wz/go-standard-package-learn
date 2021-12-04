package main

import (
	"fmt"
	"regexp"
)

func main() {
	//func Compile(expr string) (*Regexp, error)
	//Compile解析并返回一个正则表达式。如果成功返回，该Regexp就可用于匹配文本
	//s := "hello world"
	reg1 := `\d\w{2}`
	re1,_ := regexp.Compile(reg1)
	fmt.Printf("%T,%v\n",re1,re1)

	//func MustCompile(str string) *Regexp
	//MustCompile类似Compile但会在解析失败时panic，主要用于全局正则表达式变量的安全初始化
	reg2 := `\w`
	re2 := regexp.MustCompile(reg2)
	fmt.Println(re2)

	//func (re *Regexp) String() string
	//String返回用于编译成正则表达式的字符串，就是将正则转化为字符串对象
	res := re1.String()
	fmt.Printf("%T,%v\n",res,res)
}
