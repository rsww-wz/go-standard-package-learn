package main

import (
	"fmt"
	"regexp"
)

func main() {
	//func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte
	//func (re *Regexp) ReplaceAllLiteralString(src, repl string) string
	//返回src的一个拷贝，将src中所有re的匹配结果都替换为repl。repl参数被直接使用，不会使用Expand进行扩展
	re := regexp.MustCompile("a(x*)b")
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "${1}"))

	fmt.Println("-------------------------------------------------------")

	//func (re *Regexp) ReplaceAll(src, repl []byte) []byte
	//func (re *Regexp) ReplaceAllString(src, repl string) string
	//ReplaceAllLiteral返回src的一个拷贝，将src中所有re的匹配结果都替换为repl
	//在替换时，repl中的'$'符号会按照Expand方法的规则进行解释和替换，例如$1会被替换为第一个分组匹配结果
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))

	//func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte
	//func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string
	//ReplaceAllLiteral返回src的一个拷贝，将src中所有re的匹配结果（设为matched）都替换为repl(matched)
	//repl返回的字符串被直接使用，不会使用Expand进行扩展


}
