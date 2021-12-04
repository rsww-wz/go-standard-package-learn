package main

import (
	"fmt"
	"regexp"
)

func main() {
	/*
	    符号：中括号
	    匹配规则：中括号里面的内容是或关系
	        中括号内表示的是范围，在这个范围内有你需要匹配到的内容

	    理解：
	        中括号里面的正则表达式是一个范围
	        字符串有内容这个范围就可以匹配上
	        所以字符集里面是或的关系

	    中括号里面一般是元字符
	 */

	// 字符集的意思是匹配出来的也是一个字符，只不过这个字符的范围在字符集指定的范围里面就可以匹配
	// 字符集后面有数量词表示字符集范围里面的任何一个，只要符合条件都可以匹配多次（指定次数）
	str := "http://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html"
	re1 := regexp.MustCompile(`[\d\W]`)
	re2 := regexp.MustCompile(`[\d\W]+`)
	re3 := regexp.MustCompile(`[0-9]+`)
	re4 := regexp.MustCompile(`[a-z]+`)
	re5 := regexp.MustCompile(`[0-9][\W]`)
	re6 := regexp.MustCompile(`[\W][a-z]+[\W]`)
	fmt.Println(re1.FindAllString(str,-1))
	fmt.Println(re2.FindAllString(str,-1))
	fmt.Println(re3.FindAllString(str,-1))
	fmt.Println(re4.FindAllString(str,-1))
	fmt.Println(re5.FindAllString(str,-1))
	fmt.Println(re6.FindAllString(str,-1))
}
