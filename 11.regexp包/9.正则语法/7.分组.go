package main

import (
	"fmt"
	"regexp"
)

func main() {
	/*
	   匹配规则：把要匹配的字符串用小括号括起来即可
	   小括号里面是且的关系
	   组里面的才是匹配到的内容，组外面的不会被匹配
	   组的使用一般要配合其他正则表达式一起使用
	   总结：就是要匹配的内容必须跟小括号里面的字符串一样，才能匹配上
	   组的常用做法就是把表达式作为一个整体，然后用数量词重复
	*/

	str := "interesting exciting moving surprising pleasing boring tiring amazing frightening puzzling"
	str1 := "http://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html"

	// 命名分组匹配
	r := regexp.MustCompile(`(?P<year>\d+)/(?P<month>\d+)/(?P<day>\d+)`)

	// 获取所有匹配字符
	res := r.FindAllStringSubmatch(str1,-1)
	fmt.Println(res)

	// 获取分组字符（第一个元素是匹配整个字符串的字符）
	res1 := r.FindStringSubmatch(str1)
	fmt.Println(res1)

	// 获取分组名称
	//func (re *Regexp) SubexpNames() []string
	fmt.Println(r.SubexpNames())

	// 捕获分组数量
	//func (re *Regexp) NumSubexp() int
	fmt.Println(r.NumSubexp())

	// 编号分组匹配
	re := regexp.MustCompile(`(\w*)ing`)
	fmt.Println(re.FindAllString(str, -1))

	// 编号分组提取元素
	r2 := re.FindAllStringSubmatch(str, -1)
	for _, v := range r2 {
		word := v[1] + "e"
		fmt.Println(word)
	}

	// 非捕获的组 -- 不捕获分组
	re2 := regexp.MustCompile(`(?:\w*)ing`)
	fmt.Println(re2.FindAllString(str, -1))
	fmt.Println(re2.NumSubexp())

}