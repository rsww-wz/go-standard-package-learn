package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "http://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html"

	/*
	    \d：匹配0-9
	    \D：匹配除了0-9的字符
	    \w：匹配26大小写个字母和0-9和下划线
	    \W：除了w能匹配到的东西都可以匹配（包括不可见的字符，比如空号，回车等）
	    \s：专门匹配那些不可见的字符
	    \S：匹配可见字符

	    一个元字符只能匹配一个字符
	 */

	// 匹配数字
	re1 := regexp.MustCompile(`\d+`)
	re2 := regexp.MustCompile(`\D+`)
	fmt.Println(re1.FindAllString(str,-1))
	fmt.Println(re2.FindAllString(str,-1))

	// 匹配协议
	re3 := regexp.MustCompile(`\w+`)
	re4 := regexp.MustCompile(`\W+`)
	fmt.Println(re3.FindAllString(str,-1))
	fmt.Println(re4.FindAllString(str,-1))

	// 匹配空白
	re5 := regexp.MustCompile(`\s+`)
	re6 := regexp.MustCompile(`\S+`)
	fmt.Println(re5.FindAllString(str,-1))
	fmt.Println(re6.FindAllString(str,-1))
}
