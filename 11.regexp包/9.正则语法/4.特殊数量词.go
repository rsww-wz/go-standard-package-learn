package main

import (
	"fmt"
	"regexp"
)

func main() {
	/*
		*:匹配0次或者无限多次
		+:匹配1次或者无限多次
		?:匹配0次或者1次
	 */

	str := "python C C# C++ java go object-C"
	re1 := regexp.MustCompile(`\w*`)
	re2 := regexp.MustCompile(`\w+`)
	re3 := regexp.MustCompile(`\w?`)
	fmt.Println(re1.FindAllString(str,-1))
	fmt.Println(re2.FindAllString(str,-1))
	fmt.Println(re3.FindAllString(str,-1))

	// 配合非贪婪模式
	str1 := "1000abcd123"
	re4 := regexp.MustCompile(`\d*?`)  //零次
	re5 := regexp.MustCompile(`\d+?`)  //一次
	re6 := regexp.MustCompile(`\d??`)  //零次
	fmt.Println(re4.FindAllString(str1,-1))
	fmt.Println(re5.FindAllString(str1,-1))
	fmt.Println(re6.FindAllString(str1,-1))
}
