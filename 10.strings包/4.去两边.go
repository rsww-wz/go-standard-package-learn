package main

import (
	"fmt"
	"strings"
)

func main() {
	//func TrimLeft(s string, cutset string) string
	//返回将s前端所有cutset包含的utf-8码值都去掉的字符串
	str1 := " hello world"
	r1 := strings.TrimLeft(str1," ")
	fmt.Println(r1)

	//func TrimRight(s string, cutset string) string
	//返回将s后端所有cutset包含的utf-8码值都去掉的字符串
	str2 := "hello world "
	r2 := strings.TrimRight(str2," ")
	fmt.Println(r2)

	//func Trim(s string, cutset string) string
	//返回将s前后端所有cutset包含的utf-8码值都去掉的字符串
	str3 := " hello world "
	str4 := "**hello world**"
	r3 := strings.Trim(str3," ")
	r4 := strings.Trim(str4,"*")
	fmt.Println(r3)
	fmt.Println(r4)

	//func TrimSpace(s string) string
	//返回将s前后端所有空白（unicode.IsSpace指定）都去掉的字符串
	r5 := strings.TrimSpace(str3)
	fmt.Println(r5)

	//func TrimPrefix(s, prefix string) string
	//返回去除s可能的前缀prefix的字符串
	var s1 = "Goodbye,,world!"
	s1 = strings.TrimPrefix(s1, "Goodbye,")
	s1 = strings.TrimPrefix(s1, "Howdy,")
	fmt.Println("Hello" + s1)

	//func TrimSuffix(s, suffix string) string
	//返回去除s可能的后缀suffix的字符串
	var s2 = "Hello,goodbye, etc!"
	s2 = strings.TrimSuffix(s2, "goodbye, etc!")
	s2 = strings.TrimSuffix(s2, "planet")
	fmt.Println(s2 + "world!")
}
