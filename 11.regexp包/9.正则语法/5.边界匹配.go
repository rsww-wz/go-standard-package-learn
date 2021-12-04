package main

import (
	"fmt"
	"regexp"
)

func main() {
	/*
	    语法：在正则表达式的最前面加上一个乘方符号(^)，最后面加上一个美元符号($)
	    乘方符号(^)：从字符串的开头开始匹配（字符串的开头必须是你要匹配的那几个字符）
	    美元符号($)：从字符串的最后开始匹配（字符串的结尾必须是你要匹配的那几个字符）
	    这两个符号可以单独使用

		\b:是一个位置匹配，它的作用匹配单词，匹配的具体位置就是\w与\W之间的位置
		如何理解这个位置
			连续的的字符它不会匹配，因为它没有\w与\W之间的位置，都是\w
			连续的的空白字符它不会匹配，因为它没有也\w与\W之间的位置,都是\W
			只有空白符和字符连续，它才会匹配因为他们复合\w\W之间的位置

		\B:就是\b的反面的意思，非单词边界
			例如在字符串中所有位置中，扣掉\b，剩下的都是\B的
			也就是连续的字符或者连续的空白的第一个字符的后面开始匹配

	*/

	str := "https://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html"

	// 开头匹配：字符串必须以^后面的字符开头，否则不匹配
	re1 := regexp.MustCompile(`^http[\w\W]+`)
	fmt.Println(re1.FindAllString(str,-1))

	//结尾匹配：字符串必须以$后面的字符结尾，否则不匹配
	re2 := regexp.MustCompile(`http[\w\W]+html$`)
	fmt.Println(re2.FindAllString(str,-1))

	// 同时匹配:也就如果能匹配成功的话，就是匹配整个字符串了
	re3 := regexp.MustCompile(`^http[\w\W]+html$`)
	fmt.Println(re3.FindAllString(str,-1))

	str1 := "Whatever is worth doing is worth doing well"
	re4 := regexp.MustCompile(`\b[\w]+`)
	fmt.Println(re4.FindAllString(str1,-1))

	re5 := regexp.MustCompile(`\B[\w]+`)
	fmt.Println(re5.FindAllString(str1,-1))
}
