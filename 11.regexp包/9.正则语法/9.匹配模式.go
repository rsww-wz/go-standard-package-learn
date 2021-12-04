package main

import (
	"fmt"
	"regexp"
)

func main() {
	/*
		i              大小写敏感（默认关闭）
		m              ^和$在匹配文本开始和结尾之外，还可以匹配行首和行尾（默认开启）
		s              让.可以匹配\n（默认关闭）
		U              非贪婪的：交换x*和x*?、x+和x+?……的含义（默认关闭）

		(?flags)       设置当前所在分组的标志，不捕获也不匹配
	    (?flags:re)    设置re段的标志，不捕获的分组
	 */

	// 不能像python在函数中设置是否需要此模式，而是需要在正则中设置标志位flag

	str := "hello \nworld\nend\n"

	// 匹配换行符
	re := regexp.MustCompile(`(?s:.+)`)
	fmt.Println(re.FindAllString(str,-1))

	// 非贪婪
	re1 := regexp.MustCompile(`\w{4,8}?`)
	fmt.Println(re1.FindAllString(str,-1))

	re2 := regexp.MustCompile(`(?U:\w{4,8})`)
	fmt.Println(re2.FindAllString(str,-1))

	// 大小写敏感
	str1 := "python Python"
	re3 := regexp.MustCompile(`python`)
	fmt.Println(re3.FindAllString(str1,-1))

	// 大小不写敏感
	re4 := regexp.MustCompile(`(?i:python)`)
	fmt.Println(re4.FindAllString(str1,-1))

	// 大小写敏感且支持换行符
	str2 := "Hello W\norl\nd"
	re5 := regexp.MustCompile(`(?s:H.+)|(?i:.H.+)`)
	fmt.Println(re5.FindAllString(str2,-1))

}
