package main

import (
	"fmt"
	"regexp"
)

func main() {
	/*
		通配符
			. (英文句号)：匹配1次除了换行符之外的所有符号

		匹配所有字符
			.*:匹配除了换行符之外的字符无限次

		匹配全部字符
			\d\D
			\w\W
			\s\S

	    ^:表示匹配全集的补集
			^\d ==\D
			^\w == \W
			^\s == \S
	 */

	str := "http://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html"
	str1 := "http://www.flysnow.org/2018/01/20/golang-\ngoquery-examples-selector.html"
	re1 := regexp.MustCompile(`.*`)
	fmt.Println(re1.FindAllString(str,-1))
	fmt.Println(re1.FindAllString(str1,-1))
}
