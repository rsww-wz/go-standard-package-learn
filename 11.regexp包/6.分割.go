package main

import (
	"fmt"
	"regexp"
)

func main() {
	//func (re *Regexp) Split(s string, n int) []string
	//Split将re在s中匹配到的结果作为分隔符将s分割成多个字符串，并返回这些正则匹配结果之间的字符串的切片。
	//返回的切片不会包含正则匹配的结果，只包含匹配结果之间的片段。当正则表达式re中不含正则元字符时，本方法等价于strings.SplitN

	//n > 0 : 返回最多n个子字符串，最后一个子字符串是剩余未进行分割的部分。
	//n == 0: 返回nil (zero substrings)
	//n < 0 : 返回所有子字符串
	re := regexp.MustCompile(`a*`)
	fmt.Println(re.Split("say,name", -1))
	fmt.Println(re.Split("abaabaccadaaae", -1))
	fmt.Println(re.Split("abaabaccadaaae", 5))

	re1 := regexp.MustCompile(`\d`)
	fmt.Println(re1.Split("h4el7lo w8or9l99d",-1))
	fmt.Println(re1.Split("h4el7lo w8or9l99d",4))

}
