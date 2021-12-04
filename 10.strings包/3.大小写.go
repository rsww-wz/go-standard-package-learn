package main

import (
	"fmt"
	"strings"
)

func main() {
	//func Title(s string) string
	//返回s中每个单词的首字母都改为标题格式的字符串拷贝
	str1 := "hello world"
	r1 := strings.Title(str1)
	fmt.Println(r1)

	//func ToTitle(s string) string
	//返回将所有字母都转为对应的标题版本的拷贝
	str2 := "Good morning"
	r2 := strings.ToTitle(str2)
	fmt.Println(r2)

	//func ToLower(s string) string
	//返回将所有字母都转为对应的小写版本的拷贝
	name1 := "LUCY"
	result1 := strings.ToLower(name1)
	fmt.Println(result1)

	//func ToUpper(s string) string
	//返回将所有字母都转为对应的大写版本的拷贝
	name2 := "lili"
	result2 := strings.ToUpper(name2)
	fmt.Println(result2)

	// 对中文，无论是变大写还是小写都没有反应
	name3 := "中国"
	result3 := strings.ToUpper(name3)
	result4 := strings.ToLower(name3)
	fmt.Println(result3)
	fmt.Println(result4)
}
