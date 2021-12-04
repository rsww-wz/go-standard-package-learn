package main

import (
	"fmt"
	"strings"
)

func main() {
	//func EqualFold(s, t string) bool
	//判断两个utf-8编码字符串是否相同(将unicode大写、小写、标题三种格式字符视为相同)
	fmt.Println(strings.EqualFold("Go","go"))

	//func HasPrefix(s, prefix string) bool
	//判断s是否有前缀字符串prefix
	fmt.Println(strings.HasPrefix("**hello world**","*"))

	//func HasSuffix(s, suffix string) bool
	//判断s是否有后缀字符串suffix
	fmt.Println(strings.HasSuffix("**hello world**","*"))

	//func Contains(s, substr string) bool
	//判断字符串s是否包含子串substr
	fmt.Println(strings.Contains("hello world","oo"))

	//func ContainsAny(s, chars string) bool
	//判断字符串s是否包含字符串chars中的任一字符
	fmt.Println(strings.ContainsAny("hello world","oo"))

	//func Count(s, sep string) int
	//返回字符串s中有几个不重复的sep子串
	fmt.Println(strings.Count("hello world","l"))
}
