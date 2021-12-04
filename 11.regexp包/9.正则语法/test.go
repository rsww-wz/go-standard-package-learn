package main

import (
	"fmt"
	"regexp"
)

func main() {
	f4()
}

// 贪婪匹配
func f1() {
	s := `"aaaaaaaaaaaaaa"`

	// 正则语法：非贪婪模式
	re := regexp.MustCompile(`aa+?`)
	res := re.FindAllString(s,-1)
	fmt.Println("正则语法:",res)

	// go语法1：非贪婪模式
	re = regexp.MustCompile(`(?U:aa+)`)
	res = re.FindAllString(s,-1)
	fmt.Println("go语法:",res)

	// go语法2：非贪婪模式
	re = regexp.MustCompile(`(?U)aa+`)
	res = re.FindAllString(s,-1)
	fmt.Println("go语法:",res)

	// 贪婪匹配
	re = regexp.MustCompile(`aa+`)
	res  = re.FindAllString(s,-1)
	fmt.Println("贪婪匹配:",res)
	fmt.Println()
}

// 匹配换行符
func f2() {
	s1 := `尽管心里有疑惑，可是苏玉雅并没有当面提出来，她微微一笑，礼貌而刻意保持着距离 <br />;

	刘斌见苏玉雅对自己嫣然一笑，顿时感觉脑袋晕乎乎的，辨不明东南，分不清西北了 <br />`

	// 不匹配换行符
	re1 := regexp.MustCompile(`.*<br />`)
	res := re1.FindAllString(s1,-1)
	fmt.Println("不匹配换行符:","个数:",len(res),res)

	// 正则语法:匹配换行符
	re1 = regexp.MustCompile(`[\w\W]*<br />`)
	res = re1.FindAllString(s1,-1)
	fmt.Println("正则语法:匹配换行符:","个数:",len(res))

	// go语法：匹配换行符
	re1 = regexp.MustCompile(`(?s).*<br />`)
	res = re1.FindAllString(s1,-1)
	fmt.Println("go语法:匹配换行符:","个数:",len(res))
}

// 多模式匹配
func f3() {
	s1 := `尽管心里有疑惑，可是苏玉雅并没有当面提出来，她微微一笑，礼貌而刻意保持着距离 <br />;

	刘斌见苏玉雅对自己嫣然一笑，顿时感觉脑袋晕乎乎的，辨不明东南，分不清西北了 <br /> &nbsp;&nbsp`

	// go语法：多模式匹配
	re2 := regexp.MustCompile(`(?U)(?s).*&nbsp`)
	res := re2.FindAllString(s1,-1)
	fmt.Println("贪婪并匹配换行","个数:",len(res),res,res[1])
}

// 分组匹配
// func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
// 函数无论如何都会匹配不分组的正则

// func (re *Regexp) FindStringSubmatch(s string) []string
// 只会返回第一个匹配到的结果
func f4() {
	s := `interesting excited moving surprising pleased bored tiring amazing frightening puzzled`

	// 编号捕获分组
	re := regexp.MustCompile(`(\w*)ing`)
	res := re.FindAllStringSubmatch(s,-1)
	fmt.Println("编号捕获分组",res)

	// 命名并编号捕获分组
	re = regexp.MustCompile(`(?P<ing>\w*)ing`)
	res = re.FindAllStringSubmatch(s,-1)
	fmt.Println("命名并编号捕获分组",res)

	// 不捕获的分组
	re = regexp.MustCompile(`(?:\w*)ing`)
	res = re.FindAllStringSubmatch(s,-1)
	fmt.Println("不捕获的分组",res)

	// 设置当前所在分组的标志，不捕获也不匹配
	re = regexp.MustCompile(`(?is)\w*ing`)
	res = re.FindAllStringSubmatch(s,-1)
	fmt.Println("设置当前所在分组的标志，不捕获也不匹配:",res)

	// 设置re段的标志，不捕获的分组
	re = regexp.MustCompile(`(?is:\w*)ing`)
	res = re.FindAllStringSubmatch(s,-1)
	fmt.Println("设置re段的标志，不捕获的分组:",res)
}

