package main

import (
	"fmt"
	"regexp"
)

func main() {
	/*
	匹配规则
		组和字符集等，一个表达式只能匹配一个字符
		使用数量词可以匹配指定次数，相当于把组和字符集的表达式复制了几次
		是对花括号前面的那个表达式复制几次，不是对整个表达式复制几次
		所以前面一般是组，表示把组匹配多少次(重复分组多少次)

	数量词的范围
		如果匹配的内容长度不确定，数量词也是可以用范围表示的，数之间用逗号隔开
		一个数最小匹配次数和另外一个最大匹配次数
		也可以不写第一个或者第二个数字，即不限定最小值或者最大值

	贪婪与非贪婪
	    符号：问号，？
	    匹配规则：当符合匹配规则的时候，会尽可能多的匹配字符
	    当数量词范围是不确定的时候，默认是贪婪模式
	    非贪婪模式：满足匹配要求就会不继续匹配下去
	 */

	str := "http://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html"

	// 指定次数
	re1 := regexp.MustCompile(`\w{3}`)
	fmt.Println(re1.FindAllString(str,-1))

	// 指定范围：贪婪模式 -- 尽可能多匹配多的字符
	re2 := regexp.MustCompile(`\w{3,6}`)
	fmt.Println(re2.FindAllString(str,-1))

	// 指定范围：非贪婪模式 -- 匹配最少的字符
	re3 := regexp.MustCompile(`\w{3,6}?`)
	fmt.Println(re3.FindAllString(str,-1))
}
